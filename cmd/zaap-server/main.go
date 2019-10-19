package main

import (
  "context"
  "github.com/sirupsen/logrus"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "time"
)

func main() {
  logrus.Info("starting zaap-server")
  srv, err := newServerFromEnv()
  if err != nil {
    logrus.WithError(err).Fatal("failed to start the server")
  }

  done := make(chan os.Signal, 1)
  signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
  go func() {
    if err := srv.Start(); err != nil && err != http.ErrServerClosed {
      logrus.WithError(err).Fatal("failed to start the server")
    }
  }()
  logrus.Info("listening on 3000")
  <-done

  logrus.Info("stopping server")
  ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
  if err := srv.Stop(ctx); err != nil {
    logrus.WithError(err).Fatal("failed to stop the server")
  }
}
