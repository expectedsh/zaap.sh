package configs

type DaemonProxy struct {
	RabbitMQUrl string `default:"amqp://localhost"`
	Address     string `default:"localhost:9090"`
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
