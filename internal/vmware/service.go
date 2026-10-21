package vmware

import (
	"context"
	"fmt"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"math"
	"path"
	"reflect"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"time"
	"vmware-exporter/pkg/logging"
)

var _ Service = &service{}

type service struct {
	logger         *logging.Logger
	vmwareHost     string
	vmwareUser     string
	vmwarePassword string
	scrapeTimeout  time.Duration
}

type Service interface {
	statuses() ([]*Status, error)
	error(err error)
}

var interval = 20

const (
	defaultScrapeTimeout = 60 * time.Second
	// logoutTimeout limits cleanup calls, which must not depend on the (possibly expired) scrape context.
	logoutTimeout = 10 * time.Second
)

// parseTimeout accepts a Go duration ("30s") or a plain number of seconds ("30").
func parseTimeout(v string, def time.Duration) time.Duration {
	v = strings.TrimSpace(v)
	if v == "" {
		return def
	}
	if d, err := time.ParseDuration(v); err == nil && d > 0 {
		return d
	}
	if n, err := strconv.Atoi(v); err == nil && n > 0 {
		return time.Duration(n) * time.Second
	}
	return def
}

func NewService(l *logging.Logger, host string, config interface{}) Service {
	cfg := reflect.ValueOf(config).Elem()

	vmwareUser := cfg.FieldByName("VmwareUser").Interface().(string)
	if len(vmwareUser) == 0 {
		l.Debug("user name for VMWare auth not specified, use default user name")
		vmwareUser = "monitoring"
	}

	vmwarePass := cfg.FieldByName("VmwarePass").Interface().(string)
	if len(vmwarePass) == 0 {
		l.Debug("user password for VMWare auth not specified, use default password")
		vmwarePass = "password"
	}

	scrapeTimeout := parseTimeout(cfg.FieldByName("ScrapeTimeout").Interface().(string), defaultScrapeTimeout)
	if scrapeTimeout == defaultScrapeTimeout {
		l.Debug("scrape timeout not specified or invalid, use default timeout")
	}

	return &service{
		logger:         l,
		vmwareHost:     host,
		vmwareUser:     vmwareUser,
		vmwarePassword: vmwarePass,
		scrapeTimeout:  scrapeTimeout,
	}
}

func (s *service) error(err error) {
	s.logger.Error(err.Error())
}

// scrapeData is the inventory-wide data shared by the per-host collection.
type scrapeData struct {
	dsNames      map[string]string // datastore moref -> name
	dsInstances  map[string]string // datastore uuid (perf counter instance) -> name
	hostCluster  map[string]string // host moref -> cluster name
	vmPerf       map[string][]vmMetric
	vmDSPerf     map[string][]vmMetric
	hostPerf     map[string][]vmMetric
	hostInstPerf map[string][]vmMetric
}

func (s *service) perfOrLog(m map[string][]vmMetric, err error) map[string][]vmMetric {
	if err != nil {
		s.logger.Error(fmt.Sprintf("collect performance metrics from %s failed: %v", s.vmwareHost, err))
	}
	return m
}

// maxHostWorkers limits parallel per-host requests so a big vCenter is not flooded.
const maxHostWorkers = 8

// statuses collects metrics from the target, which may be a standalone ESXi host or a vCenter:
// inventory objects (hosts, datastores, VMs) are retrieved from the root folder and every host
// gets its own Status, so the metrics keep the same host_name labels in both modes.
func (s *service) statuses() ([]*Status, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.scrapeTimeout)
	defer cancel()
	c, err := NewClient(ctx, s.vmwareHost, s.vmwareUser, s.vmwarePassword)
	if err != nil {
		return nil, err
	}
	defer func() {
		// the scrape ctx may already be expired, use a fresh one so the session is always released
		lctx, lcancel := context.WithTimeout(context.Background(), logoutTimeout)
		defer lcancel()
		if err := c.Logout(lctx); err != nil {
			s.logger.Error(fmt.Sprintf("logout from %s failed: %v", s.vmwareHost, err))
		}
	}()
	m := view.NewManager(c.Client)
	root := c.ServiceContent.RootFolder

	var hss []mo.HostSystem
	if err := retrieveView(ctx, m, root, "HostSystem", []string{"summary", "datastore", "triggeredAlarmState"}, &hss); err != nil {
		return nil, fmt.Errorf("retrieve HostSystem: %w", err)
	}
	if len(hss) == 0 {
		return nil, fmt.Errorf("%s returned no HostSystem objects", s.vmwareHost)
	}

	var dss []mo.Datastore
	if err := retrieveView(ctx, m, root, "Datastore", []string{"summary", "host", "vm", "triggeredAlarmState"}, &dss); err != nil {
		return nil, fmt.Errorf("retrieve Datastore: %w", err)
	}

	var vms []mo.VirtualMachine
	if err := retrieveView(ctx, m, root, "VirtualMachine", []string{"summary", "snapshot", "storage", "guest.disk", "triggeredAlarmState"}, &vms); err != nil {
		return nil, fmt.Errorf("retrieve VirtualMachine: %w", err)
	}

	vmsRefs := make([]types.ManagedObjectReference, 0, len(vms))
	vmsByHost := make(map[string][]mo.VirtualMachine)
	for _, vm := range vms {
		vmsRefs = append(vmsRefs, vm.Self)
		if h := vm.Summary.Runtime.Host; h != nil {
			vmsByHost[h.Value] = append(vmsByHost[h.Value], vm)
		}
	}

	// a datastore is reported for every host it is mounted on, like a scrape of each ESXi would do.
	// Both sides of the relation are used (Datastore.host and HostSystem.datastore) and merged.
	dsInfo := make(map[string]totalds, len(dss))
	dsByHost := make(map[string][]totalds)
	dsSeen := make(map[[2]string]struct{})
	dsHosts := make(map[string]float64) // number of hosts a datastore is mounted on
	addDS := func(host, ds string) {
		item, ok := dsInfo[ds]
		if _, dup := dsSeen[[2]string{host, ds}]; !ok || dup {
			return
		}
		dsSeen[[2]string{host, ds}] = struct{}{}
		dsByHost[host] = append(dsByHost[host], item)
		dsHosts[ds]++
	}
	dsNames := make(map[string]string, len(dss))     // datastore moref -> name
	dsInstances := make(map[string]string, len(dss)) // datastore uuid (perf counter instance) -> name
	for _, ds := range dss {
		dsInfo[ds.Self.Value] = totalds{
			ref:         ds.Self.Value,
			dsname:      ds.Summary.Name,
			capacity:    float64(ds.Summary.Capacity),
			freespace:   float64(ds.Summary.FreeSpace),
			uncommitted: float64(ds.Summary.Uncommitted),
			dsType:      ds.Summary.Type,
			url:         ds.Summary.Url,
			accessible:  ds.Summary.Accessible,
			maintenance: datastoreMaintenanceMode(ds.Summary.MaintenanceMode),
			vms:         float64(len(ds.Vm)),
		}
		item := dsInfo[ds.Self.Value]
		item.redAlarms, item.yellowAlarms = alarmCounts(ds.TriggeredAlarmState)
		dsInfo[ds.Self.Value] = item
		dsNames[ds.Self.Value] = ds.Summary.Name
		if uuid := path.Base(strings.TrimRight(ds.Summary.Url, "/")); uuid != "." && uuid != "/" {
			dsInstances[uuid] = ds.Summary.Name
		}
	}
	for _, ds := range dss {
		for _, mount := range ds.Host {
			addDS(mount.Key.Value, ds.Self.Value)
		}
	}
	for _, h := range hss {
		for _, ref := range h.Datastore {
			addDS(h.Self.Value, ref.Value)
		}
	}

	for host, list := range dsByHost {
		for i := range list {
			list[i].hosts = dsHosts[list[i].ref]
		}
		dsByHost[host] = list
	}

	// clusters are optional (a standalone host has none): a failure only drops the cluster label
	hostCluster := make(map[string]string)
	var clusters []mo.ClusterComputeResource
	if err := retrieveView(ctx, m, root, "ClusterComputeResource", []string{"name", "host"}, &clusters); err != nil {
		s.logger.Error(fmt.Sprintf("retrieve ClusterComputeResource from %s failed: %v", s.vmwareHost, err))
	}
	for _, cl := range clusters {
		for _, h := range cl.Host {
			hostCluster[h.Value] = cl.Name
		}
	}

	// performance counters are optional: a failure must not drop all the other metrics
	data := &scrapeData{dsNames: dsNames, dsInstances: dsInstances, hostCluster: hostCluster}
	if pc, err := newPerfClient(ctx, c, s.logger); err != nil {
		s.logger.Error(fmt.Sprintf("collect performance metrics from %s failed: %v", s.vmwareHost, err))
	} else {
		var hostRefs []types.ManagedObjectReference
		for _, h := range hss {
			if h.Summary.Runtime != nil && h.Summary.Runtime.ConnectionState == types.HostSystemConnectionStateConnected &&
				h.Summary.Runtime.PowerState == types.HostSystemPowerStatePoweredOn {
				hostRefs = append(hostRefs, h.Self)
			}
		}
		data.vmPerf = s.perfOrLog(pc.query(ctx, vmsRefs, pc.allCounters(), ""))
		data.vmDSPerf = s.perfOrLog(pc.query(ctx, vmsRefs, vmInstancedCounters, "*"))
		data.hostPerf = s.perfOrLog(pc.query(ctx, hostRefs, hostCounters, ""))
		data.hostInstPerf = s.perfOrLog(pc.query(ctx, hostRefs, hostInstancedCounters, "*"))
	}

	result := make([]*Status, len(hss))
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxHostWorkers)
	for i := range hss {
		wg.Add(1)
		sem <- struct{}{}
		go func(i int) {
			defer wg.Done()
			defer func() { <-sem }()
			// a panic in this goroutine is not covered by the recover in Collect and would kill the process
			defer func() {
				if r := recover(); r != nil {
					s.logger.Error(fmt.Sprintf("panic while collecting host %s: %v\n%s", hss[i].Self.Value, r, debug.Stack()))
				}
			}()
			h := hss[i]
			result[i] = s.hostStatus(ctx, c.Client, &h, dsByHost[h.Self.Value], vmsByHost[h.Self.Value], data)
		}(i)
	}
	wg.Wait()

	statuses := make([]*Status, 0, len(result))
	for _, st := range result {
		if st != nil {
			statuses = append(statuses, st)
		}
	}
	if len(statuses) == 0 {
		return nil, fmt.Errorf("no host of %s could be collected", s.vmwareHost)
	}
	if about := c.ServiceContent.About; about.ApiType == "VirtualCenter" {
		statuses[0].VCenter = &vcenterInfo{Version: about.Version, Build: about.Build, Patch: about.PatchLevel, FullName: about.FullName, ApiType: about.ApiType}
	}
	return statuses, nil
}

// hostStatus builds the Status of one host. Details that can not be read (e.g. the host is
// disconnected) are skipped and logged, the rest of the metrics is still reported.
func (s *service) hostStatus(ctx context.Context, c *vim25.Client, h *mo.HostSystem, dss []totalds, vms []mo.VirtualMachine, data *scrapeData) *Status {
	status := Status{
		DiskOk:           []diskOk{},
		NetworkPNICSpeed: []pnic{},
		SensorInfo:       []NumericSensorInfo{},
		StorageInfo:      []StorageStateInfo{},
		DS:               dss,
		VMS:              []hvms{},
		HostCluster:      data.hostCluster[h.Self.Value],
		HostPerf:         instanceSamples(data.hostInstPerf[h.Self.Value], data.dsInstances),
	}
	for _, m := range data.hostPerf[h.Self.Value] {
		if v, err := strconv.ParseFloat(m.MetricValue, 64); err == nil && m.Instance == "-" {
			status.HostPerf = append(status.HostPerf, perfSample{Counter: m.MetricName, Value: v})
		}
	}

	sum := h.Summary
	status.HostName = sum.Config.Name
	if status.HostName == "" {
		status.HostName = h.Self.Value
	}
	if p := sum.Config.Product; p != nil {
		status.Product.Name = p.Name
		status.Product.FullName = p.FullName
		status.Product.Vendor = p.Vendor
		status.Product.Version = p.Version
		status.Product.Build = p.Build
		status.Product.OsType = p.OsType
		status.Product.ApiVersion = p.ApiVersion
		status.Product.LicenseProductName = p.LicenseProductName
		status.Product.LicenseVersion = p.LicenseProductVersion
	}
	if hw := sum.Hardware; hw != nil {
		status.TotalCpu = float64(int64(hw.CpuMhz) * int64(hw.NumCpuCores))
		status.TotalMem = float64(hw.MemorySize)
		status.HW.Vendor = hw.Vendor
		status.HW.Model = hw.Model
		status.HW.Uuid = hw.Uuid
		status.HW.CpuModel = hw.CpuModel
		status.HW.NumCpuPkgs = float64(hw.NumCpuPkgs)
		status.HW.CpuMhz = float64(hw.CpuMhz)
		status.HW.NumCpuCores = float64(hw.NumCpuCores)
		status.HW.NumCpuThreads = float64(hw.NumCpuThreads)
		status.HW.NumNics = float64(hw.NumNics)
		status.HW.NumHBAs = float64(hw.NumHBAs)
	}
	status.UsageCpu = float64(sum.QuickStats.OverallCpuUsage)
	status.UsageMem = float64(sum.QuickStats.OverallMemoryUsage) * 1024 * 1024

	status.HostRedAlarms, status.HostYellowAlarms = alarmCounts(h.TriggeredAlarmState)

	connected := true
	if rt := sum.Runtime; rt != nil {
		status.HostPowerState = powerState(rt.PowerState)
		status.HostMaintenanceMode = maintenanceMode(rt.InMaintenanceMode)
		status.HostStandbyMode = standbyMode(rt.StandbyMode)
		if rt.BootTime != nil {
			status.HostBoot = float64(rt.BootTime.Unix())
		}
		connected = rt.ConnectionState == types.HostSystemConnectionStateConnected
		if connected {
			status.HostConnected = 1
		}
		// health data is optional: the host may not report it (monitoring service down, disconnected, etc.)
		if hsr := rt.HealthSystemRuntime; hsr != nil {
			if hsr.SystemHealthInfo != nil {
				for _, sensor := range hsr.SystemHealthInfo.NumericSensorInfo {
					value := float64(sensor.CurrentReading) * math.Pow(10, float64(sensor.UnitModifier))
					status.SensorInfo = append(status.SensorInfo, NumericSensorInfo{
						Name:           sensor.Name,
						HealthState:    sensorHealthOf(sensor.HealthState),
						Value:          value,
						CurrentReading: strconv.Itoa(int(value)),
						BaseUnits:      sensor.BaseUnits,
						SensorType:     sensor.SensorType,
						Id:             sensor.Id,
						SensorNumber:   strconv.Itoa(int(sensor.SensorNumber)),
					})
				}
			}
			if hsr.HardwareStatusInfo != nil {
				for _, storageSensor := range hsr.HardwareStatusInfo.StorageStatusInfo {
					if !checkSensorIfAppend(storageSensor.Name, status.StorageInfo) {
						status.StorageInfo = append(status.StorageInfo, StorageStateInfo{
							Name:   storageSensor.Name,
							Status: sensorHealthOf(storageSensor.Status),
						})
					}
				}
			}
		}
	}

	if connected {
		s.hostDevices(ctx, c, h, &status)
	} else {
		s.logger.Info(fmt.Sprintf("host %s is not connected, storage and network details are skipped", status.HostName))
	}

	// VM names are unique only per folder in vCenter, while a duplicated label set fails the whole scrape
	seen := make(map[string]struct{}, len(vms))
	for _, vm := range vms {
		st := vmStatus(vm, data.vmPerf[vm.Self.Value], data.vmDSPerf[vm.Self.Value], data.dsNames, data.dsInstances)
		if _, dup := seen[st.VmName]; dup {
			st.VmName = fmt.Sprintf("%s (%s)", st.VmName, vm.Self.Value)
		}
		seen[st.VmName] = struct{}{}
		status.VMS = append(status.VMS, st)
	}
	return &status
}

// hostDevices reads LUN states and physical NIC speeds through the host config managers.
func (s *service) hostDevices(ctx context.Context, c *vim25.Client, h *mo.HostSystem, status *Status) {
	hs := object.NewHostSystem(c, h.Self)

	ss, err := hs.ConfigManager().StorageSystem(ctx)
	if err != nil {
		s.logger.Error(fmt.Sprintf("host %s: get storage system: %v", status.HostName, err))
	} else {
		var hostss mo.HostStorageSystem
		if err := ss.Properties(ctx, ss.Reference(), []string{"storageDeviceInfo"}, &hostss); err != nil {
			s.logger.Error(fmt.Sprintf("host %s: get storage system properties: %v", status.HostName, err))
		} else if hostss.StorageDeviceInfo != nil {
			for _, e := range hostss.StorageDeviceInfo.ScsiLun {
				lun := e.GetScsiLun()
				ok := 1.0
				for _, st := range lun.OperationalState {
					if st != "ok" {
						ok = 0
						break
					}
				}
				status.DiskOk = append(status.DiskOk, diskOk{lun.DeviceName, ok})
			}
		}
	}

	nn, err := hs.ConfigManager().NetworkSystem(ctx)
	if err != nil {
		s.logger.Error(fmt.Sprintf("host %s: get network system: %v", status.HostName, err))
		return
	}
	var hostsn mo.HostNetworkSystem
	if err := nn.Properties(ctx, nn.Reference(), []string{"networkInfo"}, &hostsn); err != nil {
		s.logger.Error(fmt.Sprintf("host %s: get network system properties: %v", status.HostName, err))
		return
	}
	if hostsn.NetworkInfo == nil {
		return
	}
	for _, ni := range hostsn.NetworkInfo.Pnic {
		var lSpeed float64
		if ni.LinkSpeed != nil {
			lSpeed = float64(ni.LinkSpeed.SpeedMb)
		}
		status.NetworkPNICSpeed = append(status.NetworkPNICSpeed, pnic{ni.Device, ni.Mac, lSpeed})
	}
}

// vmStatus converts a VM summary and its performance counters; Guest and Storage may be absent
// (powered off, inaccessible or orphaned VMs).
func vmStatus(vm mo.VirtualMachine, perfMetrics, dsPerf []vmMetric, dsNames, dsInstances map[string]string) hvms {
	var vmPerfRes vmPerf
	r := reflect.ValueOf(&vmPerfRes).Elem()
	for _, v := range perfMetrics {
		value, err := strconv.Atoi(v.MetricValue)
		if err != nil {
			continue
		}
		metricName := strings.ToUpper(strings.Replace(v.MetricName, ".", "_", -1))
		if f := r.FieldByName(metricName); f.IsValid() && f.CanSet() && f.Kind() == reflect.Float64 {
			f.SetFloat(float64(value))
		}
	}

	res := hvms{
		VmName:       vm.Summary.Config.Name,
		VmPowerState: powerStateVM(vm.Summary.Runtime.PowerState),
		VmBoot:       convertTime(vm),
		VmCpuAval:    float64(vm.Summary.Runtime.MaxCpuUsage),
		VmCpuUsage:   float64(vm.Summary.QuickStats.OverallCpuUsage),
		VmNumCpu:     float64(vm.Summary.Config.NumCpu),
		VmMemAval:    float64(vm.Summary.Config.MemorySizeMB),
		VmMemUsage:   float64(vm.Summary.QuickStats.HostMemoryUsage),
		Perf:         vmPerfRes,
		DSPerf:       instanceSamples(dsPerf, dsInstances),
	}
	if vm.Snapshot != nil {
		res.Snapshots = flattenSnapshots(vm.Snapshot.RootSnapshotList, nil)
		res.SnapshotCount = float64(len(res.Snapshots))
	}
	if vm.Storage != nil {
		for _, u := range vm.Storage.PerDatastoreUsage {
			name := dsNames[u.Datastore.Value]
			if name == "" {
				name = u.Datastore.Value
			}
			res.DSUsage = append(res.DSUsage, vmDSUsage{name, float64(u.Committed), float64(u.Uncommitted), float64(u.Unshared)})
		}
	}
	res.RedAlarms, res.YellowAlarms = alarmCounts(vm.TriggeredAlarmState)
	res.Template = boolValue(vm.Summary.Config.Template)
	if vm.Guest != nil {
		seen := make(map[string]struct{}, len(vm.Guest.Disk))
		for _, d := range vm.Guest.Disk {
			if _, dup := seen[d.DiskPath]; dup {
				continue
			}
			seen[d.DiskPath] = struct{}{}
			res.GuestDisks = append(res.GuestDisks, guestDisk{d.DiskPath, float64(d.Capacity), float64(d.FreeSpace)})
		}
	}
	if g := vm.Summary.Guest; g != nil {
		res.GuestToolsRunning = boolValue(g.ToolsRunningStatus == string(types.VirtualMachineToolsRunningStatusGuestToolsRunning))
		res.VmGuestId = g.GuestId
		res.VmGuestToolsStatus = guestToolsStatus(string(g.ToolsStatus))
		res.VmGuestToolsVersion = guestToolsVersion(g.ToolsVersionStatus)
		res.VmGuestFullName = g.GuestFullName
		res.VmGuestIpAddr = g.IpAddress
	}
	if st := vm.Summary.Storage; st != nil {
		res.VmGuestStorageCommitted = float64(st.Committed)
		res.VmGuestStorageUnCommitted = float64(st.Uncommitted)
	}
	return res
}

// flattenSnapshots returns all snapshots of the (nested) snapshot tree.
func flattenSnapshots(tree []types.VirtualMachineSnapshotTree, acc []snapshotInfo) []snapshotInfo {
	for _, sn := range tree {
		acc = append(acc, snapshotInfo{Name: sn.Name, Id: sn.Id, Created: float64(sn.CreateTime.Unix())})
		acc = flattenSnapshots(sn.ChildSnapshotList, acc)
	}
	return acc
}

// retrieveView loads properties of all objects of the given type and always destroys the view,
// even when the scrape context is already expired.
func retrieveView(ctx context.Context, m *view.Manager, root types.ManagedObjectReference, kind string, props []string, dst interface{}) error {
	v, err := m.CreateContainerView(ctx, root, []string{kind}, true)
	if err != nil {
		return err
	}
	defer func() {
		dctx, cancel := context.WithTimeout(context.Background(), logoutTimeout)
		defer cancel()
		_ = v.Destroy(dctx)
	}()
	return v.Retrieve(ctx, []string{kind}, props, dst)
}

func checkSensorIfAppend(name string, sensors []StorageStateInfo) bool {
	for _, sensor := range sensors {
		if sensor.Name == name {
			return true
		}
	}
	return false
}
