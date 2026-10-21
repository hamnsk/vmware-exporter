package app

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"strconv"
	"syscall"
	"time"
	"vmware-exporter/internal/config"
	"vmware-exporter/internal/vault"
	"vmware-exporter/internal/version"
	"vmware-exporter/internal/vmware"
	"vmware-exporter/pkg/logging"
)

type app struct {
	logger      logging.Logger
	config      interface{}
	appRouter   *mux.Router
	appSrv      *http.Server
	vaultClient vault.Client
}

const vaultKeepAlivePause = 5 * time.Second

// parseSeconds parses a Go duration ("30s") or a plain number of seconds ("30").
func parseSeconds(v string, def time.Duration) (time.Duration, error) {
	if v == "" {
		return def, nil
	}
	if d, err := time.ParseDuration(v); err == nil && d > 0 {
		return d, nil
	}
	if n, err := strconv.Atoi(v); err == nil && n > 0 {
		return time.Duration(n) * time.Second, nil
	}
	return def, fmt.Errorf("invalid duration %q", v)
}

func useVault(config interface{}) bool {
	appCfg := reflect.ValueOf(config).Elem()
	return appCfg.FieldByName("UseVault").Interface().(bool)
}

func Run() int {
	app := new()
	vmware.RegisterExporter()

	if useVault(app.config) {
		app.logger.Info("Use Vault as credential storage for VMWare authentication")
		vaultClient, err := vault.NewClient(app.config, &app.logger)
		if err != nil {
			app.logger.Error(err.Error())
		}

		app.vaultClient = vaultClient

		go func() {
			for {
				app.vaultClient.KeepAlive()
			}
		}()
	}

	app.start()

	shutdownChan := make(chan os.Signal, 1)
	hupChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, syscall.SIGABRT, syscall.SIGQUIT, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(hupChan, syscall.SIGHUP)

	for {
		select {
		case <-hupChan:
			app.config = config.GetConfig()
			app.logger.Info("Reload config from env success")
		case <-shutdownChan:
			app.shutdown()
			return 0
		}
	}

}

func new() *app {
	cfg := config.GetConfig()
	logger := logging.GetLogger()
	logger.Info(""+
		"Start application...",
		logger.String("version", version.Version),
		logger.String("build_time", version.BuildTime),
		logger.String("commit", version.Commit),
	)
	logger.Info("Application logger initialized.")
	router := mux.NewRouter()
	logger.Info("Application router initialized.")

	return &app{
		logger:      logger,
		config:      cfg,
		appRouter:   router,
		appSrv:      nil,
		vaultClient: nil,
	}
}

func (a *app) start() {
	a.parseArgs()
	a.startAppHTTPServer()
}

func (a *app) startAppHTTPServer() {
	exporterHandler := vmware.GetHandler(&a.logger, a.config, a.vaultClient)
	exporterHandler.Register(a.appRouter)

	cfg := reflect.ValueOf(a.config).Elem()

	bindAddr := cfg.FieldByName("BindAddr").Interface().(string)

	if len(bindAddr) < 1 {
		bindAddr = ":9513"
	}

	a.logger.Info("Starting server...", a.logger.String("bind_addr", bindAddr))

	wtimeout, err := parseSeconds(cfg.FieldByName("HTTPWriteTimeout").Interface().(string), 30*time.Second)
	if err != nil {
		a.logger.Error(err.Error())
		a.logger.Info("set default write timeout")
	}

	rtimeout, err := parseSeconds(cfg.FieldByName("HTTPReadTimeout").Interface().(string), 30*time.Second)
	if err != nil {
		a.logger.Error(err.Error())
		a.logger.Info("set default read timeout")
	}

	srv := &http.Server{
		Handler:      a.appRouter,
		Addr:         bindAddr,
		WriteTimeout: wtimeout,
		ReadTimeout:  rtimeout,
	}

	go func(s *http.Server) {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.fatalServer(err)
		}
	}(srv)

	a.appSrv = srv
}

func (a *app) shutdown() {
	a.logger.Info("Shutdown Application...")
	ctx, serverCancel := context.WithTimeout(context.Background(), 15*time.Second)
	err := a.appSrv.Shutdown(ctx)
	if err != nil {
		a.fatalServer(err)
	}
	serverCancel()
	a.logger.Info("Application successful shutdown")
}

func (a *app) fatalServer(err error) {
	a.logger.Fatal(err.Error())
}
