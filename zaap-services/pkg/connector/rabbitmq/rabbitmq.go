package rabbitmq

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/streadway/amqp"
)

type Config struct {
	RabbitURL string `envconfig:"RABBIT_URL" default:"amqp://localhost/"`
}

func ConfigFromEnv() (*Config, error) {
	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}
	return config, nil
}

func Connect(config *Config) (*amqp.Connection, error) {
	if config == nil {
		c, err := ConfigFromEnv()
		if err != nil {
			return nil, err
		}
		config = c
	}

	conn, err := amqp.Dial(config.RabbitURL)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
