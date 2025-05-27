package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"github.com/erminson/tasker/internal/config"
	"github.com/erminson/tasker/pkg/server"
)

type App struct {
	sp     *serviceProvider
	router *mux.Router
	log    *slog.Logger
}

func New(ctx context.Context, log *slog.Logger) (*App, error) {
	app := &App{
		log: log,
	}

	err := app.init(ctx)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Run(ctx context.Context) error {
	a.log.Info("service starting...")
	defer a.log.Info("service stopped")

	return a.runHTTPServer(ctx)
}

func (a *App) init(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPRouter,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.sp = newServiceProvider(a.log)

	return nil
}

func (a *App) initHTTPRouter(_ context.Context) error {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(fmt.Sprint("Hello World: ", time.Now().Format(time.RFC850))))
	}).Methods(http.MethodGet)

	a.router = router

	return nil
}

func (a *App) runHTTPServer(ctx context.Context) error {
	if a.router == nil {
		return fmt.Errorf("http router is nil")
	}

	srv := server.NewServer(
		a.sp.HTTPConfig().Address(),
		a.router,
	)

	errCh := make(chan error, 1)
	go func() {
		err := srv.ListenAndServe()
		errCh <- err
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errCh:
		return err
	case <-stop:
		a.log.Warn("http server shutting down...")
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		return srv.Shutdown(ctx)
	}
}
