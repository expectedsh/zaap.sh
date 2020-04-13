package main

import (
	"context"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/docker/docker/client"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"

	"github.com/remicaumette/zaap.sh/zaap-services/internal/daemon"
	"github.com/remicaumette/zaap.sh/zaap-services/pkg/configs"
)

func main() {

	logrus.Info("Starting 'daemon' ...")

	daemonConfig := configs.Daemon{}
	if err := envconfig.Process("", &daemonConfig); err != nil {
		logrus.Panic(err)
	}

	controllerWsUrl := url.URL{Scheme: "ws", Host: daemonConfig.DaemonProxyAddress, Path: "/"}

	dockerClient, err := client.NewEnvClient()
	if err != nil {
		logrus.WithError(err).Panic("unable to create docker client")
	}

	daemonCtx, daemonExit := context.WithCancel(context.Background())

	go func() {
		err := daemon.RegisterControllerConsumer(daemonCtx, controllerWsUrl, dockerClient)
		if err != nil {
			logrus.WithError(err).Panic("unable to communicate with controller websocket")
		}
	}()

	go func() {
		dockerEventConsumer, err := daemon.NewDockerEventConsumer(daemonCtx, daemonConfig, dockerClient)
		if err != nil {
			logrus.WithError(err).Panic("unable to create docker event consumer")
		}

		if err := dockerEventConsumer.Listen(); err != nil {
			logrus.WithError(err).Panic()
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)

	<-stop
	daemonExit()
	time.Sleep(1 * time.Second)
}
