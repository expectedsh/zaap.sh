package main

import (
  "context"
  "github.com/remicaumette/zaap.sh/pkg/server"
  "github.com/sirupsen/logrus"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "time"
)

func main()  {
  logrus.Info("starting zaap-server")
  config, err := server.NewConfigFromEnviron()
  if err != nil {
    logrus.WithError(err).Fatal("invalid configuration")
  }
  serv := server.New(config)

  done := make(chan os.Signal, 1)
  signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

  go func() {
    if err := serv.Start(); err != nil && err != http.ErrServerClosed {
      logrus.WithError(err).Fatal("failed to start the server")
    }
  }()

  logrus.Info("listening on 3000")
  <-done
  logrus.Info("stopping server")
  ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
  if err := serv.Stop(ctx); err != nil {
    logrus.WithError(err).Fatal("failed to stop the server")
  }
}
