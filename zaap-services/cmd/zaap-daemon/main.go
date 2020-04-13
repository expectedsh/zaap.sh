package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"

	"github.com/remicaumette/zaap.sh/zaap-services/internal/daemon"
)

func main() {
	logrus.Info("Starting 'daemon' ...")
	config := daemon.Config{}
	if err := envconfig.Process("", &config); err != nil {
		logrus.WithError(err).Fatal("could not process configuration")
	}
	d := daemon.New(config)

	ctx, exit := context.WithCancel(context.Background())
	if err := d.Start(ctx); err != nil {
		logrus.WithError(err).Fatal("could not start daemon")
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)

	<-stop
	exit()
	time.Sleep(1 * time.Second)
}
