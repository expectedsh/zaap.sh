package main

import (
  "github.com/docker/docker/client"
  "github.com/remicaumette/zaap.sh/pkg/server"
  "github.com/remicaumette/zaap.sh/pkg/server/app"
  "github.com/sirupsen/logrus"
  "os"
)

func main() {
  addr := os.Getenv("ADDR")
  if addr == "" {
    addr = ":5200"
  }
  docker, err := client.NewEnvClient()
  if err != nil {
    logrus.Error(err)
    return
  }
  defer docker.Close()
  srv := server.New(app.NewService(docker))
  if err := srv.Start(addr); err != nil {
    logrus.Error(err)
  }
}
