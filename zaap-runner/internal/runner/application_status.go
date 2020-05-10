package runner

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/runnerpb"
	"github.com/sirupsen/logrus"
)

func (r *Runner) GetApplicationStatus(_ context.Context, req *runnerpb.GetApplicationStatusRequest) (*runnerpb.GetApplicationStatusReply, error) {
	log := logrus.
		WithField("application-id", req.Id).
		WithField("deployment-id", req.DeploymentId).
		WithField("application-name", req.Name)
	log.Debug("status requested")

	status, err := r.client.DeploymentStatus(req.Id, req.DeploymentId)
	if err != nil {
		log.WithError(err).Error("failed to get application status")
		return nil, err
	}

	return &runnerpb.GetApplicationStatusReply{
		Status: status,
	}, nil
}
