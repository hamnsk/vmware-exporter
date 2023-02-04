package vmware

import (
	"context"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/performance"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"vmware-exporter/pkg/logging"
)

type vmMetric struct {
	Instance    string
	MetricName  string
	MetricValue string
	MetricUnit  string
}

func totalCpu(hs mo.HostSystem) float64 {
	totalCPU := int64(hs.Summary.Hardware.CpuMhz) * int64(hs.Summary.Hardware.NumCpuCores)
	return float64(totalCPU)
}

func convertTime(vm mo.VirtualMachine) float64 {
	if vm.Summary.Runtime.BootTime == nil {
		return 0
	}
	return float64(vm.Summary.Runtime.BootTime.Unix())
}

func powerState(s types.HostSystemPowerState) float64 {
	if s == "poweredOn" {
		return 1
	}
	if s == "poweredOff" {
		return 2
	}
	if s == "standBy" {
		return 3
	}
	return 0
}

func sensorHealth(s string) float64 {
	if s == "green" {
		return 1
	}
	if s == "yellow" {
		return 2
	}
	if s == "red" {
		return 3
	}
	return 0
}

func powerStateVM(s types.VirtualMachinePowerState) float64 {
	if s == "poweredOn" {
		return 1
	}
	if s == "poweredOff" {
		return 2
	}
	if s == "standBy" {
		return 3
	}
	return 0
}

func guestToolsStatus(s string) float64 {
	if s == "toolsOk" {
		return 1
	}

	return 0
}

func guestToolsVersion(s string) float64 {
	if s == "guestToolsUnmanaged" {
		return 0
	}
	return 1
}

func maintenanceMode(s bool) float64 {
	if s {
		return 1
	}
	return 0
}

func perfMon(ctx context.Context, c *govmomi.Client, l *logging.Logger, vms []types.ManagedObjectReference) map[string][]vmMetric {

	var perfMetricsResult []vmMetric
	metricsRes := make(map[string][]vmMetric)

	// Create a PerfManager
	perfManager := performance.NewManager(c.Client)

	// Retrieve counters name list
	counters, err := perfManager.CounterInfoByName(ctx)
	if err != nil {
		l.Fatal(err.Error())
	}

	var names []string
	for name := range counters {
		names = append(names, name)
	}

	// Create PerfQuerySpec
	spec := types.PerfQuerySpec{
		MaxSample:  1,
		MetricId:   []types.PerfMetricId{{Instance: ""}},
		IntervalId: int32(interval),
	}

	// Query metrics
	sample, err := perfManager.SampleByName(ctx, spec, names, vms)
	if err != nil {
		l.Fatal(err.Error())
	}

	result, err := perfManager.ToMetricSeries(ctx, sample)
	if err != nil {
		l.Fatal(err.Error())
	}

	// Read result
	for _, metric := range result {
		vmNum := metric.Entity.Value
		for _, v := range metric.Value {
			counter := counters[v.Name]
			units := counter.UnitInfo.GetElementDescription().Label

			instance := v.Instance
			if instance == "" {
				instance = "-"
			}

			if len(v.Value) != 0 {
				metric := vmMetric{
					Instance:    instance,
					MetricName:  v.Name,
					MetricValue: v.ValueCSV(),
					MetricUnit:  units,
				}
				perfMetricsResult = append(perfMetricsResult, metric)
			}
		}

		metricsRes[vmNum] = perfMetricsResult

	}
	return metricsRes
}
