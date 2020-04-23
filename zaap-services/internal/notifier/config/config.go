package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	PostgresURL string `envconfig:"POSTGRES_URL" default:"postgres://zaap:zaap@localhost/zaap?sslmode=disable"`
	RabbitURL   string `envconfig:"RABBIT_URL" default:"amqp://localhost/"`
}

func FromEnv() (*Config, error) {
	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}
	return config, nil
}
