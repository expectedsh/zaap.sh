package runner

import (
	"bufio"
	"errors"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/docker"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/protocol"
	"github.com/sirupsen/logrus"
	"time"
)

func (r *Runner) GetApplicationLogs(req *protocol.GetApplicationLogsRequest, srv protocol.Runner_GetApplicationLogsServer) error {
	log := logrus.WithField("application", req.Id)
	log.Info("getting logs application")

	currentApp, err := r.dockerClient.ServiceGetFromApplication(srv.Context(), req.Id)
	if err != nil {
		return err
	} else if currentApp == nil {
		return errors.New("application not found")
	}

	logs, err := r.dockerClient.ServiceGetLogs(srv.Context(), currentApp)
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
		Output:  output,
		TaskId:  reader.Labels["com.docker.swarm.task.id"],
		Time:    reader.Time.Format(time.RFC3339),
		Message: reader.Message,
	}
}
