package config

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

type HTTPConfig interface {
	Address() string
}

type httpConfig struct {
	host string
	port string
}

func NewHTTPConfig(cfg *HTTPServerConfig) (HTTPConfig, error) {

	trimmedHost := strings.TrimSpace(cfg.Host)
	if trimmedHost == "" {
		return nil, errors.New("http host not found")
	}

	host := trimmedHost
	if len(host) == 0 {
		return nil, errors.New("http host is corrupted")
	}

	port := cfg.Port
	if port == 0 {
		return nil, errors.New("http port is corrupted")
	}
	portStr := strconv.Itoa(port)

	return &httpConfig{
		host: host,
		port: portStr,
	}, nil
}

func (cfg *httpConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
