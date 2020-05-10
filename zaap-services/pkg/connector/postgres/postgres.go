package postgres

import (
	"github.com/jinzhu/gorm"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

type Config struct {
	PostgresURL string `envconfig:"POSTGRES_URL" default:"postgres://zaap:zaap@localhost/zaap?sslmode=disable"`
}

func ConfigFromEnv() (*Config, error) {
	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}
	return config, nil
}

func Connect(config *Config) (*gorm.DB, error) {
	if config == nil {
		c, err := ConfigFromEnv()
		if err != nil {
			return nil, err
		}
		config = c
	}

	return gorm.Open("postgres", config.PostgresURL)
}
