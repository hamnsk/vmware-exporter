package vmware

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/vmware/govmomi/vim25/types"
	"vmware-exporter/internal/config"
	"vmware-exporter/pkg/logging"
)

func TestParseTimeout(t *testing.T) {
	def := 60 * time.Second
	cases := map[string]time.Duration{
		"":       def,
		"30s":    30 * time.Second, // used to be multiplied by time.Second again
		"30":     30 * time.Second,
		"1m":     time.Minute,
		"junk":   def,
		"-5":     def,
		"0":      def,
		" 15 ":   15 * time.Second,
		"1500ms": 1500 * time.Millisecond,
	}
	for in, want := range cases {
		if got := parseTimeout(in, def); got != want {
			t.Errorf("parseTimeout(%q)=%v, want %v", in, got, want)
		}
	}
}

func TestSensorHealthOfNil(t *testing.T) {
	if got := sensorHealthOf(nil); got != 0 {
		t.Fatalf("nil health must map to 0, got %v", got)
	}
	if got := sensorHealthOf(&types.ElementDescription{Description: types.Description{}, Key: "Green"}); got != 1 {
		t.Fatalf("green must map to 1, got %v", got)
	}
}

// Probes of unreachable targets must fail without panics/exit, run concurrently and
// must not replace the global prometheus registry.
func TestProbeConcurrentAndGlobalRegistry(t *testing.T) {
	t.Setenv("VMWARE_EXPORTER_SCRAPE_TIMEOUT", "2s")
	logger := logging.GetLogger()
	h := GetHandler(&logger, config.GetConfig(), nil).(*exporterHandler)

	before := prometheus.DefaultGatherer
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/probe?target=127.0.0.1:1", nil)
			h.probe(rec, req)
			if rec.Code == 0 {
				t.Error("no response")
			}
		}()
	}
	wg.Add(1)
	go func() { // reload races with probes
		defer wg.Done()
		h.reload(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, "/-/reload", nil))
	}()
	wg.Wait()

	if prometheus.DefaultGatherer != before {
		t.Fatal("probe replaced prometheus.DefaultGatherer")
	}
}

type panicService struct{ msg string }

func (p panicService) statuses() ([]*Status, error) { panic(p.msg) }
func (p panicService) error(error)                  {}

// A panic inside Collect runs outside of net/http's recover and used to kill the process.
func TestCollectRecoversFromPanic(t *testing.T) {
	reg := prometheus.NewRegistry()
	reg.MustRegister(NewCollector(panicService{"boom"}))
	if _, err := reg.Gather(); err == nil {
		t.Fatal("expected gather error, got nil")
	}
}

type staticService struct{ st []*Status }

func (s staticService) statuses() ([]*Status, error) { return s.st, nil }
func (s staticService) error(error)                  {}

// Sensors and guest disks are absent on the simulator, so the output is checked on a crafted status.
// Duplicated sensors or partitions must not break the scrape (a duplicated label set fails Gather).
func TestExtraMetricsFromStatus(t *testing.T) {
	st := &Status{
		HostName: "esx1", HostConnected: 1, HostStandbyMode: 2, HostRedAlarms: 1, HostYellowAlarms: 2,
		SensorInfo: []NumericSensorInfo{
			{Name: "Fan 1", SensorType: "fan", BaseUnits: "RPM", Id: "1", Value: 5400},
			{Name: "Fan 1", SensorType: "fan", BaseUnits: "RPM", Id: "1", Value: 5400}, // duplicate
			{Name: "Temp", SensorType: "temperature", BaseUnits: "Degrees C", Id: "2", Value: 41.5},
		},
		DS: []totalds{{dsname: "ds1", capacity: 100, freespace: 40, uncommitted: 30, maintenance: 2, vms: 3, hosts: 2, redAlarms: 1}},
		VMS: []hvms{{
			VmName: "vm1", Template: 1, GuestToolsRunning: 1, RedAlarms: 1,
			GuestDisks: []guestDisk{{"/", 100, 10}, {"/var", 50, 5}},
		}},
	}
	reg := prometheus.NewRegistry()
	reg.MustRegister(NewCollector(staticService{[]*Status{st}}))
	mfs, err := reg.Gather()
	if err != nil {
		t.Fatalf("gather: %v", err)
	}
	got := map[string][]float64{}
	for _, mf := range mfs {
		for _, m := range mf.GetMetric() {
			got[mf.GetName()] = append(got[mf.GetName()], m.GetGauge().GetValue())
		}
	}
	want := map[string][]float64{
		"vmware_exporter_host_sensor_value":             {5400, 41.5},
		"vmware_exporter_host_standby_mode":             {2},
		"vmware_exporter_host_red_alarms":               {1},
		"vmware_exporter_host_yellow_alarms":            {2},
		"vmware_exporter_datastore_provisioned_size":    {90},
		"vmware_exporter_datastore_maintenance_mode":    {2},
		"vmware_exporter_datastore_vms":                 {3},
		"vmware_exporter_datastore_hosts":               {2},
		"vmware_exporter_datastore_red_alarms":          {1},
		"vmware_exporter_vm_template":                   {1},
		"vmware_exporter_vm_guest_tools_running_status": {1},
		"vmware_exporter_vm_guest_disk_capacity_size":   {100, 50},
		"vmware_exporter_vm_guest_disk_free_size":       {10, 5},
		"vmware_exporter_vm_red_alarms":                 {1},
	}
	for name, vals := range want {
		g := got[name]
		if len(g) != len(vals) {
			t.Errorf("%s: got %v, want %v", name, g, vals)
			continue
		}
		for _, v := range vals {
			found := false
			for _, x := range g {
				found = found || x == v
			}
			if !found {
				t.Errorf("%s: value %v missing in %v", name, v, g)
			}
		}
	}
}
