package main

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/deployment-manager"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/deployment-manager/config"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"time"
)

func main()  {
	logrus.Infof("starting deployment manager")

	cfg, err := config.FromEnv()
	if err != nil {
		logrus.WithError(err).Fatal("could not parse configuration")
	}

	manager := deployment_manager.New(cfg)

	go func() {
		if err := manager.Start(); err != nil {
			logrus.WithError(err).Fatal("could not start deployment manager")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	logrus.Info("shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := manager.Shutdown(ctx); err != nil {
		logrus.WithError(err).Fatal("could not gracefully shutdown the deployment manager")
	}
}
