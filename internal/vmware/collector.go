package vmware

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

type Collector struct {
	ss                        Service
	HostPowerState            *prometheus.Desc
	HostMaintenanceMode       *prometheus.Desc
	HostBoot                  *prometheus.Desc
	TotalCpu                  *prometheus.Desc
	UsageCpu                  *prometheus.Desc
	TotalMem                  *prometheus.Desc
	UsageMem                  *prometheus.Desc
	DiskOk                    *prometheus.Desc
	NetworkPNICSpeed          *prometheus.Desc
	HostHardwareInfo          *prometheus.Desc
	NumCpuPkgs                *prometheus.Desc
	CpuMhz                    *prometheus.Desc
	NumCpuCores               *prometheus.Desc
	NumCpuThreads             *prometheus.Desc
	NumNics                   *prometheus.Desc
	NumHBAs                   *prometheus.Desc
	TotalDs                   *prometheus.Desc
	UsageDs                   *prometheus.Desc
	VmGuestInfo               *prometheus.Desc
	VmGuestToolsStatus        *prometheus.Desc
	VmGuestToolsVersion       *prometheus.Desc
	VmGuestPowerState         *prometheus.Desc
	VmGuestStorageCommitted   *prometheus.Desc
	VmGuestStorageUnCommitted *prometheus.Desc
	VmBoot                    *prometheus.Desc
	VmCpuAval                 *prometheus.Desc
	VmCpuUsage                *prometheus.Desc
	VmNumCpu                  *prometheus.Desc
	VmMemAval                 *prometheus.Desc
	VmMemUsage                *prometheus.Desc
	//VM Perfomance Metrics Start Here
	CpuCostopSummation              *prometheus.Desc
	CpuDemandAverage                *prometheus.Desc
	CpuDemandentitlementratioLatest *prometheus.Desc
	CpuEntitlementLatest            *prometheus.Desc
	CpuIdleSummation                *prometheus.Desc
	CpuLatencyAverage               *prometheus.Desc
	CpuMaxlimitedSummation          *prometheus.Desc
	CpuOverlapSummation             *prometheus.Desc
	CpuReadinessAverage             *prometheus.Desc
	CpuReadySummation               *prometheus.Desc
	CpuRunSummation                 *prometheus.Desc
	CpuSwapwaitSummation            *prometheus.Desc
	CpuSystemSummation              *prometheus.Desc
	CpuUsageAverage                 *prometheus.Desc
	CpuUsageMaximum                 *prometheus.Desc
	CpuUsageMinimum                 *prometheus.Desc
	CpuUsageNone                    *prometheus.Desc
	CpuUsagemhzAverage              *prometheus.Desc
	CpuUsagemhzMaximum              *prometheus.Desc
	CpuUsagemhzMinimum              *prometheus.Desc
	CpuUsagemhzNone                 *prometheus.Desc
	CpuUsedSummation                *prometheus.Desc
	CpuWaitSummation                *prometheus.Desc
	DatastoreMaxtotallatencyLatest  *prometheus.Desc
	DiskMaxtotallatencyLatest       *prometheus.Desc
	DiskReadAverage                 *prometheus.Desc
	DiskUsageAverage                *prometheus.Desc
	DiskUsageMaximum                *prometheus.Desc
	DiskUsageMinimum                *prometheus.Desc
	DiskUsageNone                   *prometheus.Desc
	DiskWriteAverage                *prometheus.Desc
	MemActiveAverage                *prometheus.Desc
	MemActiveMaximum                *prometheus.Desc
	MemActiveMinimum                *prometheus.Desc
	MemActiveNone                   *prometheus.Desc
	MemActivewriteAverage           *prometheus.Desc
	MemCompressedAverage            *prometheus.Desc
	MemCompressionrateAverage       *prometheus.Desc
	MemConsumedAverage              *prometheus.Desc
	MemConsumedMaximum              *prometheus.Desc
	MemConsumedMinimum              *prometheus.Desc
	MemConsumedNone                 *prometheus.Desc
	MemDecompressionrateAverage     *prometheus.Desc
	MemEntitlementAverage           *prometheus.Desc
	MemGrantedAverage               *prometheus.Desc
	MemGrantedMaximum               *prometheus.Desc
	MemGrantedMinimum               *prometheus.Desc
	MemGrantedNone                  *prometheus.Desc
	MemLatencyAverage               *prometheus.Desc
	MemLlswapinrateAverage          *prometheus.Desc
	MemLlswapoutrateAverage         *prometheus.Desc
	MemLlswapusedAverage            *prometheus.Desc
	MemLlswapusedMaximum            *prometheus.Desc
	MemLlswapusedMinimum            *prometheus.Desc
	MemLlswapusedNone               *prometheus.Desc
	MemOverheadAverage              *prometheus.Desc
	MemOverheadMaximum              *prometheus.Desc
	MemOverheadMinimum              *prometheus.Desc
	MemOverheadNone                 *prometheus.Desc
	MemOverheadmaxAverage           *prometheus.Desc
	MemOverheadtouchedAverage       *prometheus.Desc
	MemSharedAverage                *prometheus.Desc
	MemSharedMaximum                *prometheus.Desc
	MemSharedMinimum                *prometheus.Desc
	MemSharedNone                   *prometheus.Desc
	MemSwapinAverage                *prometheus.Desc
	MemSwapinMaximum                *prometheus.Desc
	MemSwapinMinimum                *prometheus.Desc
	MemSwapinNone                   *prometheus.Desc
	MemSwapinrateAverage            *prometheus.Desc
	MemSwapoutAverage               *prometheus.Desc
	MemSwapoutMaximum               *prometheus.Desc
	MemSwapoutMinimum               *prometheus.Desc
	MemSwapoutNone                  *prometheus.Desc
	MemSwapoutrateAverage           *prometheus.Desc
	MemSwappedAverage               *prometheus.Desc
	MemSwappedMaximum               *prometheus.Desc
	MemSwappedMinimum               *prometheus.Desc
	MemSwappedNone                  *prometheus.Desc
	MemSwaptargetAverage            *prometheus.Desc
	MemSwaptargetMaximum            *prometheus.Desc
	MemSwaptargetMinimum            *prometheus.Desc
	MemSwaptargetNone               *prometheus.Desc
	MemUsageAverage                 *prometheus.Desc
	MemUsageMaximum                 *prometheus.Desc
	MemUsageMinimum                 *prometheus.Desc
	MemUsageNone                    *prometheus.Desc
	MemVmmemctlAverage              *prometheus.Desc
	MemVmmemctlMaximum              *prometheus.Desc
	MemVmmemctlMinimum              *prometheus.Desc
	MemVmmemctlNone                 *prometheus.Desc
	MemVmmemctltargetAverage        *prometheus.Desc
	MemVmmemctltargetMaximum        *prometheus.Desc
	MemVmmemctltargetMinimum        *prometheus.Desc
	MemVmmemctltargetNone           *prometheus.Desc
	MemZeroAverage                  *prometheus.Desc
	MemZeroMaximum                  *prometheus.Desc
	MemZeroMinimum                  *prometheus.Desc
	MemZeroNone                     *prometheus.Desc
	MemZippedLatest                 *prometheus.Desc
	MemZipsavedLatest               *prometheus.Desc
	NetBroadcastrxSummation         *prometheus.Desc
	NetBroadcasttxSummation         *prometheus.Desc
	NetBytesrxAverage               *prometheus.Desc
	NetBytestxAverage               *prometheus.Desc
	NetDroppedrxSummation           *prometheus.Desc
	NetDroppedtxSummation           *prometheus.Desc
	NetMulticastrxSummation         *prometheus.Desc
	NetMulticasttxSummation         *prometheus.Desc
	NetPacketsrxSummation           *prometheus.Desc
	NetPacketstxSummation           *prometheus.Desc
	NetPnicbytesrxAverage           *prometheus.Desc
	NetPnicbytestxAverage           *prometheus.Desc
	NetReceivedAverage              *prometheus.Desc
	NetTransmittedAverage           *prometheus.Desc
	NetUsageAverage                 *prometheus.Desc
	NetUsageMaximum                 *prometheus.Desc
	NetUsageMinimum                 *prometheus.Desc
	NetUsageNone                    *prometheus.Desc
	PowerEnergySummation            *prometheus.Desc
	PowerPowerAverage               *prometheus.Desc
	RescpuActav15Latest             *prometheus.Desc
	RescpuActav1Latest              *prometheus.Desc
	RescpuActav5Latest              *prometheus.Desc
	RescpuActpk15Latest             *prometheus.Desc
	RescpuActpk1Latest              *prometheus.Desc
	RescpuActpk5Latest              *prometheus.Desc
	RescpuMaxlimited15Latest        *prometheus.Desc
	RescpuMaxlimited1Latest         *prometheus.Desc
	RescpuMaxlimited5Latest         *prometheus.Desc
	RescpuRunav15Latest             *prometheus.Desc
	RescpuRunav1Latest              *prometheus.Desc
	RescpuRunav5Latest              *prometheus.Desc
	RescpuRunpk15Latest             *prometheus.Desc
	RescpuRunpk1Latest              *prometheus.Desc
	RescpuRunpk5Latest              *prometheus.Desc
	RescpuSamplecountLatest         *prometheus.Desc
	RescpuSampleperiodLatest        *prometheus.Desc
	SysHeartbeatLatest              *prometheus.Desc
	SysOsuptimeLatest               *prometheus.Desc
	SysUptimeLatest                 *prometheus.Desc
	VirtualdiskReadAverage          *prometheus.Desc
	VirtualdiskWriteAverage         *prometheus.Desc
}

var _ prometheus.Collector = &Collector{}

//New Create A new VMWare collector
func NewCollector(ss Service) *Collector {
	return &Collector{
		ss: ss,
		HostPowerState: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "host", "power_state"),
			"poweredOn 1, poweredOff 2, standBy 3, other 0",
			[]string{"host_name"},
			nil,
		),
		HostMaintenanceMode: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "host", "maintenance_mode"),
			"Maintenance mode 1 true, 0 false",
			[]string{"host_name"},
			nil,
		),

		HostBoot: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "host", "boot_timestamp_seconds"),
			"Uptime host",
			[]string{"host_name"},
			nil,
		),
		TotalCpu: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "host", "cpu_max"),
			"CPU total",
			[]string{"host_name"},
			nil,
		),
		UsageCpu: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "host", "cpu_usage"),
			"CPU Usage",
			[]string{"host_name"},
			nil,
		),
		TotalMem: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "host", "memory_max"),
			"Memory max",
			[]string{"host_name"},
			nil,
		),
		UsageMem: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "host", "memory_usage"),
			"Memory Usage",
			[]string{"host_name"},
			nil,
		),
		DiskOk: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "host", "disk_ok"),
			"Disk is working normally",
			[]string{"host_name", "device"},
			nil,
		),
		NetworkPNICSpeed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "host", "network_pnic_speed"),
			"Network PNIC Speed",
			[]string{"host_name", "device", "mac"},
			nil,
		),
		HostHardwareInfo: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "host", "hardware_info"),
			"Vmware Host Hardware info",
			[]string{"host_name", "vendor", "model", "uuid", "cpu_model"},
			nil,
		),
		NumCpuPkgs: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "host", "num_cpu_pkgs"),
			"Vmware Host Number CPU Pkgs",
			[]string{"host_name"},
			nil,
		),
		CpuMhz: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "host", "cpu_mhz"),
			"Vmware Host CPU Mhz",
			[]string{"host_name"},
			nil,
		),
		NumCpuCores: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "host", "cpu_cores"),
			"Vmware Host CPU Cores",
			[]string{"host_name"},
			nil,
		),
		NumCpuThreads: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "host", "cpu_threads"),
			"Vmware Host CPU Threads",
			[]string{"host_name"},
			nil,
		),
		NumNics: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "host", "num_nics"),
			"Vmware Host Number of NIC's",
			[]string{"host_name"},
			nil,
		),
		NumHBAs: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "host", "num_hbas"),
			"Vmware Host Number of HBA's",
			[]string{"host_name"},
			nil,
		),
		TotalDs: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "datastore", "capacity_size"),
			"Datastore total",
			[]string{"ds_name", "host_name"},
			nil,
		),
		UsageDs: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "datastore", "freespace_size"),
			"Datastore free",
			[]string{"ds_name", "host_name"},
			nil,
		),
		VmGuestInfo: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "guest_info"),
			"VMWare Guest VM info",
			[]string{"vm_name", "host_name", "guest_id", "guest_full_name", "ip_addr"},
			nil,
		),
		VmGuestPowerState: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "power_state"),
			"VMWare Guest Power State info",
			[]string{"vm_name", "host_name"},
			nil,
		),
		VmGuestToolsStatus: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "guest_tools_status"),
			"VMWare Guest Tools Status info",
			[]string{"vm_name", "host_name"},
			nil,
		),
		VmGuestToolsVersion: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "guest_tools_version"),
			"VMWare Guest Tools Status info",
			[]string{"vm_name", "host_name"},
			nil,
		),
		VmGuestStorageCommitted: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "guest_storage_committed"),
			"VMWare Guest VM storage committed",
			[]string{"vm_name", "host_name"},
			nil,
		),
		VmGuestStorageUnCommitted: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "guest_storage_uncommitted"),
			"VMWare Guest VM storage uncommitted",
			[]string{"vm_name", "host_name"},
			nil,
		),
		VmBoot: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "boot_timestamp_seconds"),
			"VMWare VM boot time in seconds",
			[]string{"vm_name", "host_name"},
			nil,
		),
		VmCpuAval: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "cpu_avaleblemhz"),
			"VMWare VM usage CPU",
			[]string{"vm_name", "host_name"},
			nil,
		),
		VmCpuUsage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "cpu_usagemhz"),
			"VMWare VM usage CPU",
			[]string{"vm_name", "host_name"},
			nil,
		),
		VmNumCpu: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "num_cpu"),
			"Available number of cores",
			[]string{"vm_name", "host_name"},
			nil,
		),
		VmMemAval: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_avaleble"),
			"Available memory",
			[]string{"vm_name", "host_name"},
			nil,
		),
		VmMemUsage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_usage"),
			"Usage memory",
			[]string{"vm_name", "host_name"},
			nil,
		),
		//VM Performance Metrics
		CpuCostopSummation: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "cpu_costop_summation"),
			"Time the virtual machine is ready to run, but is unable to run due to co-scheduling constraints in millisecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
		CpuDemandAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "cpu_demand_average"),
			"The amount of CPU resources a virtual machine would use if there were no CPU contention or CPU limit in Mhz.",
			[]string{"vm_name", "host_name"},
			nil,
		),
		CpuDemandentitlementratioLatest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "cpu_demandentitlementratio_latest"),
			"CPU resource entitlement to CPU demand ratio (in percents)",
			[]string{"vm_name", "host_name"},
			nil,
		),
		CpuEntitlementLatest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "cpu_entitlement_latest"),
			"CPU resources devoted by the ESXi scheduler in Mhz.",
			[]string{"vm_name", "host_name"},
			nil,
		),
		CpuIdleSummation: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "cpu_idle_summation"),
			"Total time that the CPU spent in an idle state in percent.",
			[]string{"vm_name", "host_name"},
			nil,
		),
		CpuLatencyAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "cpu_latency_average"),
			"Percent of time the virtual machine is unable to run because it is contending for access to the physical CPU(s)",
			[]string{"vm_name", "host_name"},
			nil,
		),
		CpuMaxlimitedSummation: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "cpu_maxlimited_summation"),
			"Time the virtual machine is ready to run, but is not running because it has reached its maximum CPU limit setting in milliseconds.",
			[]string{"vm_name", "host_name"},
			nil,
		),
		CpuOverlapSummation: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "cpu_overlap_summation"),
			"Time the virtual machine was interrupted to perform system services on behalf of itself or other virtual machines in milliseconds.",
			[]string{"vm_name", "host_name"},
			nil,
		),
		CpuReadinessAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "cpu_readiness_average"),
			"Percentage of time that the virtual machine was ready, but could not get scheduled to run on the physical CPU.",
			[]string{"vm_name", "host_name"},
			nil,
		),
		CpuReadySummation: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "cpu_ready_summation"),
			"Time that the virtual machine was ready, but could not get scheduled to run on the physical CPU during last measurement interval. CPU ready time is dependent on the number of virtual machines on the host and their CPU loads.",
			[]string{"vm_name", "host_name"},
			nil,
		),
		CpuRunSummation: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "cpu_run_summation"),
			"Time the virtual machine is scheduled to run in millisecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
		CpuSwapwaitSummation: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "cpu_swapwait_summation"),
			"CPU time spent waiting for swap-in in millisecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
		CpuSystemSummation: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "cpu_system_summation"),
			"Amount of time spent on system processes on each virtual CPU in the virtual machine. This is the host view of the CPU usage, not the guest operating system view. Msec",
			[]string{"vm_name", "host_name"},
			nil,
		),
		CpuUsageAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "cpu_usage_average"),
			"CPU Usage in percent",
			[]string{"vm_name", "host_name"},
			nil,
		),
		CpuUsageMaximum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "cpu_usage_maximum"),
			"CPU Usage in percent",
			[]string{"vm_name", "host_name"},
			nil,
		),
		CpuUsageMinimum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "cpu_usage_minimum"),
			"CPU Usage in percent",
			[]string{"vm_name", "host_name"},
			nil,
		),
		CpuUsageNone: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "cpu_usage_none"),
			"CPU Usage in percent",
			[]string{"vm_name", "host_name"},
			nil,
		),
		CpuUsagemhzAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "cpu_usagemhz_average"),
			"CPU Usage in Mhz",
			[]string{"vm_name", "host_name"},
			nil,
		),
		CpuUsagemhzMaximum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "cpu_usagemhz_maximum"),
			"CPU Usage in Mhz",
			[]string{"vm_name", "host_name"},
			nil,
		),
		CpuUsagemhzMinimum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "cpu_usagemhz_minimum"),
			"CPU Usage in Mhz",
			[]string{"vm_name", "host_name"},
			nil,
		),
		CpuUsagemhzNone: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "cpu_usagemhz_none"),
			"CPU Usage in Mhz",
			[]string{"vm_name", "host_name"},
			nil,
		),
		CpuUsedSummation: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "cpu_used_summation"),
			"Time accounted to the virtual machine. If a system service runs on behalf of this virtual machine, the time spent by that service (represented by cpu.system) should be charged to this virtual machine. If not, the time spent (represented by cpu.overlap) should not be charged against this virtual machine.",
			[]string{"vm_name", "host_name"},
			nil,
		),
		CpuWaitSummation: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "cpu_wait_summation"),
			"Total CPU time spent in wait state.The wait total includes time spent the CPU Idle, CPU Swap Wait, and CPU I/O Wait states.",
			[]string{"vm_name", "host_name"},
			nil,
		),
		DatastoreMaxtotallatencyLatest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "datastore_maxtotallatency_latest"),
			"Highest latency value across all datastores used by the host in milliseconds",
			[]string{"vm_name", "host_name"},
			nil,
		),
		DiskMaxtotallatencyLatest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "disk_maxtotallatency_latest"),
			"Highest latency value across all disks used by the host. Latency measures the time taken to process a SCSI command issued by the guest OS to the virtual machine. The kernel latency is the time VMkernel takes to process an IO request. The device latency is the time it takes the hardware to handle the request. Msec",
			[]string{"vm_name", "host_name"},
			nil,
		),
		DiskReadAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "disk_read_average"),
			"Average number of kilobytes read from the disk each second during the collection interval. kiloBytesPerSecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
		DiskUsageAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "disk_usage_average"),
			"Aggregated disk I/O rate. For hosts, this metric includes the rates for all virtual machines running on the host during the collection interval. kiloBytesPerSecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
		DiskUsageMaximum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "disk_usage_maximum"),
			"Aggregated disk I/O rate. For hosts, this metric includes the rates for all virtual machines running on the host during the collection interval. kiloBytesPerSecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
		DiskUsageMinimum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "disk_usage_minimum"),
			"Aggregated disk I/O rate. For hosts, this metric includes the rates for all virtual machines running on the host during the collection interval. kiloBytesPerSecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
		DiskUsageNone: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "disk_usage_none"),
			"Aggregated disk I/O rate. For hosts, this metric includes the rates for all virtual machines running on the host during the collection interval. kiloBytesPerSecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
		DiskWriteAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "disk_write_average"),
			"Average number of kilobytes written to disk each second during the collection interval. kiloBytesPerSecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemActiveAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_active_average"),
			"Amount of memory that is actively used, as estimated by VMkernel based on recently touched memory pages. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemActiveMaximum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_active_maximum"),
			"Amount of memory that is actively used, as estimated by VMkernel based on recently touched memory pages. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemActiveMinimum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_active_minimum"),
			"Amount of memory that is actively used, as estimated by VMkernel based on recently touched memory pages. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemActiveNone: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_active_none"),
			"Amount of memory that is actively used, as estimated by VMkernel based on recently touched memory pages. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemActivewriteAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_activewrite_average"),
			"Estimate for the amount of memory actively being written to by the virtual machine. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemCompressedAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_compressed_average"),
			"Amount of memory reserved by userworlds. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemCompressionrateAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_compressionrate_average"),
			"Rate of memory compression for the virtual machine. kiloBytesPerSecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemConsumedAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_consumed_average"),
			"Amount of host physical memory consumed by a virtual machine, host, or cluster. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemConsumedMaximum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_consumed_maximum"),
			"Amount of host physical memory consumed by a virtual machine, host, or cluster. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemConsumedMinimum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_consumed_minimum"),
			"Amount of host physical memory consumed by a virtual machine, host, or cluster. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemConsumedNone: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_consumed_none"),
			"Amount of host physical memory consumed by a virtual machine, host, or cluster. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemDecompressionrateAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_decompressionrate_average"),
			"Rate of memory decompression for the virtual machine. kiloBytesPerSecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemEntitlementAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_entitlement_average"),
			"Amount of host physical memory the virtual machine is entitled to, as determined by the ESX scheduler. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemGrantedAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_granted_average"),
			"Amount of host physical memory or physical memory that is mapped for a virtual machine or a host. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemGrantedMaximum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_granted_maximum"),
			"Amount of host physical memory or physical memory that is mapped for a virtual machine or a host. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemGrantedMinimum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_granted_minimum"),
			"Amount of host physical memory or physical memory that is mapped for a virtual machine or a host. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemGrantedNone: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_granted_none"),
			"Amount of host physical memory or physical memory that is mapped for a virtual machine or a host. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemLatencyAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_latency_average"),
			"Percentage of time the virtual machine is waiting to access swapped or compressed memory",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemLlswapinrateAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_llswapinrate_average"),
			"Rate at which memory is being swapped from host cache into active memory in kiloBytesPerSecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemLlswapoutrateAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_llswapoutrate_average"),
			"Rate at which memory is being swapped from active memory to host cache in kiloBytesPerSecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemLlswapusedAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_llswapused_maximum"),
			"Space used for caching swapped pages in the host cache KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemLlswapusedMaximum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_llswapused_average"),
			"Space used for caching swapped pages in the host cache KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemLlswapusedMinimum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_llswapused_minimum"),
			"Space used for caching swapped pages in the host cache KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemLlswapusedNone: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_llswapused_none"),
			"Space used for caching swapped pages in the host cache KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemOverheadAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_overhead_average"),
			"Host physical memory (KB) consumed by the virtualization infrastructure for running the virtual machine.",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemOverheadMaximum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_overhead_maximum"),
			"Host physical memory (KB) consumed by the virtualization infrastructure for running the virtual machine.",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemOverheadMinimum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_overhead_minimum"),
			"Host physical memory (KB) consumed by the virtualization infrastructure for running the virtual machine.",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemOverheadNone: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_overhead_none"),
			"Host physical memory (KB) consumed by the virtualization infrastructure for running the virtual machine.",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemOverheadmaxAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_overheadmax_average"),
			"Host physical memory (KB) reserved for use as the virtualization overhead for the virtual machine",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemOverheadtouchedAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_overheadtouched_average"),
			"Actively touched overhead host physical memory (KB) reserved for use as the virtualization overhead for the virtual machine",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemSharedAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_shared_average"),
			"Amount of guest physical memory that is shared with other virtual machines, relative to a single virtual machine or to all powered-on virtual machines on a host. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemSharedMaximum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_shared_maximum"),
			"Amount of guest physical memory that is shared with other virtual machines, relative to a single virtual machine or to all powered-on virtual machines on a host. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemSharedMinimum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_shared_minimum"),
			"Amount of guest physical memory that is shared with other virtual machines, relative to a single virtual machine or to all powered-on virtual machines on a host. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemSharedNone: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_shared_none"),
			"Amount of guest physical memory that is shared with other virtual machines, relative to a single virtual machine or to all powered-on virtual machines on a host. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemSwapinAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_swapin_average"),
			"Amount swapped-in to memory from disk. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemSwapinMaximum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_swapin_maximum"),
			"Amount swapped-in to memory from disk. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemSwapinMinimum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_swapin_minimum"),
			"Amount swapped-in to memory from disk. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemSwapinNone: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_swapin_none"),
			"Amount swapped-in to memory from disk. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemSwapinrateAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_swapinrate_average"),
			"Rate at which memory is swapped from disk into active memory during the interval. This counter applies to virtual machines and is generally more useful than the swapin counter to determine if the virtual machine is running slow due to swapping, especially when looking at real-time statistics. kiloBytesPerSecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemSwapoutAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_swapout_average"),
			"Amount of memory swapped-out to disk. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemSwapoutMaximum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_swapout_maximum"),
			"Amount of memory swapped-out to disk. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemSwapoutMinimum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_swapout_minimum"),
			"Amount of memory swapped-out to disk. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemSwapoutNone: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_swapout_none"),
			"Amount of memory swapped-out to disk. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemSwapoutrateAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_swapoutrate_average"),
			"Rate at which memory is being swapped from active memory to disk during the current interval. This counter applies to virtual machines and is generally more useful than the swapout counter to determine if the virtual machine is running slow due to swapping, especially when looking at real-time statistics. kiloBytesPerSecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemSwappedAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_swapped_average"),
			"Current amount of guest physical memory swapped out to the virtual machine swap file by the VMkernel. Swapped memory stays on disk until the virtual machine needs it. This statistic refers to VMkernel swapping and not to guest OS swapping. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemSwappedMaximum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_swapped_maximum"),
			"Current amount of guest physical memory swapped out to the virtual machine swap file by the VMkernel. Swapped memory stays on disk until the virtual machine needs it. This statistic refers to VMkernel swapping and not to guest OS swapping. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemSwappedMinimum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_swapped_minimum"),
			"Current amount of guest physical memory swapped out to the virtual machine swap file by the VMkernel. Swapped memory stays on disk until the virtual machine needs it. This statistic refers to VMkernel swapping and not to guest OS swapping. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemSwappedNone: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_swapped_none"),
			"Current amount of guest physical memory swapped out to the virtual machine swap file by the VMkernel. Swapped memory stays on disk until the virtual machine needs it. This statistic refers to VMkernel swapping and not to guest OS swapping. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemSwaptargetAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_swaptarget_average"),
			"Target size for the virtual machine swap file. The VMkernel manages swapping by comparing swaptarget against swapped. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemSwaptargetMaximum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_swaptarget_maximum"),
			"Target size for the virtual machine swap file. The VMkernel manages swapping by comparing swaptarget against swapped. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemSwaptargetMinimum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_swaptarget_minimum"),
			"Target size for the virtual machine swap file. The VMkernel manages swapping by comparing swaptarget against swapped. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemSwaptargetNone: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_swaptarget_none"),
			"Target size for the virtual machine swap file. The VMkernel manages swapping by comparing swaptarget against swapped. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemUsageAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_usage_average"),
			"Memory usage as percent of total configured or available memory. A value between 0 and 10,000 expressed as a hundredth of a percent (1 = 0.01%).",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemUsageMaximum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_usage_maximum"),
			"Memory usage as percent of total configured or available memory. A value between 0 and 10,000 expressed as a hundredth of a percent (1 = 0.01%).",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemUsageMinimum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_usage_minimum"),
			"Memory usage as percent of total configured or available memory. A value between 0 and 10,000 expressed as a hundredth of a percent (1 = 0.01%).",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemUsageNone: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_usage_none"),
			"Memory usage as percent of total configured or available memory. A value between 0 and 10,000 expressed as a hundredth of a percent (1 = 0.01%).",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemVmmemctlAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_vmmemctl_average"),
			"Amount of memory allocated by the virtual machine memory control driver (vmmemctl), which is installed with VMware Tools. It is a VMware exclusive memory-management driver that controls ballooning. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemVmmemctlMaximum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_vmmemctl_maximum"),
			"Amount of memory allocated by the virtual machine memory control driver (vmmemctl), which is installed with VMware Tools. It is a VMware exclusive memory-management driver that controls ballooning. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemVmmemctlMinimum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_vmmemctl_minimum"),
			"Amount of memory allocated by the virtual machine memory control driver (vmmemctl), which is installed with VMware Tools. It is a VMware exclusive memory-management driver that controls ballooning. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemVmmemctlNone: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_vmmemctl_none"),
			"Amount of memory allocated by the virtual machine memory control driver (vmmemctl), which is installed with VMware Tools. It is a VMware exclusive memory-management driver that controls ballooning. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemVmmemctltargetAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_vmmemctltarget_average"),
			"Target value set by VMkernal for the virtual machine's memory balloon size. In conjunction with vmmemctl metric, this metric is used by VMkernel to inflate and deflate the balloon for a virtual machine KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemVmmemctltargetMaximum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_vmmemctltarget_maximum"),
			"Target value set by VMkernal for the virtual machine's memory balloon size. In conjunction with vmmemctl metric, this metric is used by VMkernel to inflate and deflate the balloon for a virtual machine KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemVmmemctltargetMinimum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_vmmemctltarget_minimum"),
			"Target value set by VMkernal for the virtual machine's memory balloon size. In conjunction with vmmemctl metric, this metric is used by VMkernel to inflate and deflate the balloon for a virtual machine KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemVmmemctltargetNone: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_vmmemctltarget_none"),
			"Target value set by VMkernal for the virtual machine's memory balloon size. In conjunction with vmmemctl metric, this metric is used by VMkernel to inflate and deflate the balloon for a virtual machine KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemZeroAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_zero_average"),
			"Memory that contains 0s only.Included in shared amount. Through transparent page sharing, zero memory pages can be shared among virtual machines that run the same operating system. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemZeroMaximum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_zero_maximum"),
			"Memory that contains 0s only.Included in shared amount. Through transparent page sharing, zero memory pages can be shared among virtual machines that run the same operating system. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemZeroMinimum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_zero_minimum"),
			"Memory that contains 0s only.Included in shared amount. Through transparent page sharing, zero memory pages can be shared among virtual machines that run the same operating system. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemZeroNone: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_zero_none"),
			"Memory that contains 0s only.Included in shared amount. Through transparent page sharing, zero memory pages can be shared among virtual machines that run the same operating system. KB",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemZippedLatest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_zipped_latest"),
			"Memory (KB) zipped.",
			[]string{"vm_name", "host_name"},
			nil,
		),
		MemZipsavedLatest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "mem_zipsaved_latest"),
			"Memory (KB) saved due to memory zipping.",
			[]string{"vm_name", "host_name"},
			nil,
		),
		NetBroadcastrxSummation: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "net_broadcastrx_summation"),
			"Number of broadcast packets received during the sampling interval.",
			[]string{"vm_name", "host_name"},
			nil,
		),
		NetBroadcasttxSummation: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "net_broadcasttx_summation"),
			"Number of broadcast packets transmitted during the sampling interval.",
			[]string{"vm_name", "host_name"},
			nil,
		),
		NetBytesrxAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "net_bytesrx_average"),
			"Average amount of data received per second. kiloBytesPerSecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
		NetBytestxAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "net_bytestx_average"),
			"Average amount of data transmitted per second. kiloBytesPerSecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
		NetDroppedrxSummation: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "net_droppedrx_summation"),
			"Number of received packets dropped during the collection interval.",
			[]string{"vm_name", "host_name"},
			nil,
		),
		NetDroppedtxSummation: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "net_droppedtx_summation"),
			"Number of transmitted packets dropped during the collection interval.",
			[]string{"vm_name", "host_name"},
			nil,
		),
		NetMulticastrxSummation: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "net_multicastrx_summation"),
			"Number of multicast packets received during the sampling interval.",
			[]string{"vm_name", "host_name"},
			nil,
		),
		NetMulticasttxSummation: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "net_multicasttx_summation"),
			"Number of multicast packets transmitted during the sampling interval.",
			[]string{"vm_name", "host_name"},
			nil,
		),
		NetPacketsrxSummation: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "net_packetsrx_summation"),
			"Number of packets received during the interval.",
			[]string{"vm_name", "host_name"},
			nil,
		),
		NetPacketstxSummation: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "net_packetstx_summation"),
			"Number of packets transmitted during the interval.",
			[]string{"vm_name", "host_name"},
			nil,
		),
		NetPnicbytesrxAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "net_pnicbytesrx_average"),
			"Average pNic I/O",
			[]string{"vm_name", "host_name"},
			nil,
		),
		NetPnicbytestxAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "net_pnicbytestx_average"),
			"Average pNic I/O",
			[]string{"vm_name", "host_name"},
			nil,
		),
		NetReceivedAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "net_received_average"),
			"Average rate at which data was received during the interval. This represents the bandwidth of the network. kiloBytesPerSecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
		NetTransmittedAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "net_transmitted_average"),
			"Average rate at which data was transmitted during the interval. This represents the bandwidth of the network. kiloBytesPerSecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
		NetUsageAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "net_usage_average"),
			"Network utilization (combined transmit- and receive-rates) during the interval. kiloBytesPerSecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
		NetUsageMaximum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "net_usage_maximum"),
			"Network utilization (combined transmit- and receive-rates) during the interval. kiloBytesPerSecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
		NetUsageMinimum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "net_usage_minimum"),
			"Network utilization (combined transmit- and receive-rates) during the interval. kiloBytesPerSecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
		NetUsageNone: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "net_usage_none"),
			"Network utilization (combined transmit- and receive-rates) during the interval. kiloBytesPerSecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
		PowerEnergySummation: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "power_energy_summation"),
			"power_energy_summation. joule",
			[]string{"vm_name", "host_name"},
			nil,
		),
		PowerPowerAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "power_power_average"),
			"Current power usage. watt",
			[]string{"vm_name", "host_name"},
			nil,
		),
		RescpuActav15Latest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "rescpu_actav15_latest"),
			"CPU active average over 15 minutes. percent",
			[]string{"vm_name", "host_name"},
			nil,
		),
		RescpuActav1Latest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "rescpu_actav1_latest"),
			"CPU active average over 1 minutes. percent",
			[]string{"vm_name", "host_name"},
			nil,
		),
		RescpuActav5Latest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "rescpu_actav5_latest"),
			"CPU active average over 5 minutes. percent",
			[]string{"vm_name", "host_name"},
			nil,
		),
		RescpuActpk15Latest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "rescpu_actpk15_latest"),
			"CPU active peak over 15 minutes. percent",
			[]string{"vm_name", "host_name"},
			nil,
		),
		RescpuActpk1Latest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "rescpu_actpk1_latest"),
			"CPU active peak over 1 minutes. percent",
			[]string{"vm_name", "host_name"},
			nil,
		),
		RescpuActpk5Latest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "rescpu_actpk5_latest"),
			"CPU active peak over 1 minutes. percent",
			[]string{"vm_name", "host_name"},
			nil,
		),
		RescpuMaxlimited15Latest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "rescpu_maxlimited15_latest"),
			"Amount of CPU resources over the limit that were refused, average over 15 minutes. percent",
			[]string{"vm_name", "host_name"},
			nil,
		),
		RescpuMaxlimited1Latest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "rescpu_maxlimited1_latest"),
			"Amount of CPU resources over the limit that were refused, average over 1 minutes. percent",
			[]string{"vm_name", "host_name"},
			nil,
		),
		RescpuMaxlimited5Latest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "rescpu_maxlimited5_latest"),
			"Amount of CPU resources over the limit that were refused, average over 5 minutes. percent",
			[]string{"vm_name", "host_name"},
			nil,
		),
		RescpuRunav15Latest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "rescpu_runav15_latest"),
			"CPU running average over 15 minutes. percent",
			[]string{"vm_name", "host_name"},
			nil,
		),
		RescpuRunav1Latest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "rescpu_runav1_latest"),
			"CPU running average over 1 minutes. percent",
			[]string{"vm_name", "host_name"},
			nil,
		),
		RescpuRunav5Latest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "rescpu_runav5_latest"),
			"CPU running average over 5 minutes. percent",
			[]string{"vm_name", "host_name"},
			nil,
		),
		RescpuRunpk15Latest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "rescpu_runpk15_latest"),
			"CPU running peak over 15 minutes. percent",
			[]string{"vm_name", "host_name"},
			nil,
		),
		RescpuRunpk1Latest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "rescpu_runpk1_latest"),
			"CPU running peak over 1 minutes. percent",
			[]string{"vm_name", "host_name"},
			nil,
		),
		RescpuRunpk5Latest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "rescpu_runpk5_latest"),
			"CPU running peak over 5 minutes. percent",
			[]string{"vm_name", "host_name"},
			nil,
		),
		RescpuSamplecountLatest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "rescpu_samplecount_latest"),
			"Group CPU sample count. number",
			[]string{"vm_name", "host_name"},
			nil,
		),
		RescpuSampleperiodLatest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "rescpu_sampleperiod_latest"),
			"Group CPU sample period. millisecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
		SysHeartbeatLatest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "sys_heartbeat_latest"),
			"Number of heartbeats issued per virtual machine during the interval. number",
			[]string{"vm_name", "host_name"},
			nil,
		),
		SysOsuptimeLatest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "sys_osuptime_latest"),
			"Total time elapsed, in seconds, since last operating system boot-up. second",
			[]string{"vm_name", "host_name"},
			nil,
		),
		SysUptimeLatest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "sys_uptime_latest"),
			"Total time elapsed, in seconds, since last system startup. second",
			[]string{"vm_name", "host_name"},
			nil,
		),
		VirtualdiskReadAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "virtualdisk_read_average"),
			"Rate of reading data from the virtual disk. kiloBytesPerSecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
		VirtualdiskWriteAverage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "vm", "virtualdisk_write_average"),
			"Rate of writing data to the virtual disk. kiloBytesPerSecond",
			[]string{"vm_name", "host_name"},
			nil,
		),
	}
}

func (c Collector) Describe(descs chan<- *prometheus.Desc) {
	ds := []*prometheus.Desc{
		c.HostPowerState,
		c.HostMaintenanceMode,
		c.HostBoot,
		c.TotalCpu,
		c.UsageCpu,
		c.TotalMem,
		c.UsageMem,
		c.DiskOk,
		c.NetworkPNICSpeed,
		c.HostHardwareInfo,
		c.NumCpuPkgs,
		c.CpuMhz,
		c.NumCpuCores,
		c.NumCpuThreads,
		c.NumNics,
		c.NumHBAs,
		c.TotalDs,
		c.UsageDs,
		c.VmGuestInfo,
		c.VmGuestPowerState,
		c.VmGuestToolsStatus,
		c.VmGuestToolsVersion,
		c.VmGuestStorageCommitted,
		c.VmGuestStorageUnCommitted,
		c.VmBoot,
		c.VmCpuAval,
		c.VmCpuUsage,
		c.VmNumCpu,
		c.VmMemAval,
		c.VmMemUsage,
		//VM Performance Metrics
		c.CpuCostopSummation,
		c.CpuDemandAverage,
		c.CpuDemandentitlementratioLatest,
		c.CpuEntitlementLatest,
		c.CpuIdleSummation,
		c.CpuLatencyAverage,
		c.CpuMaxlimitedSummation,
		c.CpuOverlapSummation,
		c.CpuReadinessAverage,
		c.CpuReadySummation,
		c.CpuRunSummation,
		c.CpuSwapwaitSummation,
		c.CpuSystemSummation,
		c.CpuUsageAverage,
		c.CpuUsageMaximum,
		c.CpuUsageMinimum,
		c.CpuUsageNone,
		c.CpuUsagemhzAverage,
		c.CpuUsagemhzMaximum,
		c.CpuUsagemhzMinimum,
		c.CpuUsagemhzNone,
		c.CpuUsedSummation,
		c.CpuWaitSummation,
		c.DatastoreMaxtotallatencyLatest,
		c.DiskMaxtotallatencyLatest,
		c.DiskReadAverage,
		c.DiskUsageAverage,
		c.DiskUsageMaximum,
		c.DiskUsageMinimum,
		c.DiskUsageNone,
		c.DiskWriteAverage,
		c.MemActiveAverage,
		c.MemActiveMaximum,
		c.MemActiveMinimum,
		c.MemActiveNone,
		c.MemActivewriteAverage,
		c.MemCompressedAverage,
		c.MemCompressionrateAverage,
		c.MemConsumedAverage,
		c.MemConsumedMaximum,
		c.MemConsumedMinimum,
		c.MemConsumedNone,
		c.MemDecompressionrateAverage,
		c.MemEntitlementAverage,
		c.MemGrantedAverage,
		c.MemGrantedMaximum,
		c.MemGrantedMinimum,
		c.MemGrantedNone,
		c.MemLatencyAverage,
		c.MemLlswapinrateAverage,
		c.MemLlswapoutrateAverage,
		c.MemLlswapusedAverage,
		c.MemLlswapusedMaximum,
		c.MemLlswapusedMinimum,
		c.MemLlswapusedNone,
		c.MemOverheadAverage,
		c.MemOverheadMaximum,
		c.MemOverheadMinimum,
		c.MemOverheadNone,
		c.MemOverheadmaxAverage,
		c.MemOverheadtouchedAverage,
		c.MemSharedAverage,
		c.MemSharedMaximum,
		c.MemSharedMinimum,
		c.MemSharedNone,
		c.MemSwapinAverage,
		c.MemSwapinMaximum,
		c.MemSwapinMinimum,
		c.MemSwapinNone,
		c.MemSwapinrateAverage,
		c.MemSwapoutAverage,
		c.MemSwapoutMaximum,
		c.MemSwapoutMinimum,
		c.MemSwapoutNone,
		c.MemSwapoutrateAverage,
		c.MemSwappedAverage,
		c.MemSwappedMaximum,
		c.MemSwappedMinimum,
		c.MemSwappedNone,
		c.MemSwaptargetAverage,
		c.MemSwaptargetMaximum,
		c.MemSwaptargetMinimum,
		c.MemSwaptargetNone,
		c.MemUsageAverage,
		c.MemUsageMaximum,
		c.MemUsageMinimum,
		c.MemUsageNone,
		c.MemVmmemctlAverage,
		c.MemVmmemctlMaximum,
		c.MemVmmemctlMinimum,
		c.MemVmmemctlNone,
		c.MemVmmemctltargetAverage,
		c.MemVmmemctltargetMaximum,
		c.MemVmmemctltargetMinimum,
		c.MemVmmemctltargetNone,
		c.MemZeroAverage,
		c.MemZeroMaximum,
		c.MemZeroMinimum,
		c.MemZeroNone,
		c.MemZippedLatest,
		c.MemZipsavedLatest,
		c.NetBroadcastrxSummation,
		c.NetBroadcasttxSummation,
		c.NetBytesrxAverage,
		c.NetBytestxAverage,
		c.NetDroppedrxSummation,
		c.NetDroppedtxSummation,
		c.NetMulticastrxSummation,
		c.NetMulticasttxSummation,
		c.NetPacketsrxSummation,
		c.NetPacketstxSummation,
		c.NetPnicbytesrxAverage,
		c.NetPnicbytestxAverage,
		c.NetReceivedAverage,
		c.NetTransmittedAverage,
		c.NetUsageAverage,
		c.NetUsageMaximum,
		c.NetUsageMinimum,
		c.NetUsageNone,
		c.PowerEnergySummation,
		c.PowerPowerAverage,
		c.RescpuActav15Latest,
		c.RescpuActav1Latest,
		c.RescpuActav5Latest,
		c.RescpuActpk15Latest,
		c.RescpuActpk1Latest,
		c.RescpuActpk5Latest,
		c.RescpuMaxlimited15Latest,
		c.RescpuMaxlimited1Latest,
		c.RescpuMaxlimited5Latest,
		c.RescpuRunav15Latest,
		c.RescpuRunav1Latest,
		c.RescpuRunav5Latest,
		c.RescpuRunpk15Latest,
		c.RescpuRunpk1Latest,
		c.RescpuRunpk5Latest,
		c.RescpuSamplecountLatest,
		c.RescpuSampleperiodLatest,
		c.SysHeartbeatLatest,
		c.SysOsuptimeLatest,
		c.SysUptimeLatest,
		c.VirtualdiskReadAverage,
		c.VirtualdiskWriteAverage,
	}

	for _, d := range ds {
		descs <- d
	}
}

func (c Collector) Collect(ch chan<- prometheus.Metric) {
	//Call Api and return metrics

	s, err := c.ss.status()

	if err != nil {
		c.ss.error(fmt.Errorf("failed collecting VMWare metrics: %v", err))
		ch <- prometheus.NewInvalidMetric(c.HostHardwareInfo, err)
		return
	}

	ch <- prometheus.MustNewConstMetric(
		c.HostPowerState,
		prometheus.GaugeValue,
		s.HostPowerState,
		s.HostName,
	)
	ch <- prometheus.MustNewConstMetric(
		c.HostMaintenanceMode,
		prometheus.GaugeValue,
		s.HostMaintenanceMode,
		s.HostName,
	)
	ch <- prometheus.MustNewConstMetric(
		c.HostBoot,
		prometheus.GaugeValue,
		s.HostBoot,
		s.HostName,
	)
	ch <- prometheus.MustNewConstMetric(
		c.TotalCpu,
		prometheus.GaugeValue,
		s.TotalCpu,
		s.HostName,
	)
	ch <- prometheus.MustNewConstMetric(
		c.UsageCpu,
		prometheus.GaugeValue,
		s.UsageCpu,
		s.HostName,
	)
	ch <- prometheus.MustNewConstMetric(
		c.TotalMem,
		prometheus.GaugeValue,
		s.TotalMem,
		s.HostName,
	)
	ch <- prometheus.MustNewConstMetric(
		c.UsageMem,
		prometheus.GaugeValue,
		s.UsageMem,
		s.HostName,
	)
	for _, d := range s.DiskOk {
		ch <- prometheus.MustNewConstMetric(
			c.DiskOk,
			prometheus.GaugeValue,
			d.ok,
			s.HostName, d.device,
		)
	}

	for _, n := range s.NetworkPNICSpeed {
		ch <- prometheus.MustNewConstMetric(
			c.NetworkPNICSpeed,
			prometheus.GaugeValue,
			n.speed,
			s.HostName, n.device, n.mac,
		)
	}

	ch <- prometheus.MustNewConstMetric(
		c.HostHardwareInfo,
		prometheus.GaugeValue,
		1,
		s.HostName, s.HW.Vendor, s.HW.Model, s.HW.Uuid, s.HW.CpuModel,
	)
	ch <- prometheus.MustNewConstMetric(
		c.NumCpuPkgs,
		prometheus.GaugeValue,
		s.HW.NumCpuPkgs,
		s.HostName,
	)
	ch <- prometheus.MustNewConstMetric(
		c.CpuMhz,
		prometheus.GaugeValue,
		s.HW.CpuMhz,
		s.HostName,
	)
	ch <- prometheus.MustNewConstMetric(
		c.NumCpuCores,
		prometheus.GaugeValue,
		s.HW.NumCpuCores,
		s.HostName,
	)
	ch <- prometheus.MustNewConstMetric(
		c.NumCpuThreads,
		prometheus.GaugeValue,
		s.HW.NumCpuThreads,
		s.HostName,
	)
	ch <- prometheus.MustNewConstMetric(
		c.NumNics,
		prometheus.GaugeValue,
		s.HW.NumNics,
		s.HostName,
	)
	ch <- prometheus.MustNewConstMetric(
		c.NumHBAs,
		prometheus.GaugeValue,
		s.HW.NumHBAs,
		s.HostName,
	)

	for _, ds := range s.DS {
		ch <- prometheus.MustNewConstMetric(
			c.TotalDs,
			prometheus.GaugeValue,
			ds.capacity,
			ds.dsname, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.UsageDs,
			prometheus.GaugeValue,
			ds.freespace,
			ds.dsname, s.HostName,
		)
	}
	for _, vm := range s.VMS {
		ch <- prometheus.MustNewConstMetric(
			c.VmGuestInfo,
			prometheus.GaugeValue,
			1,
			vm.VmName, s.HostName, vm.VmGuestId, vm.VmGuestFullName, vm.VmGuestIpAddr,
		)
		ch <- prometheus.MustNewConstMetric(
			c.VmGuestPowerState,
			prometheus.GaugeValue,
			vm.VmPowerState,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.VmGuestToolsStatus,
			prometheus.GaugeValue,
			vm.VmGuestToolsStatus,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.VmGuestToolsVersion,
			prometheus.GaugeValue,
			vm.VmGuestToolsVersion,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.VmGuestStorageCommitted,
			prometheus.GaugeValue,
			vm.VmGuestStorageCommitted,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.VmGuestStorageUnCommitted,
			prometheus.GaugeValue,
			vm.VmGuestStorageUnCommitted,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.VmBoot,
			prometheus.GaugeValue,
			vm.VmBoot,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.VmCpuAval,
			prometheus.GaugeValue,
			vm.VmCpuAval,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.VmCpuUsage,
			prometheus.GaugeValue,
			vm.VmCpuUsage,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.VmNumCpu,
			prometheus.GaugeValue,
			vm.VmNumCpu,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.VmMemAval,
			prometheus.GaugeValue,
			vm.VmMemAval,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.VmMemUsage,
			prometheus.GaugeValue,
			vm.VmMemUsage,
			vm.VmName, s.HostName,
		)
		//Performance Metrics Start This
		ch <- prometheus.MustNewConstMetric(
			c.CpuCostopSummation,
			prometheus.GaugeValue,
			vm.Perf.CPU_COSTOP_SUMMATION,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.CpuDemandAverage,
			prometheus.GaugeValue,
			vm.Perf.CPU_DEMANDENTITLEMENTRATIO_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.CpuDemandentitlementratioLatest,
			prometheus.GaugeValue,
			vm.Perf.CPU_DEMAND_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.CpuEntitlementLatest,
			prometheus.GaugeValue,
			vm.Perf.CPU_ENTITLEMENT_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.CpuIdleSummation,
			prometheus.GaugeValue,
			vm.Perf.CPU_IDLE_SUMMATION,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.CpuLatencyAverage,
			prometheus.GaugeValue,
			vm.Perf.CPU_LATENCY_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.CpuMaxlimitedSummation,
			prometheus.GaugeValue,
			vm.Perf.CPU_MAXLIMITED_SUMMATION,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.CpuOverlapSummation,
			prometheus.GaugeValue,
			vm.Perf.CPU_OVERLAP_SUMMATION,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.CpuReadinessAverage,
			prometheus.GaugeValue,
			vm.Perf.CPU_READINESS_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.CpuReadySummation,
			prometheus.GaugeValue,
			vm.Perf.CPU_READY_SUMMATION,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.CpuRunSummation,
			prometheus.GaugeValue,
			vm.Perf.CPU_RUN_SUMMATION,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.CpuSwapwaitSummation,
			prometheus.GaugeValue,
			vm.Perf.CPU_SWAPWAIT_SUMMATION,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.CpuSystemSummation,
			prometheus.GaugeValue,
			vm.Perf.CPU_SYSTEM_SUMMATION,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.CpuUsageAverage,
			prometheus.GaugeValue,
			vm.Perf.CPU_USAGEMHZ_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.CpuUsageMaximum,
			prometheus.GaugeValue,
			vm.Perf.CPU_USAGEMHZ_MAXIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.CpuUsageMinimum,
			prometheus.GaugeValue,
			vm.Perf.CPU_USAGEMHZ_MINIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.CpuUsageNone,
			prometheus.GaugeValue,
			vm.Perf.CPU_USAGEMHZ_NONE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.CpuUsagemhzAverage,
			prometheus.GaugeValue,
			vm.Perf.CPU_USAGE_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.CpuUsagemhzMaximum,
			prometheus.GaugeValue,
			vm.Perf.CPU_USAGE_MAXIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.CpuUsagemhzMinimum,
			prometheus.GaugeValue,
			vm.Perf.CPU_USAGE_MINIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.CpuUsagemhzNone,
			prometheus.GaugeValue,
			vm.Perf.CPU_USAGE_NONE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.CpuUsedSummation,
			prometheus.GaugeValue,
			vm.Perf.CPU_USED_SUMMATION,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.CpuWaitSummation,
			prometheus.GaugeValue,
			vm.Perf.CPU_WAIT_SUMMATION,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.DatastoreMaxtotallatencyLatest,
			prometheus.GaugeValue,
			vm.Perf.DATASTORE_MAXTOTALLATENCY_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.DiskMaxtotallatencyLatest,
			prometheus.GaugeValue,
			vm.Perf.DISK_MAXTOTALLATENCY_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.DiskReadAverage,
			prometheus.GaugeValue,
			vm.Perf.DISK_READ_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.DiskUsageAverage,
			prometheus.GaugeValue,
			vm.Perf.DISK_USAGE_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.DiskUsageMaximum,
			prometheus.GaugeValue,
			vm.Perf.DISK_USAGE_MAXIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.DiskUsageMinimum,
			prometheus.GaugeValue,
			vm.Perf.DISK_USAGE_MINIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.DiskUsageNone,
			prometheus.GaugeValue,
			vm.Perf.DISK_USAGE_NONE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.DiskWriteAverage,
			prometheus.GaugeValue,
			vm.Perf.DISK_WRITE_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemActiveAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_ACTIVEWRITE_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemActiveMaximum,
			prometheus.GaugeValue,
			vm.Perf.MEM_ACTIVE_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemActiveMinimum,
			prometheus.GaugeValue,
			vm.Perf.MEM_ACTIVE_MAXIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemActiveNone,
			prometheus.GaugeValue,
			vm.Perf.MEM_ACTIVE_MINIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemActivewriteAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_ACTIVE_NONE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemCompressedAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_COMPRESSED_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemCompressionrateAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_COMPRESSIONRATE_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemConsumedAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_CONSUMED_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemConsumedMaximum,
			prometheus.GaugeValue,
			vm.Perf.MEM_CONSUMED_MAXIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemConsumedMinimum,
			prometheus.GaugeValue,
			vm.Perf.MEM_CONSUMED_MINIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemConsumedNone,
			prometheus.GaugeValue,
			vm.Perf.MEM_CONSUMED_NONE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemDecompressionrateAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_DECOMPRESSIONRATE_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemEntitlementAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_ENTITLEMENT_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemGrantedAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_GRANTED_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemGrantedMaximum,
			prometheus.GaugeValue,
			vm.Perf.MEM_GRANTED_MAXIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemGrantedMinimum,
			prometheus.GaugeValue,
			vm.Perf.MEM_GRANTED_MINIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemGrantedNone,
			prometheus.GaugeValue,
			vm.Perf.MEM_GRANTED_NONE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemLatencyAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_LATENCY_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemLlswapinrateAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_LLSWAPINRATE_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemLlswapoutrateAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_LLSWAPOUTRATE_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemLlswapusedAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_LLSWAPUSED_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemLlswapusedMaximum,
			prometheus.GaugeValue,
			vm.Perf.MEM_LLSWAPUSED_MAXIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemLlswapusedMinimum,
			prometheus.GaugeValue,
			vm.Perf.MEM_LLSWAPUSED_MINIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemLlswapusedNone,
			prometheus.GaugeValue,
			vm.Perf.MEM_LLSWAPUSED_NONE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemOverheadAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_OVERHEADMAX_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemOverheadMaximum,
			prometheus.GaugeValue,
			vm.Perf.MEM_OVERHEADTOUCHED_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemOverheadMinimum,
			prometheus.GaugeValue,
			vm.Perf.MEM_OVERHEAD_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemOverheadNone,
			prometheus.GaugeValue,
			vm.Perf.MEM_OVERHEAD_MAXIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemOverheadmaxAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_OVERHEAD_MINIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemOverheadtouchedAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_OVERHEAD_NONE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemSharedAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_SHARED_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemSharedMaximum,
			prometheus.GaugeValue,
			vm.Perf.MEM_SHARED_MAXIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemSharedMinimum,
			prometheus.GaugeValue,
			vm.Perf.MEM_SHARED_MINIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemSharedNone,
			prometheus.GaugeValue,
			vm.Perf.MEM_SHARED_NONE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemSwapinAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_SWAPINRATE_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemSwapinMaximum,
			prometheus.GaugeValue,
			vm.Perf.MEM_SWAPIN_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemSwapinMinimum,
			prometheus.GaugeValue,
			vm.Perf.MEM_SWAPIN_MAXIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemSwapinNone,
			prometheus.GaugeValue,
			vm.Perf.MEM_SWAPIN_MINIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemSwapinrateAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_SWAPIN_NONE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemSwapoutAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_SWAPOUTRATE_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemSwapoutMaximum,
			prometheus.GaugeValue,
			vm.Perf.MEM_SWAPOUT_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemSwapoutMinimum,
			prometheus.GaugeValue,
			vm.Perf.MEM_SWAPOUT_MAXIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemSwapoutNone,
			prometheus.GaugeValue,
			vm.Perf.MEM_SWAPOUT_MINIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemSwapoutrateAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_SWAPOUT_NONE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemSwappedAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_SWAPPED_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemSwappedMaximum,
			prometheus.GaugeValue,
			vm.Perf.MEM_SWAPPED_MAXIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemSwappedMinimum,
			prometheus.GaugeValue,
			vm.Perf.MEM_SWAPPED_MINIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemSwappedNone,
			prometheus.GaugeValue,
			vm.Perf.MEM_SWAPPED_NONE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemSwaptargetAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_SWAPTARGET_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemSwaptargetMaximum,
			prometheus.GaugeValue,
			vm.Perf.MEM_SWAPTARGET_MAXIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemSwaptargetMinimum,
			prometheus.GaugeValue,
			vm.Perf.MEM_SWAPTARGET_MINIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemSwaptargetNone,
			prometheus.GaugeValue,
			vm.Perf.MEM_SWAPTARGET_NONE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemUsageAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_USAGE_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemUsageMaximum,
			prometheus.GaugeValue,
			vm.Perf.MEM_USAGE_MAXIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemUsageMinimum,
			prometheus.GaugeValue,
			vm.Perf.MEM_USAGE_MINIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemUsageNone,
			prometheus.GaugeValue,
			vm.Perf.MEM_USAGE_NONE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemVmmemctlAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_VMMEMCTLTARGET_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemVmmemctlMaximum,
			prometheus.GaugeValue,
			vm.Perf.MEM_VMMEMCTLTARGET_MAXIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemVmmemctlMinimum,
			prometheus.GaugeValue,
			vm.Perf.MEM_VMMEMCTLTARGET_MINIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemVmmemctlNone,
			prometheus.GaugeValue,
			vm.Perf.MEM_VMMEMCTLTARGET_NONE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemVmmemctltargetAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_VMMEMCTL_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemVmmemctltargetMaximum,
			prometheus.GaugeValue,
			vm.Perf.MEM_VMMEMCTL_MAXIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemVmmemctltargetMinimum,
			prometheus.GaugeValue,
			vm.Perf.MEM_VMMEMCTL_MINIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemVmmemctltargetNone,
			prometheus.GaugeValue,
			vm.Perf.MEM_VMMEMCTL_NONE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemZeroAverage,
			prometheus.GaugeValue,
			vm.Perf.MEM_ZERO_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemZeroMaximum,
			prometheus.GaugeValue,
			vm.Perf.MEM_ZERO_MAXIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemZeroMinimum,
			prometheus.GaugeValue,
			vm.Perf.MEM_ZERO_MINIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemZeroNone,
			prometheus.GaugeValue,
			vm.Perf.MEM_ZERO_NONE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemZippedLatest,
			prometheus.GaugeValue,
			vm.Perf.MEM_ZIPPED_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MemZipsavedLatest,
			prometheus.GaugeValue,
			vm.Perf.MEM_ZIPSAVED_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.NetBroadcastrxSummation,
			prometheus.GaugeValue,
			vm.Perf.NET_BROADCASTRX_SUMMATION,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.NetBroadcasttxSummation,
			prometheus.GaugeValue,
			vm.Perf.NET_BROADCASTTX_SUMMATION,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.NetBytesrxAverage,
			prometheus.GaugeValue,
			vm.Perf.NET_BYTESRX_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.NetBytestxAverage,
			prometheus.GaugeValue,
			vm.Perf.NET_BYTESTX_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.NetDroppedrxSummation,
			prometheus.GaugeValue,
			vm.Perf.NET_DROPPEDRX_SUMMATION,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.NetDroppedtxSummation,
			prometheus.GaugeValue,
			vm.Perf.NET_DROPPEDTX_SUMMATION,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.NetMulticastrxSummation,
			prometheus.GaugeValue,
			vm.Perf.NET_MULTICASTRX_SUMMATION,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.NetMulticasttxSummation,
			prometheus.GaugeValue,
			vm.Perf.NET_MULTICASTTX_SUMMATION,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.NetPacketsrxSummation,
			prometheus.GaugeValue,
			vm.Perf.NET_PACKETSRX_SUMMATION,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.NetPacketstxSummation,
			prometheus.GaugeValue,
			vm.Perf.NET_PACKETSTX_SUMMATION,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.NetPnicbytesrxAverage,
			prometheus.GaugeValue,
			vm.Perf.NET_PNICBYTESRX_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.NetPnicbytestxAverage,
			prometheus.GaugeValue,
			vm.Perf.NET_PNICBYTESTX_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.NetReceivedAverage,
			prometheus.GaugeValue,
			vm.Perf.NET_RECEIVED_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.NetTransmittedAverage,
			prometheus.GaugeValue,
			vm.Perf.NET_TRANSMITTED_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.NetUsageAverage,
			prometheus.GaugeValue,
			vm.Perf.NET_USAGE_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.NetUsageMaximum,
			prometheus.GaugeValue,
			vm.Perf.NET_USAGE_MAXIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.NetUsageMinimum,
			prometheus.GaugeValue,
			vm.Perf.NET_USAGE_MINIMUM,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.NetUsageNone,
			prometheus.GaugeValue,
			vm.Perf.NET_USAGE_NONE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.PowerEnergySummation,
			prometheus.GaugeValue,
			vm.Perf.POWER_ENERGY_SUMMATION,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.PowerPowerAverage,
			prometheus.GaugeValue,
			vm.Perf.POWER_POWER_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.RescpuActav15Latest,
			prometheus.GaugeValue,
			vm.Perf.RESCPU_ACTAV15_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.RescpuActav1Latest,
			prometheus.GaugeValue,
			vm.Perf.RESCPU_ACTAV1_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.RescpuActav5Latest,
			prometheus.GaugeValue,
			vm.Perf.RESCPU_ACTAV5_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.RescpuActpk15Latest,
			prometheus.GaugeValue,
			vm.Perf.RESCPU_ACTPK15_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.RescpuActpk1Latest,
			prometheus.GaugeValue,
			vm.Perf.RESCPU_ACTPK1_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.RescpuActpk5Latest,
			prometheus.GaugeValue,
			vm.Perf.RESCPU_ACTPK5_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.RescpuMaxlimited15Latest,
			prometheus.GaugeValue,
			vm.Perf.RESCPU_MAXLIMITED15_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.RescpuMaxlimited1Latest,
			prometheus.GaugeValue,
			vm.Perf.RESCPU_MAXLIMITED1_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.RescpuMaxlimited5Latest,
			prometheus.GaugeValue,
			vm.Perf.RESCPU_MAXLIMITED5_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.RescpuRunav15Latest,
			prometheus.GaugeValue,
			vm.Perf.RESCPU_RUNAV15_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.RescpuRunav1Latest,
			prometheus.GaugeValue,
			vm.Perf.RESCPU_RUNAV1_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.RescpuRunav5Latest,
			prometheus.GaugeValue,
			vm.Perf.RESCPU_RUNAV5_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.RescpuRunpk15Latest,
			prometheus.GaugeValue,
			vm.Perf.RESCPU_RUNPK15_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.RescpuRunpk1Latest,
			prometheus.GaugeValue,
			vm.Perf.RESCPU_RUNPK1_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.RescpuRunpk5Latest,
			prometheus.GaugeValue,
			vm.Perf.RESCPU_RUNPK5_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.RescpuSamplecountLatest,
			prometheus.GaugeValue,
			vm.Perf.RESCPU_SAMPLECOUNT_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.RescpuSampleperiodLatest,
			prometheus.GaugeValue,
			vm.Perf.RESCPU_SAMPLEPERIOD_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.SysHeartbeatLatest,
			prometheus.GaugeValue,
			vm.Perf.SYS_HEARTBEAT_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.SysOsuptimeLatest,
			prometheus.GaugeValue,
			vm.Perf.SYS_OSUPTIME_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.SysUptimeLatest,
			prometheus.GaugeValue,
			vm.Perf.SYS_UPTIME_LATEST,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.VirtualdiskReadAverage,
			prometheus.GaugeValue,
			vm.Perf.VIRTUALDISK_READ_AVERAGE,
			vm.VmName, s.HostName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.VirtualdiskWriteAverage,
			prometheus.GaugeValue,
			vm.Perf.VIRTUALDISK_WRITE_AVERAGE,
			vm.VmName, s.HostName,
		)

	}

}
