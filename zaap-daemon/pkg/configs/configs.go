package configs

import "fmt"

// amqp://user:bitnami@localhost:5672
type DaemonProxy struct {
	RabbitMQ RabbitMQConfig
	Address  string `default:"localhost:9090"`
}

type Daemon struct {
	DaemonProxyAddress string `default:"localhost:9090"`
}

type RabbitMQConfig struct {
	User     string `default:"user"`
	Password string `default:"bitnami"`
	Host     string `default:"localhost"`
	Port     string `default:"5672"`
}

func (r RabbitMQConfig) Url() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s", r.User, r.Password, r.Host, r.Port)
}
