package main

import (
	"github.com/expected.sh/zaap.sh/zaap-runner/internal/runner"
	"github.com/expected.sh/zaap.sh/zaap-runner/internal/runner/config"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func main() {
	cfg, err := config.FromEnv()
	if err != nil {
		logrus.WithError(err).Fatal("could not parse configuration")
		return
	}

	r := runner.New(cfg)

	go func() {
		if err := r.Start(); err != nil {
			logrus.WithError(err).Fatal("could not start runner")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	logrus.Info("shutting down")
	r.Shutdown()
}
