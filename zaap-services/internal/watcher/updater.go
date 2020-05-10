package watcher

import (
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/runnerpb"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
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
			} else if err := s.runnerService.NotifyStatusChanged(&runner); err != nil {
				log.WithError(err).Error("failed to notify runner status changed")
			}
		}
	}()

	client, conn, err := s.runnerService.NewConnection(&runner)
	if err != nil {
		runner.Status = core.RunnerStatusOffline
		return
	}
	defer conn.Close()

	_, err = client.Ping(s.context, &runnerpb.PingRequest{
		Time: time.Now().Unix(),
	})
	if err != nil {
		logrus.Info(err)
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

		res, err := client.GetApplicationStatus(s.context, &runnerpb.GetApplicationStatusRequest{
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
				appLog.WithError(err).Error("failed to update application status")
			} else if err = s.applicationService.NotifyStatusChanged(&application); err != nil {
				appLog.WithError(err).Error("failed to notify application status changed")
			}
		}
	}
}
