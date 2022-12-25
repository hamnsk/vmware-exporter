package vmware

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/version"
)

const namespace = "vmware_exporter"

type diskOk struct {
	device string
	ok     float64
}

type pnic struct {
	device, mac string
	speed       float64
}

type totalds struct {
	dsname              string
	capacity, freespace float64
}

type hvms struct {
	VmName              string
	VmPowerState        float64
	VmGuestId           string
	VmGuestToolsStatus  float64
	VmGuestToolsVersion float64

	VmGuestFullName           string
	VmGuestIpAddr             string
	VmGuestStorageCommitted   float64
	VmGuestStorageUnCommitted float64
	VmBoot                    float64
	VmCpuAval                 float64
	VmCpuUsage                float64
	VmNumCpu                  float64
	VmMemAval                 float64
	VmMemUsage                float64
	Perf                      vmPerf
}
type vmPerf struct {
	CPU_COSTOP_SUMMATION              float64
	CPU_DEMANDENTITLEMENTRATIO_LATEST float64
	CPU_DEMAND_AVERAGE                float64
	CPU_ENTITLEMENT_LATEST            float64
	CPU_IDLE_SUMMATION                float64
	CPU_LATENCY_AVERAGE               float64
	CPU_MAXLIMITED_SUMMATION          float64
	CPU_OVERLAP_SUMMATION             float64
	CPU_READINESS_AVERAGE             float64
	CPU_READY_SUMMATION               float64
	CPU_RUN_SUMMATION                 float64
	CPU_SWAPWAIT_SUMMATION            float64
	CPU_SYSTEM_SUMMATION              float64
	CPU_USAGEMHZ_AVERAGE              float64
	CPU_USAGEMHZ_MAXIMUM              float64
	CPU_USAGEMHZ_MINIMUM              float64
	CPU_USAGEMHZ_NONE                 float64
	CPU_USAGE_AVERAGE                 float64
	CPU_USAGE_MAXIMUM                 float64
	CPU_USAGE_MINIMUM                 float64
	CPU_USAGE_NONE                    float64
	CPU_USED_SUMMATION                float64
	CPU_WAIT_SUMMATION                float64
	DATASTORE_MAXTOTALLATENCY_LATEST  float64
	DISK_MAXTOTALLATENCY_LATEST       float64
	DISK_READ_AVERAGE                 float64
	DISK_USAGE_AVERAGE                float64
	DISK_USAGE_MAXIMUM                float64
	DISK_USAGE_MINIMUM                float64
	DISK_USAGE_NONE                   float64
	DISK_WRITE_AVERAGE                float64
	MEM_ACTIVEWRITE_AVERAGE           float64
	MEM_ACTIVE_AVERAGE                float64
	MEM_ACTIVE_MAXIMUM                float64
	MEM_ACTIVE_MINIMUM                float64
	MEM_ACTIVE_NONE                   float64
	MEM_COMPRESSED_AVERAGE            float64
	MEM_COMPRESSIONRATE_AVERAGE       float64
	MEM_CONSUMED_AVERAGE              float64
	MEM_CONSUMED_MAXIMUM              float64
	MEM_CONSUMED_MINIMUM              float64
	MEM_CONSUMED_NONE                 float64
	MEM_DECOMPRESSIONRATE_AVERAGE     float64
	MEM_ENTITLEMENT_AVERAGE           float64
	MEM_GRANTED_AVERAGE               float64
	MEM_GRANTED_MAXIMUM               float64
	MEM_GRANTED_MINIMUM               float64
	MEM_GRANTED_NONE                  float64
	MEM_LATENCY_AVERAGE               float64
	MEM_LLSWAPINRATE_AVERAGE          float64
	MEM_LLSWAPOUTRATE_AVERAGE         float64
	MEM_LLSWAPUSED_AVERAGE            float64
	MEM_LLSWAPUSED_MAXIMUM            float64
	MEM_LLSWAPUSED_MINIMUM            float64
	MEM_LLSWAPUSED_NONE               float64
	MEM_OVERHEADMAX_AVERAGE           float64
	MEM_OVERHEADTOUCHED_AVERAGE       float64
	MEM_OVERHEAD_AVERAGE              float64
	MEM_OVERHEAD_MAXIMUM              float64
	MEM_OVERHEAD_MINIMUM              float64
	MEM_OVERHEAD_NONE                 float64
	MEM_SHARED_AVERAGE                float64
	MEM_SHARED_MAXIMUM                float64
	MEM_SHARED_MINIMUM                float64
	MEM_SHARED_NONE                   float64
	MEM_SWAPINRATE_AVERAGE            float64
	MEM_SWAPIN_AVERAGE                float64
	MEM_SWAPIN_MAXIMUM                float64
	MEM_SWAPIN_MINIMUM                float64
	MEM_SWAPIN_NONE                   float64
	MEM_SWAPOUTRATE_AVERAGE           float64
	MEM_SWAPOUT_AVERAGE               float64
	MEM_SWAPOUT_MAXIMUM               float64
	MEM_SWAPOUT_MINIMUM               float64
	MEM_SWAPOUT_NONE                  float64
	MEM_SWAPPED_AVERAGE               float64
	MEM_SWAPPED_MAXIMUM               float64
	MEM_SWAPPED_MINIMUM               float64
	MEM_SWAPPED_NONE                  float64
	MEM_SWAPTARGET_AVERAGE            float64
	MEM_SWAPTARGET_MAXIMUM            float64
	MEM_SWAPTARGET_MINIMUM            float64
	MEM_SWAPTARGET_NONE               float64
	MEM_USAGE_AVERAGE                 float64
	MEM_USAGE_MAXIMUM                 float64
	MEM_USAGE_MINIMUM                 float64
	MEM_USAGE_NONE                    float64
	MEM_VMMEMCTLTARGET_AVERAGE        float64
	MEM_VMMEMCTLTARGET_MAXIMUM        float64
	MEM_VMMEMCTLTARGET_MINIMUM        float64
	MEM_VMMEMCTLTARGET_NONE           float64
	MEM_VMMEMCTL_AVERAGE              float64
	MEM_VMMEMCTL_MAXIMUM              float64
	MEM_VMMEMCTL_MINIMUM              float64
	MEM_VMMEMCTL_NONE                 float64
	MEM_ZERO_AVERAGE                  float64
	MEM_ZERO_MAXIMUM                  float64
	MEM_ZERO_MINIMUM                  float64
	MEM_ZERO_NONE                     float64
	MEM_ZIPPED_LATEST                 float64
	MEM_ZIPSAVED_LATEST               float64
	NET_BROADCASTRX_SUMMATION         float64
	NET_BROADCASTTX_SUMMATION         float64
	NET_BYTESRX_AVERAGE               float64
	NET_BYTESTX_AVERAGE               float64
	NET_DROPPEDRX_SUMMATION           float64
	NET_DROPPEDTX_SUMMATION           float64
	NET_MULTICASTRX_SUMMATION         float64
	NET_MULTICASTTX_SUMMATION         float64
	NET_PACKETSRX_SUMMATION           float64
	NET_PACKETSTX_SUMMATION           float64
	NET_PNICBYTESRX_AVERAGE           float64
	NET_PNICBYTESTX_AVERAGE           float64
	NET_RECEIVED_AVERAGE              float64
	NET_TRANSMITTED_AVERAGE           float64
	NET_USAGE_AVERAGE                 float64
	NET_USAGE_MAXIMUM                 float64
	NET_USAGE_MINIMUM                 float64
	NET_USAGE_NONE                    float64
	POWER_ENERGY_SUMMATION            float64
	POWER_POWER_AVERAGE               float64
	RESCPU_ACTAV15_LATEST             float64
	RESCPU_ACTAV1_LATEST              float64
	RESCPU_ACTAV5_LATEST              float64
	RESCPU_ACTPK15_LATEST             float64
	RESCPU_ACTPK1_LATEST              float64
	RESCPU_ACTPK5_LATEST              float64
	RESCPU_MAXLIMITED15_LATEST        float64
	RESCPU_MAXLIMITED1_LATEST         float64
	RESCPU_MAXLIMITED5_LATEST         float64
	RESCPU_RUNAV15_LATEST             float64
	RESCPU_RUNAV1_LATEST              float64
	RESCPU_RUNAV5_LATEST              float64
	RESCPU_RUNPK15_LATEST             float64
	RESCPU_RUNPK1_LATEST              float64
	RESCPU_RUNPK5_LATEST              float64
	RESCPU_SAMPLECOUNT_LATEST         float64
	RESCPU_SAMPLEPERIOD_LATEST        float64
	SYS_HEARTBEAT_LATEST              float64
	SYS_OSUPTIME_LATEST               float64
	SYS_UPTIME_LATEST                 float64
	VIRTUALDISK_READ_AVERAGE          float64
	VIRTUALDISK_WRITE_AVERAGE         float64
}

type hwInfo struct {
	Vendor        string
	Model         string
	Uuid          string
	CpuModel      string
	NumCpuPkgs    float64
	CpuMhz        float64
	NumCpuCores   float64
	NumCpuThreads float64
	NumNics       float64
	NumHBAs       float64
}

type Status struct {
	HostName            string
	HostPowerState      float64
	HostMaintenanceMode float64
	HostBoot            float64
	TotalCpu            float64
	UsageCpu            float64
	TotalMem            float64
	UsageMem            float64
	DiskOk              []diskOk
	NetworkPNICSpeed    []pnic
	HW                  hwInfo
	DS                  []totalds
	VMS                 []hvms
}

func RegisterExporter() {
	prometheus.MustRegister(version.NewCollector("vmware_exporter"))
}
