package scheduler

import (
	"context"
	"github.com/remicaumette/zaap.sh/zaap-scheduler/pkg/docker"
	"github.com/remicaumette/zaap.sh/zaap-scheduler/pkg/protocol"
	"github.com/sirupsen/logrus"
	"os"
)

type Server struct {
	dockerClient *docker.Docker
}

func New(dockerClient *docker.Docker) *Server {
	return &Server{
		dockerClient: dockerClient,
	}
}

func (s *Server) TestConnection(_ context.Context, r *protocol.TestConnectionRequest) (*protocol.TestConnectionResponse, error) {
	return &protocol.TestConnectionResponse{
		Ok: os.Getenv("SCHEDULER_TOKEN") == r.Token,
	}, nil
}

func (s *Server) DeployApplication(ctx context.Context, r *protocol.DeployApplicationRequest) (*protocol.DeployApplicationResponse, error) {
	log := logrus.WithField("application", r.Application.Id)
	log.Info("deployment requested")
	currentApp, err := s.dockerClient.ServiceGetFromApplication(ctx, r.Application.Id)
	if err != nil {
		return nil, err
	}

	if currentApp == nil {
		log.Info("application does not exists, creating")
		err = s.dockerClient.ServiceCreate(ctx, docker.ConvertApplication(r.Application))
	} else {
		log.Info("application already exists, updating")
		err = s.dockerClient.ServiceUpdate(ctx, docker.ConvertApplication(r.Application), currentApp)
	}

	if err != nil {
		return nil, err
	}
	return &protocol.DeployApplicationResponse{}, nil
}

func (s *Server) DeleteApplication(ctx context.Context, r *protocol.DeleteApplicationRequest) (*protocol.DeleteApplicationResponse, error) {
	log := logrus.WithField("application", r.Id)
	log.Info("deletion requested")
	currentApp, err := s.dockerClient.ServiceGetFromApplication(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	if currentApp == nil {
		log.Info("application does not exists")
		return nil, nil
	}
	if err := s.dockerClient.ServiceDelete(ctx, currentApp); err != nil {
		return nil, err
	}
	return &protocol.DeleteApplicationResponse{}, nil
}

func (s *Server) GetApplicationLogs(r *protocol.GetApplicationLogsRequest, srv protocol.Scheduler_GetApplicationLogsServer) error {
	return nil
}
