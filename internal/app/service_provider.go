package app

import (
	"context"
	"log"
	"log/slog"

	"github.com/erminson/tasker/internal/config"
	"github.com/erminson/tasker/internal/repository"
	usersRepo "github.com/erminson/tasker/internal/repository/user"
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
		s.userRepo = usersRepo.NewUserRepository(s.DBClient(ctx))
	}

	return s.userRepo
}
