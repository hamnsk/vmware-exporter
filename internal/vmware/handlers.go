package vmware

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"reflect"
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

	if len(target) > 1 {
		appCfg := reflect.ValueOf(h.cfg).Elem()
		if appCfg.FieldByName("UseVault").Interface().(bool) {
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

			appCfg.FieldByName("VmwareUser").SetString(data.Data["username"].(string))
			appCfg.FieldByName("VmwarePass").SetString(data.Data["password"].(string))
		}

		svc := NewService(h.logger, target, h.cfg)
		reg := prometheus.NewRegistry()
		prometheus.DefaultRegisterer = reg
		prometheus.DefaultGatherer = reg
		reg.MustRegister(NewCollector(svc))
	}
	promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{}).ServeHTTP(w, r)
}

func (h *exporterHandler) reload(w http.ResponseWriter, r *http.Request) {
	h.cfg = config.GetConfig()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode("Reload config from env success")
	h.logger.Info("Reload config from env success")
}
