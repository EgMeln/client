// Package config to env
package config

import "github.com/caarlos0/env/v6"

// Config struct to grpc config env
type Config struct {
	PositionServicePort string `env:"POSITION_PORT" envDefault:":8083"`
	PriceServicePort    string `env:"PRICE_PORT" envDefault:":8089"`
}

// New contract config
func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
