package vmware

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"vmware-exporter/internal/config"
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
			</body>
			</html>`

var _ Handler = &exporterHandler{}

type exporterHandler struct {
	logger *logging.Logger
	cfg    interface{}
}

type Handler interface {
	Register(router *mux.Router)
}

func GetHandler(logger *logging.Logger, cfg interface{}) Handler {
	h := exporterHandler{
		logger: logger,
		cfg:    cfg,
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
	host := r.URL.Query().Get("target")
	if len(host) > 1 {
		service := NewService(h.logger, host, h.cfg)
		go func() {
			service.hostMetrics()
		}()
		go func() {
			service.dsMetrics()
		}()
		go func() {
			service.vmMetrics()
		}()
	}
	w.WriteHeader(http.StatusOK)
}

func (h *exporterHandler) reload(w http.ResponseWriter, r *http.Request) {
	h.cfg = config.GetConfig()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode("Reload config from env success")
	h.logger.Info("Reload config from env success")
}
