package vmware

import (
	"crypto/tls"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/vmware/govmomi/simulator"
	"vmware-exporter/pkg/logging"
)

// newSimService starts a vCenter simulator (several hosts in clusters, shared datastores, VMs)
// and returns a service pointed at it.
func newSimService(t *testing.T, model *simulator.Model) *service {
	t.Helper()
	if err := model.Create(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(model.Remove)
	model.Service.TLS = new(tls.Config) // the exporter always talks HTTPS
	srv := model.Service.NewServer()
	t.Cleanup(srv.Close)

	logger := logging.GetLogger()
	return &service{
		logger:         &logger,
		vmwareHost:     srv.URL.Host,
		vmwareUser:     "user",
		vmwarePassword: "pass",
		scrapeTimeout:  30 * time.Second,
	}
}

// vCenter has several hosts: every host must be reported with its own VMs and datastores.
func TestVCenterReportsAllHosts(t *testing.T) {
	model := simulator.VPX()
	model.Cluster = 2     // 2 clusters
	model.ClusterHost = 3 // 3 hosts each
	model.Host = 1        // + 1 standalone host
	model.Machine = 2     // VMs per host/resource pool
	svc := newSimService(t, model)

	statuses, err := svc.statuses()
	if err != nil {
		t.Fatalf("statuses: %v", err)
	}
	if len(statuses) < 7 {
		t.Fatalf("want >=7 hosts, got %d", len(statuses))
	}

	names := map[string]bool{}
	totalVMs := 0
	for _, st := range statuses {
		if st.HostName == "" || names[st.HostName] {
			t.Errorf("empty or duplicated host name %q", st.HostName)
		}
		names[st.HostName] = true
		totalVMs += len(st.VMS)
		if len(st.DS) == 0 {
			t.Errorf("host %s has no datastores", st.HostName)
		}
	}
	if totalVMs == 0 {
		t.Fatal("no VMs reported")
	}

	// the whole exposition must be valid: no duplicated series between hosts
	reg := prometheus.NewRegistry()
	reg.MustRegister(NewCollector(svc))
	mfs, err := reg.Gather()
	if err != nil {
		t.Fatalf("gather: %v", err)
	}
	hostSeries := 0
	for _, mf := range mfs {
		if mf.GetName() == "vmware_exporter_host_power_state" {
			hostSeries = len(mf.GetMetric())
		}
	}
	if hostSeries != len(statuses) {
		t.Fatalf("host_power_state series=%d, hosts=%d", hostSeries, len(statuses))
	}
	count := func(name string) int {
		for _, mf := range mfs {
			if mf.GetName() == name {
				return len(mf.GetMetric())
			}
		}
		return 0
	}
	if n := count("vmware_exporter_vcenter_info"); n != 1 {
		t.Errorf("vcenter_info must be exported once, got %d", n)
	}
	if n := count("vmware_exporter_host_cluster_info"); n != 6 { // 2 clusters x 3 hosts, the standalone host has none
		t.Errorf("host_cluster_info series=%d, want 6", n)
	}
	for _, name := range []string{"host_connected", "datastore_info", "datastore_accessible", "datastore_provisioned_size", "vm_snapshot_count"} {
		if count("vmware_exporter_"+name) == 0 {
			t.Errorf("metric %s is missing", name)
		}
	}
	if n, err := testutil.GatherAndCount(reg); err != nil || n == 0 {
		t.Fatalf("count=%d err=%v", n, err)
	}
}

// A standalone ESXi (single host) keeps working.
func TestSingleHostTarget(t *testing.T) {
	svc := newSimService(t, simulator.ESX())
	statuses, err := svc.statuses()
	if err != nil {
		t.Fatalf("statuses: %v", err)
	}
	if len(statuses) != 1 {
		t.Fatalf("want 1 host, got %d", len(statuses))
	}
	if statuses[0].HostName == "" || len(statuses[0].DS) == 0 {
		t.Fatalf("incomplete status: %+v", statuses[0])
	}
}
