package vmware

import (
	"fmt"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

// Performance counters read from the hosts and from the VMs (instanced ones come per NIC, datastore, ...).
// Only counters that return data on ESXi 8 and vCenter 8 are listed.
var (
	hostCounters = []string{
		"cpu.usage.average", "cpu.usagemhz.average", "cpu.demand.average", "cpu.latency.average",
		"cpu.ready.summation", "cpu.readiness.average", "cpu.costop.summation",
		"mem.usage.average", "mem.active.average", "mem.consumed.average", "mem.granted.average",
		"mem.shared.average", "mem.sysUsage.average", "mem.swapused.average", "mem.vmmemctl.average",
		"disk.maxTotalLatency.latest", "disk.usage.average", "power.power.average", "sys.uptime.latest",
	}
	hostInstancedCounters = []string{
		"net.bytesRx.average", "net.bytesTx.average", "net.packetsRx.summation", "net.packetsTx.summation",
		"net.errorsRx.summation", "net.errorsTx.summation", "net.droppedRx.summation", "net.droppedTx.summation",
		"datastore.read.average", "datastore.write.average",
		"datastore.numberReadAveraged.average", "datastore.numberWriteAveraged.average",
		"datastore.totalReadLatency.average", "datastore.totalWriteLatency.average",
		"disk.read.average", "disk.write.average",
	}
	vmInstancedCounters = []string{
		"datastore.read.average", "datastore.write.average",
		"datastore.numberReadAveraged.average", "datastore.numberWriteAveraged.average",
		"datastore.totalReadLatency.average", "datastore.totalWriteLatency.average",
	}
)

// extraDescs describes metrics that do not fit the fixed Collector fields.
type extraDescs struct {
	vcenterInfo       *prometheus.Desc
	hostConnected     *prometheus.Desc
	hostClusterInfo   *prometheus.Desc
	dsInfo            *prometheus.Desc
	dsAccessible      *prometheus.Desc
	dsUncommitted     *prometheus.Desc
	dsProvisioned     *prometheus.Desc
	vmSnapshotCount   *prometheus.Desc
	vmSnapshotCreated *prometheus.Desc
	vmDSCommitted     *prometheus.Desc
	vmDSUncommitted   *prometheus.Desc
	vmDSUnshared      *prometheus.Desc

	hostStandbyMode   *prometheus.Desc
	hostRedAlarms     *prometheus.Desc
	hostYellowAlarms  *prometheus.Desc
	hostSensorValue   *prometheus.Desc
	dsMaintenanceMode *prometheus.Desc
	dsVMs             *prometheus.Desc
	dsHosts           *prometheus.Desc
	dsRedAlarms       *prometheus.Desc
	dsYellowAlarms    *prometheus.Desc
	vmTemplate        *prometheus.Desc
	vmToolsRunning    *prometheus.Desc
	vmGuestDiskCap    *prometheus.Desc
	vmGuestDiskFree   *prometheus.Desc
	vmRedAlarms       *prometheus.Desc
	vmYellowAlarms    *prometheus.Desc

	hostPerf     map[string]*prometheus.Desc // counter -> desc, labels: host_name
	hostInstPerf map[string]*prometheus.Desc // labels: host_name, ds_name|nic|device
	vmDSPerf     map[string]*prometheus.Desc // labels: vm_name, host_name, ds_name
}

// perfMetricName turns a counter like "net.bytesRx.average" into "<prefix>_net_bytesrx_average".
func perfMetricName(prefix, counter string) string {
	return prefix + "_" + strings.ToLower(strings.ReplaceAll(counter, ".", "_"))
}

// instanceLabel names the label of an instanced counter after what the instance is. It must not be
// "instance": Prometheus uses that name for the scrape target and would rename it to exported_instance.
func instanceLabel(counter string) string {
	switch {
	case strings.HasPrefix(counter, "datastore."):
		return "ds_name"
	case strings.HasPrefix(counter, "net."):
		return "nic"
	}
	return "device"
}

// newPerfDescs creates one Desc per counter; instanced counters get one more label for the instance.
func newPerfDescs(prefix string, counters []string, labels []string, instanced bool) map[string]*prometheus.Desc {
	m := make(map[string]*prometheus.Desc, len(counters))
	for _, c := range counters {
		l := labels
		if instanced {
			l = append(append([]string{}, labels...), instanceLabel(c))
		}
		m[c] = prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", perfMetricName(prefix, c)),
			fmt.Sprintf("Vmware performance counter %s (real-time sample, vSphere units)", c),
			l, nil,
		)
	}
	return m
}

func newExtraDescs() *extraDescs {
	desc := func(sub, name, help string, labels ...string) *prometheus.Desc {
		return prometheus.NewDesc(prometheus.BuildFQName(namespace, sub, name), help, labels, nil)
	}
	return &extraDescs{
		vcenterInfo:       desc("vcenter", "info", "Vmware vCenter version info", "version", "build", "patch", "full_name", "api_type"),
		hostConnected:     desc("host", "connected", "Vmware Host connection state to vCenter, connected 1, other 0", "host_name"),
		hostClusterInfo:   desc("host", "cluster_info", "Vmware Host membership in a cluster", "host_name", "cluster"),
		dsInfo:            desc("datastore", "info", "Datastore type and url", "ds_name", "host_name", "type", "url"),
		dsAccessible:      desc("datastore", "accessible", "Datastore is accessible 1, not 0", "ds_name", "host_name"),
		dsUncommitted:     desc("datastore", "uncommitted_size", "Datastore uncommitted (thin provisioned, not yet allocated) space in bytes", "ds_name", "host_name"),
		dsProvisioned:     desc("datastore", "provisioned_size", "Datastore provisioned space in bytes: used + uncommitted", "ds_name", "host_name"),
		vmSnapshotCount:   desc("vm", "snapshot_count", "Number of VM snapshots", "vm_name", "host_name"),
		vmSnapshotCreated: desc("vm", "snapshot_created_timestamp_seconds", "VM snapshot creation time, unix timestamp", "vm_name", "host_name", "snapshot_name", "snapshot_id"),
		vmDSCommitted:     desc("vm", "datastore_committed_size", "VM space committed on a datastore in bytes", "vm_name", "host_name", "ds_name"),
		vmDSUncommitted:   desc("vm", "datastore_uncommitted_size", "VM space uncommitted on a datastore in bytes", "vm_name", "host_name", "ds_name"),
		vmDSUnshared:      desc("vm", "datastore_unshared_size", "VM unshared space on a datastore in bytes", "vm_name", "host_name", "ds_name"),

		hostStandbyMode:   desc("host", "standby_mode", "Vmware Host standby mode, none 0, entering 1, in 2, exiting 3", "host_name"),
		hostRedAlarms:     desc("host", "red_alarms", "Number of triggered red alarms of the host", "host_name"),
		hostYellowAlarms:  desc("host", "yellow_alarms", "Number of triggered yellow alarms of the host", "host_name"),
		hostSensorValue:   desc("host", "sensor_value", "Vmware Host numeric sensor reading (unit modifier applied, unit in base_units)", "host_name", "name", "sensor_type", "base_units", "id"),
		dsMaintenanceMode: desc("datastore", "maintenance_mode", "Datastore maintenance mode, normal 0, enteringMaintenance 1, inMaintenance 2", "ds_name", "host_name"),
		dsVMs:             desc("datastore", "vms", "Number of VMs on the datastore", "ds_name", "host_name"),
		dsHosts:           desc("datastore", "hosts", "Number of hosts the datastore is mounted on", "ds_name", "host_name"),
		dsRedAlarms:       desc("datastore", "red_alarms", "Number of triggered red alarms of the datastore", "ds_name", "host_name"),
		dsYellowAlarms:    desc("datastore", "yellow_alarms", "Number of triggered yellow alarms of the datastore", "ds_name", "host_name"),
		vmTemplate:        desc("vm", "template", "VM is a template 1, a regular VM 0", "vm_name", "host_name"),
		vmToolsRunning:    desc("vm", "guest_tools_running_status", "VMware Tools are running 1, not running 0", "vm_name", "host_name"),
		vmGuestDiskCap:    desc("vm", "guest_disk_capacity_size", "Guest file system capacity in bytes (needs VMware Tools)", "vm_name", "host_name", "partition"),
		vmGuestDiskFree:   desc("vm", "guest_disk_free_size", "Guest file system free space in bytes (needs VMware Tools)", "vm_name", "host_name", "partition"),
		vmRedAlarms:       desc("vm", "red_alarms", "Number of triggered red alarms of the VM", "vm_name", "host_name"),
		vmYellowAlarms:    desc("vm", "yellow_alarms", "Number of triggered yellow alarms of the VM", "vm_name", "host_name"),

		hostPerf:     newPerfDescs("host", hostCounters, []string{"host_name"}, false),
		hostInstPerf: newPerfDescs("host", hostInstancedCounters, []string{"host_name"}, true),
		vmDSPerf:     newPerfDescs("vm", vmInstancedCounters, []string{"vm_name", "host_name"}, true),
	}
}

func (e *extraDescs) describe(ch chan<- *prometheus.Desc) {
	for _, d := range []*prometheus.Desc{
		e.vcenterInfo, e.hostConnected, e.hostClusterInfo, e.dsInfo, e.dsAccessible, e.dsUncommitted, e.dsProvisioned,
		e.vmSnapshotCount, e.vmSnapshotCreated, e.vmDSCommitted, e.vmDSUncommitted, e.vmDSUnshared,
		e.hostStandbyMode, e.hostRedAlarms, e.hostYellowAlarms, e.hostSensorValue,
		e.dsMaintenanceMode, e.dsVMs, e.dsHosts, e.dsRedAlarms, e.dsYellowAlarms,
		e.vmTemplate, e.vmToolsRunning, e.vmGuestDiskCap, e.vmGuestDiskFree, e.vmRedAlarms, e.vmYellowAlarms,
	} {
		ch <- d
	}
	for _, m := range []map[string]*prometheus.Desc{e.hostPerf, e.hostInstPerf, e.vmDSPerf} {
		for _, d := range m {
			ch <- d
		}
	}
}

func boolValue(b bool) float64 {
	if b {
		return 1
	}
	return 0
}

// collect emits the metrics of one host: the host itself, its datastores and its VMs.
func (e *extraDescs) collect(ch chan<- prometheus.Metric, s *Status) {
	gauge := func(d *prometheus.Desc, v float64, labels ...string) {
		ch <- prometheus.MustNewConstMetric(d, prometheus.GaugeValue, v, labels...)
	}

	if s.VCenter != nil {
		gauge(e.vcenterInfo, 1, s.VCenter.Version, s.VCenter.Build, s.VCenter.Patch, s.VCenter.FullName, s.VCenter.ApiType)
	}

	gauge(e.hostConnected, s.HostConnected, s.HostName)
	if s.HostCluster != "" {
		gauge(e.hostClusterInfo, 1, s.HostName, s.HostCluster)
	}
	gauge(e.hostStandbyMode, s.HostStandbyMode, s.HostName)
	gauge(e.hostRedAlarms, s.HostRedAlarms, s.HostName)
	gauge(e.hostYellowAlarms, s.HostYellowAlarms, s.HostName)
	sensorSeen := make(map[string]struct{}, len(s.SensorInfo))
	for _, sn := range s.SensorInfo {
		// the label set must be unique, otherwise the whole scrape fails
		key := sn.Name + "\x00" + sn.SensorType + "\x00" + sn.BaseUnits + "\x00" + sn.Id
		if _, dup := sensorSeen[key]; dup {
			continue
		}
		sensorSeen[key] = struct{}{}
		gauge(e.hostSensorValue, sn.Value, s.HostName, sn.Name, sn.SensorType, sn.BaseUnits, sn.Id)
	}
	for _, p := range s.HostPerf {
		if p.Instance == "" {
			if d, ok := e.hostPerf[p.Counter]; ok {
				gauge(d, p.Value, s.HostName)
			}
		} else if d, ok := e.hostInstPerf[p.Counter]; ok {
			gauge(d, p.Value, s.HostName, p.Instance)
		}
	}

	for _, ds := range s.DS {
		gauge(e.dsInfo, 1, ds.dsname, s.HostName, ds.dsType, ds.url)
		gauge(e.dsAccessible, boolValue(ds.accessible), ds.dsname, s.HostName)
		gauge(e.dsUncommitted, ds.uncommitted, ds.dsname, s.HostName)
		gauge(e.dsProvisioned, ds.capacity-ds.freespace+ds.uncommitted, ds.dsname, s.HostName)
		gauge(e.dsMaintenanceMode, ds.maintenance, ds.dsname, s.HostName)
		gauge(e.dsVMs, ds.vms, ds.dsname, s.HostName)
		gauge(e.dsHosts, ds.hosts, ds.dsname, s.HostName)
		gauge(e.dsRedAlarms, ds.redAlarms, ds.dsname, s.HostName)
		gauge(e.dsYellowAlarms, ds.yellowAlarms, ds.dsname, s.HostName)
	}

	for _, vm := range s.VMS {
		gauge(e.vmSnapshotCount, vm.SnapshotCount, vm.VmName, s.HostName)
		gauge(e.vmTemplate, vm.Template, vm.VmName, s.HostName)
		gauge(e.vmToolsRunning, vm.GuestToolsRunning, vm.VmName, s.HostName)
		gauge(e.vmRedAlarms, vm.RedAlarms, vm.VmName, s.HostName)
		gauge(e.vmYellowAlarms, vm.YellowAlarms, vm.VmName, s.HostName)
		for _, d := range vm.GuestDisks {
			gauge(e.vmGuestDiskCap, d.Capacity, vm.VmName, s.HostName, d.Path)
			gauge(e.vmGuestDiskFree, d.Free, vm.VmName, s.HostName, d.Path)
		}
		// snapshot names are not unique: the id keeps the series distinct
		for _, sn := range vm.Snapshots {
			gauge(e.vmSnapshotCreated, sn.Created, vm.VmName, s.HostName, sn.Name, fmt.Sprint(sn.Id))
		}
		for _, u := range vm.DSUsage {
			gauge(e.vmDSCommitted, u.Committed, vm.VmName, s.HostName, u.DsName)
			gauge(e.vmDSUncommitted, u.Uncommitted, vm.VmName, s.HostName, u.DsName)
			gauge(e.vmDSUnshared, u.Unshared, vm.VmName, s.HostName, u.DsName)
		}
		for _, p := range vm.DSPerf {
			if d, ok := e.vmDSPerf[p.Counter]; ok {
				gauge(d, p.Value, vm.VmName, s.HostName, p.Instance)
			}
		}
	}
}
