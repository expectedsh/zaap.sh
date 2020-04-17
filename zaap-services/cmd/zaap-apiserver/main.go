package main

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/internal/apiserver"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"time"
)

func main() {
	config, err := apiserver.ConfigFromEnv()
	if err != nil {
		logrus.WithError(err).Fatal("could not parse configuration")
	}

	server := apiserver.New(config)

	go func() {
		logrus.Infof("listening on %v", config.Addr)
		if err := server.Start(); err != nil {
			logrus.WithError(err).Fatal("could not start apiserver")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	logrus.Info("server is shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logrus.WithError(err).Fatal("could not gracefully shutdown the server")
	}

	logrus.Info("server stopped")
}
