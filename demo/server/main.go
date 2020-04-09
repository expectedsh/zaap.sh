package main

import (
  "github.com/remicaumette/zaap.sh/demo/protocol"
  "github.com/remicaumette/zaap.sh/demo/server/registry"
  "github.com/remicaumette/zaap.sh/demo/server/scheduler"
  "github.com/sirupsen/logrus"
  "google.golang.org/grpc"
  "net"
)

func main() {
  server := grpc.NewServer()
  schedulerInstance := scheduler.New()
  registryInstance := registry.NewHookServer(schedulerInstance)
  protocol.RegisterSchedulerServer(server, schedulerInstance)
  go registryInstance.Start()
  addr := ":8090"
  lis, err := net.Listen("tcp", addr)
  if err != nil {
    logrus.WithError(err).Fatalf("could not listen on %v", addr)
    return
  }
  logrus.Infof("listening on %v", addr)
  if err = server.Serve(lis); err != nil {
    logrus.WithError(err).Fatal("failed to start the server")
  }
}
