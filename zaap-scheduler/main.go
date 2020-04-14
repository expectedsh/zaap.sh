package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"github.com/remicaumette/zaap.sh/zaap-scheduler/protocol"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
)

type SchedulerServer struct {
	dockerClient *client.Client
}

func main() {
	dockerClient, err := client.NewEnvClient()
	if err != nil {
		logrus.WithError(err).Fatal("could not initialize docker client")
		return
	}

	addr := ":8090"
	server := grpc.NewServer()
	protocol.RegisterSchedulerServer(server, &SchedulerServer{dockerClient})

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

func (s *SchedulerServer) TestConnection(_ context.Context, r *protocol.TestConnectionRequest) (*protocol.TestConnectionResponse, error) {
	return &protocol.TestConnectionResponse{
		Ok: os.Getenv("SCHEDULER_TOKEN") == r.Token,
	}, nil
}

func (s *SchedulerServer) DeployApplication(ctx context.Context, r *protocol.DeployApplicationRequest) (*protocol.DeployApplicationResponse, error) {
	log := logrus.WithField("application", r.Id)
	log.Info("deployment requested")
	currentApp, err := s.getApplication(r.Id)
	if err != nil {
		return nil, err
	}

	if currentApp == nil {
		log.Info("application does not exists, creating")
		err = s.createApplication(r)
	} else {
		log.Info("application already exists, updating")
		err = s.updateApplication(r, currentApp)
	}

	if err != nil {
		return nil, err
	}
	return &protocol.DeployApplicationResponse{State: protocol.ApplicationState_STARTING}, nil
}

func (s *SchedulerServer) getApplication(id string) (*swarm.Service, error) {
	args := filters.NewArgs()
	args.Add("label", fmt.Sprintf("zaap-app-id=%s", id))
	services, err := s.dockerClient.ServiceList(context.Background(), types.ServiceListOptions{
		Filters: args,
	})
	if err != nil {
		return nil, err
	}
	if len(services) == 0 {
		return nil, nil
	}
	return &services[0], nil
}

func (s *SchedulerServer) createApplication(application *protocol.DeployApplicationRequest) error {
	var replicas = uint64(1)
	_, err := s.dockerClient.ServiceCreate(context.Background(), swarm.ServiceSpec{
		Annotations: swarm.Annotations{
			Name: application.Name + "_" + application.Id,
			Labels: map[string]string{
				"zaap-app-id":   application.Id,
				"zaap-app-name": application.Name,
			},
		},
		TaskTemplate: swarm.TaskSpec{
			ContainerSpec: swarm.ContainerSpec{
				Image: application.Name,
				Env:   convertEnvironmentVariables(application.Environment),
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

func (s *SchedulerServer) updateApplication(application *protocol.DeployApplicationRequest, service *swarm.Service) error {
	var replicas = uint64(1)
	_, err := s.dockerClient.ServiceUpdate(context.Background(), service.ID, service.Version, swarm.ServiceSpec{
		Annotations: swarm.Annotations{
			Name: application.Name + "_" + application.Id,
			Labels: map[string]string{
				"zaap-app-id":   application.Id,
				"zaap-app-name": application.Name,
			},
		},
		TaskTemplate: swarm.TaskSpec{
			ContainerSpec: swarm.ContainerSpec{
				Image: application.Name,
				Env:   convertEnvironmentVariables(application.Environment),
			},
		},
		Mode: swarm.ServiceMode{
			Replicated: &swarm.ReplicatedService{
				Replicas: &replicas,
			},
		},
	}, types.ServiceUpdateOptions{})
	return err
}

func convertEnvironmentVariables(environment map[string]string) []string {
	var env []string
	for key, value := range environment {
		env = append(env, fmt.Sprintf("%s=%s", key, value))
	}

	return env
}
