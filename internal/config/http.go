package config

import (
	"fmt"
	"net"
	"os"
)

const (
	httpHostEnvName = "HTTP_HOST"
	httpPortEnvName = "HTTP_PORT"
)

type HTTPConfig interface {
	Address() string
}
type httpConfig struct {
	host string
	port string
}

func NewHTTPConfig() (HTTPConfig, error) {
	host := os.Getenv(httpHostEnvName)
	if host == "" {
		return nil, fmt.Errorf("environment variable %s is not set", httpHostEnvName)
	}

	port := os.Getenv(httpPortEnvName)
	if port == "" {
		return nil, fmt.Errorf("environment variable %s is not set", httpPortEnvName)
	}

	return &httpConfig{
		host: host,
		port: port,
	}, nil
}

func (c *httpConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
