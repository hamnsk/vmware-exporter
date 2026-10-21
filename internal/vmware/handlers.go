package vmware

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"reflect"
	"sync"
	"vmware-exporter/internal/config"
	"vmware-exporter/internal/vault"
	"vmware-exporter/pkg/logging"
)

const (
	home    = "/"
	metrics = "/metrics"
	probe   = "/probe"
	reload  = "/-/reload"
)

const homeResponse = `<html>
			<head><title>Vmware Prometheus Exporter</title></head>
			<body>
			<h1>Vmware Prometheus Exporter</h1>
			<p><a href="` + "/metrics" + `">Show Metrics</a></p>
			<h2>More information:</h2>
			<p><a href="https://github.com/hamnsk/vmware-exporter">github.com/hamnsk/vmware-exporter</a></p>
			</body>
			</html>`

var _ Handler = &exporterHandler{}

type exporterHandler struct {
	logger      *logging.Logger
	mu          sync.RWMutex // protects cfg, which can be replaced by reload
	cfg         interface{}
	vaultClient vault.Client
}

type Handler interface {
	Register(router *mux.Router)
}

func GetHandler(logger *logging.Logger, cfg interface{}, vaultClient vault.Client) Handler {
	h := exporterHandler{
		logger:      logger,
		cfg:         cfg,
		vaultClient: vaultClient,
	}
	return &h
}

func (h *exporterHandler) Register(router *mux.Router) {
	router.HandleFunc(home, h.home).Methods(http.MethodGet)
	router.HandleFunc(probe, h.probe).Methods(http.MethodGet)
	router.HandleFunc(reload, h.reload).Methods(http.MethodPost)
	router.Handle(metrics, promhttp.Handler()).Methods(http.MethodGet)
}

func (h *exporterHandler) home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(homeResponse))
}

func (h *exporterHandler) probe(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	target := query.Get("target")
	if len(query["target"]) != 1 || target == "" {
		http.Error(w, "'target' parameter must be specified once", http.StatusBadRequest)
		return
	}

	// every probe works on its own copy of the config: credentials from Vault are per target
	// and must never leak into the shared config used by concurrent probes
	cfg := h.configCopy()
	appCfg := reflect.ValueOf(cfg).Elem()

	if appCfg.FieldByName("UseVault").Interface().(bool) {
		if h.vaultClient == nil {
			h.logger.Error("vault client is not initialized")
			http.Error(w, "vault client is not initialized", http.StatusInternalServerError)
			return
		}

		secretStoreName := appCfg.FieldByName("VaultSecretStoreName").Interface().(string)
		if len(secretStoreName) == 0 {
			h.logger.Error("name of kv2 secret store must be specified")
			http.Error(w, "name of kv2 secret store must be specified", http.StatusBadRequest)
			return
		}

		secretStorePath := appCfg.FieldByName("VaultSecretStorePath").Interface().(string)
		if len(secretStorePath) == 0 {
			h.logger.Error("path for secret must be specified")
			http.Error(w, "path for secret must be specified", http.StatusBadRequest)
			return
		}

		data, err := h.vaultClient.GetClient().KVv2(secretStoreName).Get(
			r.Context(),
			fmt.Sprintf("%s/%s", secretStorePath, target),
		)
		if err != nil {
			h.logger.Error(
				fmt.Sprintf("error occurred when get credentials from Vault for target: %s", target),
			)
			h.logger.Error(err.Error())
			http.Error(
				w,
				fmt.Sprintf("error occurred when get credentials from Vault for target: %s", target),
				http.StatusBadRequest,
			)
			return
		}

		username, uok := data.Data["username"].(string)
		password, pok := data.Data["password"].(string)
		if !uok || !pok {
			h.logger.Error(fmt.Sprintf("secret for target %s must contain string fields 'username' and 'password'", target))
			http.Error(w, "secret must contain 'username' and 'password'", http.StatusBadRequest)
			return
		}
		appCfg.FieldByName("VmwareUser").SetString(username)
		appCfg.FieldByName("VmwarePass").SetString(password)
	}

	// a separate registry per request: never touch the global prometheus.Default* variables,
	// otherwise concurrent probes mix up metrics and /metrics stops serving the default registry
	reg := prometheus.NewRegistry()
	reg.MustRegister(NewCollector(NewService(h.logger, target, cfg)))
	promhttp.HandlerFor(reg, promhttp.HandlerOpts{}).ServeHTTP(w, r)
}

// configCopy returns a shallow copy of the current config (it consists of plain values only).
func (h *exporterHandler) configCopy() interface{} {
	h.mu.RLock()
	defer h.mu.RUnlock()
	src := reflect.ValueOf(h.cfg).Elem()
	dst := reflect.New(src.Type())
	dst.Elem().Set(src)
	return dst.Interface()
}

func (h *exporterHandler) reload(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	h.cfg = config.GetConfig()
	h.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode("Reload config from env success")
	h.logger.Info("Reload config from env success")
}
