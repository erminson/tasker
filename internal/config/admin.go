package config

import (
	"fmt"
	"os"
)

const (
	adminLoginEnvName = "ADMIN_LOGIN"
	adminPassEnvName  = "ADMIN_PASS"
)

type AdminConfig interface {
	Credentials() (string, string)
}
type adminConfig struct {
	login string
	pass  string
}

func NewAdminConfig() (AdminConfig, error) {
	login := os.Getenv(adminLoginEnvName)
	if login == "" {
		return nil, fmt.Errorf("environment variable %s is not set", adminLoginEnvName)
	}

	pass := os.Getenv(adminPassEnvName)
	if pass == "" {
		return nil, fmt.Errorf("environment variable %s is not set", adminPassEnvName)
	}

	return &adminConfig{
		login: login,
		pass:  pass,
	}, nil
}

func (c *adminConfig) Credentials() (string, string) {
	return c.login, c.pass
}
