package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Addr        string `envconfig:"ADDR" default:":3000"`
	PostgresURL string `envconfig:"POSTGRES_URL" default:"postgres://zaap:zaap@localhost/zaap?sslmode=disable"`
	SecretKey   string `envconfig:"SECRET_KEY" default:"changeme"`
}

func FromEnv() (*Config, error) {
	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}
	return config, nil
}
