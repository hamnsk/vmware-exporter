package vmware

import (
	"context"
	"flag"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
	"reflect"
	"strconv"
	"vmware-exporter/pkg/logging"
)

var _ Service = &service{}

type service struct {
	logger         *logging.Logger
	vmwareHost     string
	vmwareUser     string
	vmwarePassword string
}

type Service interface {
	hostMetrics()
	dsMetrics()
	vmMetrics()
}

var interval = flag.Int("i", 20, "Interval ID")

func NewService(l *logging.Logger, host string, config interface{}) Service {
	cfg := reflect.ValueOf(config).Elem()

	vmwareUser := cfg.FieldByName("VmwareUser").Interface().(string)
	if len(vmwareUser) == 0 {
		l.Info("set default user name")
		vmwareUser = "monitoring"
	}

	vmwarePass := cfg.FieldByName("VmwarePass").Interface().(string)
	if len(vmwarePass) == 0 {
		l.Info("set default password ")
		vmwarePass = "password"
	}

	return &service{
		logger:         l,
		vmwareHost:     host,
		vmwareUser:     vmwareUser,
		vmwarePassword: vmwarePass,
	}
}

func (s *service) dsMetrics() {
	ctx := context.Background()
	c, err := NewClient(ctx, s.vmwareHost, s.vmwareUser, s.vmwarePassword)
	if err != nil {
		s.logger.Fatal(err.Error())
	}
	defer c.Logout(ctx)
	m := view.NewManager(c.Client)
	v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"Datastore"}, true)
	if err != nil {
		s.logger.Fatal(err.Error())
	}
	defer v.Destroy(ctx)
	var dss []mo.Datastore
	err = v.Retrieve(ctx, []string{"Datastore"}, []string{"summary"}, &dss)
	if err != nil {
		s.logger.Fatal(err.Error())
	}
	for _, ds := range dss {
		dsname := ds.Summary.Name
		prometheusTotalDs.WithLabelValues(dsname, s.vmwareHost).Set(float64(ds.Summary.Capacity))
		prometheusUsageDs.WithLabelValues(dsname, s.vmwareHost).Set(float64(ds.Summary.FreeSpace))
	}
}

func (s *service) vmMetrics() {
	ctx := context.Background()
	c, err := NewClient(ctx, s.vmwareHost, s.vmwareUser, s.vmwarePassword)
	if err != nil {
		s.logger.Fatal(err.Error())
	}
	defer c.Logout(ctx)
	m := view.NewManager(c.Client)
	v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
	if err != nil {
		s.logger.Fatal(err.Error())
	}
	defer v.Destroy(ctx)
	var vms []mo.VirtualMachine
	err = v.Retrieve(ctx, []string{"VirtualMachine"}, []string{"summary"}, &vms)
	if err != nil {
		s.logger.Fatal(err.Error())
	}

	vmsRefs, err := v.Find(ctx, []string{"VirtualMachine"}, nil)
	if err != nil {
		s.logger.Fatal(err.Error())
	}

	vmMetrics := perfMon(ctx, c, s.logger, vmsRefs)

	for _, vm := range vms {
		vmNum := vm.GetManagedEntity().Self.Value
		vmname := vm.Summary.Config.Name

		metrics, ok := vmMetrics[vmNum]

		if ok {
			for _, v := range metrics {
				value, err := strconv.Atoi(v.MetricValue)
				if err == nil {
					prometheusVmPerfMon.WithLabelValues(
						vmname,
						s.vmwareHost,
						v.Instance,
						v.MetricName,
						v.MetricUnit,
					).Set(float64(value))
				}

			}
		}

		prometheusVmBoot.WithLabelValues(vmname, s.vmwareHost).Set(convertTime(vm))
		prometheusVmCpuAval.WithLabelValues(vmname, s.vmwareHost).Set(float64(vm.Summary.Runtime.MaxCpuUsage) * 1000 * 1000)
		prometheusVmCpuUsage.WithLabelValues(vmname, s.vmwareHost).Set(float64(vm.Summary.QuickStats.OverallCpuUsage) * 1000 * 1000)
		prometheusVmNumCpu.WithLabelValues(vmname, s.vmwareHost).Set(float64(vm.Summary.Config.NumCpu))
		prometheusVmMemAval.WithLabelValues(vmname, s.vmwareHost).Set(float64(vm.Summary.Config.MemorySizeMB))
		prometheusVmMemUsage.WithLabelValues(vmname, s.vmwareHost).Set(float64(vm.Summary.QuickStats.GuestMemoryUsage) * 1024 * 1024)

		prometheusVmGuestInfo.WithLabelValues(
			vmname,
			s.vmwareHost,
			vm.Summary.Guest.GuestId,
			vm.Summary.Guest.GuestFullName,
			vm.Summary.Guest.IpAddress,
		).Set(1.0)

		prometheusVmGuestStorageCommitted.WithLabelValues(vmname, s.vmwareHost).Set(float64(vm.Summary.Storage.Committed))
		prometheusVmGuestStorageUnCommitted.WithLabelValues(vmname, s.vmwareHost).Set(float64(vm.Summary.Storage.Uncommitted))

	}
}

func (s *service) hostMetrics() {
	ctx := context.Background()
	c, err := NewClient(ctx, s.vmwareHost, s.vmwareUser, s.vmwarePassword)
	if err != nil {
		s.logger.Fatal(err.Error())
	}
	defer c.Logout(ctx)
	m := view.NewManager(c.Client)
	v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"HostSystem"}, true)
	if err != nil {
		s.logger.Fatal(err.Error())
	}
	defer v.Destroy(ctx)
	var hss []mo.HostSystem
	err = v.Retrieve(ctx, []string{"HostSystem"}, []string{"summary"}, &hss)
	if err != nil {
		s.logger.Fatal(err.Error())
	}
	for _, hs := range hss {
		prometheusHostPowerState.WithLabelValues(s.vmwareHost).Set(powerState(hs.Summary.Runtime.PowerState))
		prometheusHostBoot.WithLabelValues(s.vmwareHost).Set(float64(hs.Summary.Runtime.BootTime.Unix()))
		prometheusTotalCpu.WithLabelValues(s.vmwareHost).Set(totalCpu(hs))
		prometheusUsageCpu.WithLabelValues(s.vmwareHost).Set(float64(hs.Summary.QuickStats.OverallCpuUsage))
		prometheusTotalMem.WithLabelValues(s.vmwareHost).Set(float64(hs.Summary.Hardware.MemorySize))
		prometheusUsageMem.WithLabelValues(s.vmwareHost).Set(float64(hs.Summary.QuickStats.OverallMemoryUsage) * 1024 * 1024)

		prometheusHostHardwareInfo.WithLabelValues(
			s.vmwareHost,
			hs.Summary.Hardware.Vendor,
			hs.Summary.Hardware.Model,
			hs.Summary.Hardware.Uuid,
			hs.Summary.Hardware.CpuModel,
			strconv.Itoa(int(hs.Summary.Hardware.NumCpuPkgs)),
			strconv.Itoa(int(hs.Summary.Hardware.CpuMhz)),
			strconv.Itoa(int(hs.Summary.Hardware.NumCpuCores)),
			strconv.Itoa(int(hs.Summary.Hardware.NumCpuThreads)),
			strconv.Itoa(int(hs.Summary.Hardware.NumNics)),
			strconv.Itoa(int(hs.Summary.Hardware.NumHBAs))).Set(1.0)

	}
	finder := find.NewFinder(c.Client)
	hs, err := finder.DefaultHostSystem(ctx)
	if err != nil {
		s.logger.Fatal(err.Error())
	}
	ss, err := hs.ConfigManager().StorageSystem(ctx)
	if err != nil {
		s.logger.Fatal(err.Error())
	}
	var hostss mo.HostStorageSystem
	err = ss.Properties(ctx, ss.Reference(), nil, &hostss)
	if err != nil {
		s.logger.Fatal(err.Error())
	}

	for _, e := range hostss.StorageDeviceInfo.ScsiLun {
		lun := e.GetScsiLun()
		ok := 1.0
		for _, s := range lun.OperationalState {
			if s != "ok" {
				ok = 0
				break
			}
		}
		prometheusDiskOk.WithLabelValues(s.vmwareHost, lun.DeviceName).Set(ok)
	}

	nn, err := hs.ConfigManager().NetworkSystem(ctx)
	if err != nil {
		s.logger.Fatal(err.Error())
	}
	var hostsn mo.HostNetworkSystem
	err = nn.Properties(ctx, nn.Reference(), nil, &hostsn)
	if err != nil {
		s.logger.Fatal(err.Error())
	}
	for _, ni := range hostsn.NetworkInfo.Pnic {
		prometheusNetworkPNICSpeed.WithLabelValues(s.vmwareHost, ni.Device, ni.Mac).Set(float64(ni.LinkSpeed.SpeedMb))
	}
}
