package app

import (
	"context"
	"log"
	"log/slog"

	apiUser "github.com/erminson/tasker/internal/api/user"
	"github.com/erminson/tasker/internal/config"
	"github.com/erminson/tasker/internal/repository"
	taskRepo "github.com/erminson/tasker/internal/repository/task"
	userRepo "github.com/erminson/tasker/internal/repository/user"
	"github.com/erminson/tasker/internal/service"
	"github.com/erminson/tasker/internal/service/task"
	"github.com/erminson/tasker/internal/service/user"
	database "github.com/erminson/tasker/pkg/db"
	"github.com/erminson/tasker/pkg/logger"
)

type serviceProvider struct {
	log *slog.Logger

	pgCfg    config.PGConfig
	httpCfg  config.HTTPConfig
	adminCfg config.AdminConfig

	db database.Driver

	userRepo repository.UserRepository
	taskRepo repository.TaskRepository

	userService service.UserService
	taskService service.TaskService

	userApi *apiUser.Implementation
}

func newServiceProvider(log *slog.Logger) *serviceProvider {
	return &serviceProvider{
		log: log,
	}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgCfg == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to load pg config: %s", err.Error())
			return nil
		}

		s.pgCfg = cfg
	}

	return s.pgCfg
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpCfg == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to load http config: %s", err)
			return nil
		}

		s.httpCfg = cfg
	}

	return s.httpCfg
}

func (s *serviceProvider) AdminConfig() config.PGConfig {
	if s.adminCfg == nil {
		cfg, err := config.NewAdminConfig()
		if err != nil {
			log.Fatalf("failed to load admin config: %s", err.Error())
			return nil
		}

		s.adminCfg = cfg
	}

	return s.pgCfg
}

func (s *serviceProvider) DBClient(_ context.Context) database.Driver {
	if s.db == nil {
		cl := database.NewDatabase(s.log, s.PGConfig().DSN())

		err := cl.Ping()
		if err != nil {
			s.log.Error("ping", logger.Err(err))
		}

		s.db = cl
	}

	return s.db
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepo == nil {
		s.userRepo = userRepo.NewRepository(s.DBClient(ctx))
	}

	return s.userRepo
}
func (s *serviceProvider) TaskRepository(ctx context.Context) repository.TaskRepository {
	if s.taskRepo == nil {
		s.taskRepo = taskRepo.NewRepository(s.DBClient(ctx))
	}

	return s.taskRepo
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = user.NewService(s.UserRepository(ctx))
	}

	return s.userService
}

func (s *serviceProvider) TaskService(ctx context.Context) service.TaskService {
	if s.taskService == nil {
		s.taskService = task.NewService(s.TaskRepository(ctx))
	}

	return s.taskService
}

func (s *serviceProvider) UserApi(ctx context.Context) *apiUser.Implementation {
	if s.userApi == nil {
		s.userApi = apiUser.NewApi(s.UserService(ctx), s.TaskService(ctx))
	}

	return s.userApi
}
