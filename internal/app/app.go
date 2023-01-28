package app

import (
	"context"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/gorilla/mux"
	"github.com/hashicorp/vault/api"
	"net/http"
	"os"
	"os/signal"
	"reflect"
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
	tokenTTL    int
	vaultClient *api.Client
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
		vaultClient, err := vault.NewClient(app.config)
		if err != nil {
			app.logger.Error(err.Error())
		}

		//app.tokenTTL = vaultClient.GetTokenTTL()
		//app.vaultClient = vaultClient.GetClient()

		app.tokenTTL = vaultClient.Ttl
		app.vaultClient = vaultClient.Client

		interval := time.Duration(app.tokenTTL) * time.Second
		nextTime := time.Now().Add(time.Since(time.Now()) + interval)

		go func() {
			app.logger.Info(
				fmt.Sprintf("Start job for renew Vault token every: %d seconds", app.tokenTTL),
				app.logger.String("vault_addr", app.vaultClient.Address()),
			)
			s := gocron.NewScheduler(time.Local)
			_, err := s.Every(app.tokenTTL).Second().Do(func() {
				app.logger.Info(fmt.Sprintf("Renew Vault Token by TTL. TTL is: %d", app.tokenTTL))
				vaultClient, err := vault.NewClient(app.config)
				if err != nil {
					app.logger.Error(err.Error())
				}
				//app.tokenTTL = vaultClient.GetTokenTTL()
				//app.vaultClient = vaultClient.GetClient()
				app.tokenTTL = vaultClient.Ttl
				app.vaultClient = vaultClient.Client
			})
			if err != nil {
				app.logger.Error(err.Error())
			}
			s.StartAt(nextTime)
			s.StartBlocking()
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
		logger:    logger,
		config:    cfg,
		appRouter: router,
		appSrv:    nil,
		//vaultToken: "",
		tokenTTL:    0,
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
