package config

import (
	"fmt"
	"os"
)

const (
	jwtSecretEnvName = "JWT_SECRET"
)

type JWTConfig interface {
	Secret() string
}
type jwtConfig struct {
	secret string
}

func NewJWTConfig() (JWTConfig, error) {
	secret := os.Getenv(jwtSecretEnvName)
	if secret == "" {
		return nil, fmt.Errorf("environment variable %s is not set", jwtSecretEnvName)
	}

	return &jwtConfig{
		secret: secret,
	}, nil
}

func (c *jwtConfig) Secret() string {
	return c.secret
}
