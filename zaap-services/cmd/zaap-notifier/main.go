package main

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/notifier"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/notifier/config"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"time"
)

func main() {
	logrus.Info("starting notifier")

	cfg, err := config.FromEnv()
	if err != nil {
		logrus.WithError(err).Fatal("could not parse configuration")
	}

	server := notifier.New(cfg)

	go func() {
		if err := server.Start(); err != nil {
			logrus.WithError(err).Fatal("could not start")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	logrus.Info("shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logrus.WithError(err).Fatal("could not gracefully shutdown")
	}
}
