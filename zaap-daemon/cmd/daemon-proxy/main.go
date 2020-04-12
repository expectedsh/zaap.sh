package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"

	"github.com/remicaumette/zaap.sh/internal/proxy"
	"github.com/remicaumette/zaap.sh/pkg/configs"

	"github.com/kelseyhightower/envconfig"
)

func main() {

	logrus.Info("Starting 'daemon-proxy' ...")

	daemonProxyConfig := configs.DaemonProxy{}
	if err := envconfig.Process("", &daemonProxyConfig); err != nil {
		logrus.Panic(err)
	}

	rabbitConnection, err := amqp.Dial(daemonProxyConfig.RabbitMQ.Url())
	if err != nil {
		logrus.Panic(err)
	}

	dockerEventQueue, err := proxy.NewDockerEventQueue(rabbitConnection)
	if err != nil {
		logrus.Panic(err)
	}

	http.HandleFunc("/daemon", proxy.DaemonSocketHandler(dockerEventQueue))
	http.HandleFunc("/api", proxy.APISocketHandler)

	server := &http.Server{Addr: daemonProxyConfig.Address, Handler: http.DefaultServeMux}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				logrus.Panic(err)
			}
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logrus.WithField("closer", "http-server").Panic(err)
	}
	if err := rabbitConnection.Close(); err != nil {
		logrus.WithField("closer", "rabbit-mq").Panic(err)
	}
}
