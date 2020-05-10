package rabbitmq

import (
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/backoff"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
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

func Connect(config *Config) (*Connection, error) {
	if config == nil {
		c, err := ConfigFromEnv()
		if err != nil {
			return nil, err
		}
		config = c
	}

	rawConn, err := amqp.Dial(config.RabbitURL)
	if err != nil {
		return nil, err
	}

	conn := &Connection{rawConn}

	go func() {
		errors := conn.NotifyClose(make(chan *amqp.Error))
		for {
			amqpErr, ok := <-errors
			if !ok {
				break
			}
			err = backoff.New("rabbitmq connection closed, reconnecting...", func() error {
				rawConn, err = amqp.Dial(config.RabbitURL)
				if err != nil {
					return err
				}
				return nil
			}, logrus.WithError(amqpErr)).Run()
			if err != nil {
				logrus.WithError(err).Fatal("could not reconnect")
				return
			}
			conn.Connection = rawConn
		}
	}()

	return conn, nil
}
