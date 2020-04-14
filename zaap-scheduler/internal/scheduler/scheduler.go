package scheduler

import (
	"bufio"
	"context"
	"errors"
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
	log := logrus.WithField("application", r.Id)
	log.Info("getting logs application")

	currentApp, err := s.dockerClient.ServiceGetFromApplication(srv.Context(), r.Id)
	if err != nil {
		return err
	} else if currentApp == nil {
		return errors.New("application not found")
	}

	logs, err := s.dockerClient.ServiceGetLogs(srv.Context(), currentApp)
	if err != nil {
		return err
	}
	defer logs.Close()

	reader := docker.NewLogReader(logs)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if err := srv.Send(readLogLine(reader)); err != nil {
			return err
		}
	}

	return scanner.Err()
}

func readLogLine(reader *docker.LogReader) *protocol.GetApplicationLogsResponse {
	output := protocol.GetApplicationLogsResponse_STDOUT
	if reader.Output == docker.OutputStderr {
		output = protocol.GetApplicationLogsResponse_STDERR
	}

	return &protocol.GetApplicationLogsResponse{
		Output: output,
		TaskId: reader.Labels["com.docker.swarm.task.id"],
		Time: &protocol.Timestamp{
			Second:     int64(reader.Time.Second()),
			NanoSecond: int64(reader.Time.Nanosecond()),
		},
		Message: reader.Message,
	}
}
