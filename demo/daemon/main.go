package main

import (
  "context"
  "github.com/docker/docker/api/types"
  "github.com/docker/docker/api/types/swarm"
  "github.com/docker/docker/client"
  "github.com/remicaumette/zaap.sh/demo/protocol"
  "github.com/sirupsen/logrus"
  "google.golang.org/grpc"
  "google.golang.org/grpc/metadata"
)

func deployApp(docker *client.Client, r *protocol.AppDeploymentRequest) error {
  replicas := uint64(1)
  resources := &swarm.Resources{}
  _, err := docker.ServiceCreate(context.Background(), swarm.ServiceSpec{
    Annotations: swarm.Annotations{
      Name: r.AppName,
    },
    TaskTemplate: swarm.TaskSpec{
      ContainerSpec: swarm.ContainerSpec{
        Image: r.Repository,
      },
      Resources: &swarm.ResourceRequirements{
        Limits:       resources,
        Reservations: resources,
      },
    },
    Mode: swarm.ServiceMode{
      Replicated: &swarm.ReplicatedService{
        Replicas: &replicas,
      },
    },
  }, types.ServiceCreateOptions{})
  return err
}

func main() {
  docker, err := client.NewEnvClient()
  if err != nil {
    logrus.WithError(err).Fatal("could not connect to docker")
    return
  }
  defer docker.Close()
  addr := ":8090"
  conn, err := grpc.Dial(addr, grpc.WithInsecure())
  if err != nil {
    logrus.WithError(err).Fatalf("could not dial with %v", addr)
    return
  }
  defer conn.Close()
  scheduler := protocol.NewSchedulerClient(conn)
  logrus.Info("connection open, trying to obtain a token")
  tokenReq, err := scheduler.GetToken(context.Background(), &protocol.GetTokenRequest{})
  if err != nil {
    logrus.WithError(err).Fatal("failed to obtain a token")
    return
  }
  logrus.WithField("token", tokenReq.Token).Info("ok")
  ctx := metadata.AppendToOutgoingContext(context.Background(), "token", tokenReq.Token)
  events, err := scheduler.DeploymentEvents(ctx)
  if err != nil {
    logrus.WithError(err).Fatal("could not get deployment events")
    return
  }
  for {
    event, err := events.Recv()
    if err != nil {
      logrus.WithError(err).Fatal("could not read event")
      break
    }
    switch event.Type {
    case protocol.DeploymentEventRequestType_DEPLOY_APP:
      logrus.Info("deployment requested")
      if err := deployApp(docker, event.GetAppDeploymentRequest()); err != nil {
        events.Send(&protocol.DeploymentEventResponse{
          Type:    protocol.DeploymentEventResponseType_ERROR,
          Message: err.Error(),
        })
      } else {
        events.Send(&protocol.DeploymentEventResponse{
          Type: protocol.DeploymentEventResponseType_OK,
        })
      }
      break
    }
  }
}
