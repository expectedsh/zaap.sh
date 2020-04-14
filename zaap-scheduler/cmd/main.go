package main

import (
	"github.com/docker/docker/client"
	"github.com/remicaumette/zaap.sh/zaap-scheduler/internal/scheduler"
	"github.com/remicaumette/zaap.sh/zaap-scheduler/pkg/docker"
	"github.com/remicaumette/zaap.sh/zaap-scheduler/pkg/protocol"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func main() {
	dockerClient, err := client.NewEnvClient()
	if err != nil {
		logrus.WithError(err).Fatal("could not initialize docker client")
		return
	}

	addr := ":8090"
	server := grpc.NewServer()
	protocol.RegisterSchedulerServer(server, scheduler.New(&docker.Docker{Client: dockerClient}))

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
