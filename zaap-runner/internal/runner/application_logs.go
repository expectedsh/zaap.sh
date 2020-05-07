package runner

import (
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/protocol"
	"github.com/sirupsen/logrus"
)

func (r *Runner) GetApplicationLogs(req *protocol.GetApplicationLogsRequest, srv protocol.Runner_GetApplicationLogsServer) error {
	log := logrus.WithField("application", req.Id)
	log.Info("getting logs application")

	logs, err := r.client.Logs(srv.Context(), req.Id)
	if err != nil {
		return err
	}

	for log := range logs {
		err = srv.Send(&protocol.GetApplicationLogsResponse{
			Time:    log.Time.String(),
			Pod:     log.Pod,
			Message: log.Message,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
