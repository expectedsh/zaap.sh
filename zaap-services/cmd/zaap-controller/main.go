package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/remicaumette/zaap.sh/zaap-services/internal/controller"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("starting 'controller' ...")

	config := controller.Config{}
	if err := envconfig.Process("", &config); err != nil {
		logrus.WithError(err).Fatal("could not process configuration")
	}
	ctrl := controller.New(config)

	go func() {
		if err := ctrl.Start(); err != nil {
			if err != http.ErrServerClosed {
				logrus.WithError(err).Fatal("could not start the server")
			}
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := ctrl.Shutdown(ctx); err != nil {
		logrus.WithError(err).Fatal("failed to shutdown the server")
	}
}
