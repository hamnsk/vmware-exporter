package app

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"
	"vmware-exporter/internal/config"
	"vmware-exporter/internal/version"
	"vmware-exporter/internal/vmware"
	"vmware-exporter/pkg/logging"
)

type app struct {
	logger    logging.Logger
	config    interface{}
	appRouter *mux.Router
	appSrv    *http.Server
}

func Run() int {
	app := new()
	vmware.RegisterExporter()
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
		logger:    logger,
		config:    cfg,
		appRouter: router,
		appSrv:    nil,
	}
}

func (a *app) start() {
	a.parseArgs()
	a.startAppHTTPServer()
}

func (a *app) startAppHTTPServer() {
	exporterHandler := vmware.GetHandler(&a.logger, a.config)
	exporterHandler.Register(a.appRouter)

	cfg := reflect.ValueOf(a.config).Elem()

	bindAddr := cfg.FieldByName("BindAddr").Interface().(string)

	if len(bindAddr) < 1 {
		bindAddr = ":9513"
	}

	a.logger.Info("Starting server...", a.logger.String("bind_addr", bindAddr))

	wtimeout, err := time.ParseDuration(cfg.FieldByName("HTTPWriteTimeout").Interface().(string))
	if err != nil {
		a.logger.Error(err.Error())
		a.logger.Info("set default write timeout")
		wtimeout = 30
	}

	rtimeout, err := time.ParseDuration(cfg.FieldByName("HTTPReadTimeout").Interface().(string))
	if err != nil {
		a.logger.Error(err.Error())
		a.logger.Info("set default read timeout")
		rtimeout = 30
	}

	srv := &http.Server{
		Handler:      a.appRouter,
		Addr:         bindAddr,
		WriteTimeout: wtimeout * time.Second,
		ReadTimeout:  rtimeout * time.Second,
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
