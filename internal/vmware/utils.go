package vmware

import (
	"context"
	"fmt"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/performance"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"strconv"
	"strings"
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

// sensorHealthOf is nil-safe: hosts may omit the health state of a sensor.
func sensorHealthOf(s types.BaseElementDescription) float64 {
	if s == nil {
		return 0
	}
	return sensorHealth(strings.ToLower(s.GetElementDescription().Key))
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

// perfBatchSize limits the number of entities in one QueryPerf request: a single request for every VM
// of a big vCenter may exceed the scrape timeout or the server limits.
const perfBatchSize = 100

type perfClient struct {
	manager  *performance.Manager
	counters map[string]*types.PerfCounterInfo
	logger   *logging.Logger
}

func newPerfClient(ctx context.Context, c *govmomi.Client, l *logging.Logger) (*perfClient, error) {
	manager := performance.NewManager(c.Client)
	counters, err := manager.CounterInfoByName(ctx)
	if err != nil {
		return nil, fmt.Errorf("counter info: %w", err)
	}
	return &perfClient{manager: manager, counters: counters, logger: l}, nil
}

// allCounters returns the names of every counter known to the server.
func (p *perfClient) allCounters() []string {
	names := make([]string, 0, len(p.counters))
	for name := range p.counters {
		names = append(names, name)
	}
	return names
}

// query returns the last real-time sample of the counters per entity (key is the managed object id).
// instance is "" for the aggregated value or "*" for every instance (NIC, datastore, ...).
// Counters unknown to the server are skipped; a failed batch is logged and the other batches still run.
func (p *perfClient) query(ctx context.Context, refs []types.ManagedObjectReference, names []string, instance string) (map[string][]vmMetric, error) {
	known := make([]string, 0, len(names))
	for _, n := range names {
		if _, ok := p.counters[n]; ok {
			known = append(known, n)
		}
	}
	if len(refs) == 0 || len(known) == 0 {
		return nil, nil
	}

	spec := types.PerfQuerySpec{
		MaxSample:  1,
		MetricId:   []types.PerfMetricId{{Instance: instance}},
		IntervalId: int32(interval),
	}

	metricsRes := make(map[string][]vmMetric)
	for start := 0; start < len(refs); start += perfBatchSize {
		end := start + perfBatchSize
		if end > len(refs) {
			end = len(refs)
		}
		if ctx.Err() != nil {
			return metricsRes, ctx.Err()
		}

		sample, err := p.manager.SampleByName(ctx, spec, known, refs[start:end])
		if err != nil {
			p.logger.Error(fmt.Sprintf("perf sample of entities %d-%d failed: %v", start, end, err))
			continue
		}
		result, err := p.manager.ToMetricSeries(ctx, sample)
		if err != nil {
			p.logger.Error(fmt.Sprintf("perf metric series of entities %d-%d failed: %v", start, end, err))
			continue
		}

		for _, metric := range result {
			for _, v := range metric.Value {
				if len(v.Value) == 0 {
					continue
				}
				instance := v.Instance
				if instance == "" {
					instance = "-"
				}
				var unit string
				if ci, ok := p.counters[v.Name]; ok && ci.UnitInfo != nil {
					unit = ci.UnitInfo.GetElementDescription().Label
				}
				metricsRes[metric.Entity.Value] = append(metricsRes[metric.Entity.Value], vmMetric{
					Instance:    instance,
					MetricName:  v.Name,
					MetricValue: v.ValueCSV(),
					MetricUnit:  unit,
				})
			}
		}
	}
	return metricsRes, nil
}

// instanceSamples converts instanced metrics to samples; the aggregated value ("-") is skipped
// because the aggregates are exported as separate metrics. resolve may rename an instance.
func instanceSamples(metrics []vmMetric, resolve map[string]string) []perfSample {
	var res []perfSample
	for _, m := range metrics {
		if m.Instance == "-" {
			continue
		}
		v, err := strconv.ParseFloat(m.MetricValue, 64)
		if err != nil {
			continue
		}
		inst := m.Instance
		if name, ok := resolve[inst]; ok {
			inst = name
		}
		res = append(res, perfSample{Counter: m.MetricName, Instance: inst, Value: v})
	}
	return res
}

// alarmCounts returns the number of red and yellow triggered alarms of an entity.
func alarmCounts(states []types.AlarmState) (red, yellow float64) {
	for _, st := range states {
		switch st.OverallStatus {
		case types.ManagedEntityStatusRed:
			red++
		case types.ManagedEntityStatusYellow:
			yellow++
		}
	}
	return red, yellow
}

func datastoreMaintenanceMode(s string) float64 {
	switch s {
	case "enteringMaintenance":
		return 1
	case "inMaintenance":
		return 2
	}
	return 0
}

func standbyMode(s string) float64 {
	switch s {
	case "entering":
		return 1
	case "in":
		return 2
	case "exiting":
		return 3
	}
	return 0
}

// dedupSensors drops sensors with an identical label set: a duplicated series fails the whole scrape.
func dedupSensors(in []NumericSensorInfo) []NumericSensorInfo {
	seen := make(map[[6]string]struct{}, len(in))
	out := make([]NumericSensorInfo, 0, len(in))
	for _, s := range in {
		key := [6]string{s.Name, s.CurrentReading, s.BaseUnits, s.SensorType, s.Id, s.SensorNumber}
		if _, dup := seen[key]; dup {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, s)
	}
	return out
}
