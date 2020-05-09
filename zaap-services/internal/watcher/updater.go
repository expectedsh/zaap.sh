package watcher

import (
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/protocol"
	"github.com/sirupsen/logrus"
	"time"
)

func (s *Server) updateRunner(runner core.Runner) {
	log := logrus.WithField("runner-id", runner.ID)
	oldStatus := runner.Status

	runner.Status = core.RunnerStatusOnline
	defer func() {
		if oldStatus != runner.Status {
			if err := s.runnerStore.Update(s.context, &runner); err != nil {
				log.WithError(err).Error("failed to update runner status")
			}
		}
	}()

	client, conn, err := s.runnerService.NewConnection(&runner)
	if err != nil {
		runner.Status = core.RunnerStatusOffline
		return
	}
	defer conn.Close()

	_, err = client.Ping(s.context, &protocol.PingRequest{Time: time.Now().Unix()})
	if err != nil {
		runner.Status = core.RunnerStatusOffline
		return
	}

	applications, err := s.applicationStore.ListByRunner(s.context, runner.ID)
	if err != nil {
		log.WithError(err).Error("failed to list applications by runner")
		return
	}

	for _, application := range *applications {
		appLog := log.WithField("application-id", application.ID)
		appOldStatus := application.Status

		res, err := client.GetApplicationStatus(s.context, &protocol.GetApplicationStatusRequest{
			Id:           application.ID.String(),
			DeploymentId: application.CurrentDeploymentID.String(),
			Name:         application.Name,
		})
		if err != nil {
			application.Status = core.ApplicationStatusUnknown
		} else {
			application.Status = core.ApplicationStatusFromRunner(res.Status)
		}

		if appOldStatus != application.Status {
			if err = s.applicationStore.Update(s.context, &application); err != nil {
				appLog.WithError(err).Error("failed to update runner status")
			}
		}
	}
}
