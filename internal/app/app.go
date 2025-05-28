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

	"github.com/erminson/tasker"
	"github.com/erminson/tasker/internal/config"
	"github.com/erminson/tasker/internal/rest"
	"github.com/erminson/tasker/pkg/crypto"
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
		a.initMigration,
		a.initAdmin,
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
	// TODO: refactor
	a.sp.authMid = rest.NewAuthMiddleware(a.sp.JWTConfig().Secret())

	return nil
}

func (a *App) initHTTPRouter(ctx context.Context) error {
	if a.sp.authMid == nil {
		return fmt.Errorf("auth middleware is nil")
	}

	router := mux.NewRouter()

	v1public := router.PathPrefix("/api/v1/").Subrouter()
	v1 := router.PathPrefix("/api/v1/").Subrouter()
	v1.Use(a.sp.authMid.Middleware)

	userApi := a.sp.UserApi(ctx)
	// TODO: transfer setting router to method
	v1public.HandleFunc("/login", userApi.Login).Methods(http.MethodPost)

	v1public.HandleFunc("/users/leaderboard", userApi.LeaderBoard).Methods(http.MethodGet)
	v1.HandleFunc("/users/{id}/task/complete", userApi.CompleteTask).Methods(http.MethodPost)
	v1.HandleFunc("/users/{id}/referrer", userApi.Referrer).Methods(http.MethodPost)
	v1.HandleFunc("/users", userApi.CreateUser).Methods(http.MethodPost)
	v1.HandleFunc("/users/{id}", userApi.UpdateUser).Methods(http.MethodPatch)
	v1.HandleFunc("/users/{id}/status", userApi.UserInfo).Methods(http.MethodGet)

	a.router = router

	return nil
}

func (a *App) initMigration(ctx context.Context) error {
	cl := a.sp.DBClient(ctx)
	return cl.ApplyMigrations(tasker.Migrations)
}

func (a *App) initAdmin(ctx context.Context) error {
	repo := a.sp.UserRepository(ctx)

	// TODO: use tx for Count and Save methods
	count, err := repo.Count(ctx)
	if err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	cfg, err := config.NewAdminConfig()
	if err != nil {
		return err
	}

	login, pass := cfg.Credentials()
	passHash := crypto.BCrypto(pass)

	return repo.Save(ctx, login, passHash)
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
