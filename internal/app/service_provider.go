package app

import (
	"log"
	"log/slog"

	"github.com/erminson/tasker/internal/config"
)

type serviceProvider struct {
	pgCfg   config.PGConfig
	httpCfg config.HTTPConfig
}

func newServiceProvider(log *slog.Logger) *serviceProvider {
	return &serviceProvider{}
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
