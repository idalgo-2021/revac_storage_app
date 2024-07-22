package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"revac_storage_app/internal/config"
)

type App struct {
	Config          config.Config
	serviceProvider *serviceProvider
	httpServer      *http.Server
	db              *sql.DB
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize dependencies: %w", err)
	}

	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initDB,
		a.initServiceProvider,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	var cfg config.Config
	err := config.LoadConfig(&cfg, "config.json")
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}
	a.Config = cfg

	return nil
}

func (a *App) initDB(_ context.Context) error {
	dbConn, err := config.NewDBConnection(&a.Config.DBConnectionConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	a.db = dbConn
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.db, a.Config)
	return nil
}

func (a *App) initHTTPServer(_ context.Context) error {
	handler := a.serviceProvider.Router()
	a.httpServer = &http.Server{
		Addr:    a.serviceProvider.HTTPConfig(&a.Config.HTTPServerConfig).Address(),
		Handler: handler,
	}
	return nil
}

func (a *App) Run() error {
	return a.runHTTPServer()
}

func (a *App) runHTTPServer() error {

	log.Printf("HTTP server is running on %s", a.serviceProvider.HTTPConfig(&a.Config.HTTPServerConfig).Address())

	if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to run HTTP server: %w", err)
	}

	return nil
}
