package configs

type Controller struct {
	RabbitMQUrl string `default:"amqp://localhost"`
	Address     string `default:"localhost:9090"`
}

type Daemon struct {
	DaemonProxyAddress string `default:"localhost:9090"`
	SchedulerToken     string `default:"jesuisuntoken"`
}
