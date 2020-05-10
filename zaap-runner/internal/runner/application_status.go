package runner

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/runnerpb"
	"github.com/sirupsen/logrus"
)

func (r *Runner) GetApplicationStatus(ctx context.Context, req *runnerpb.GetApplicationStatusRequest) (*runnerpb.GetApplicationStatusResponse, error) {
	log := logrus.
		WithField("application-id", req.Id).
		WithField("deployment-id", req.DeploymentId).
		WithField("application-name", req.Name)
	log.Debug("status requested")

	status, err := r.client.GetStatus(req.Id, req.DeploymentId)
	if err != nil {
		log.WithError(err).Error("failed to get application status")
		return nil, err
	}

	return &runnerpb.GetApplicationStatusResponse{
		Status: status,
	}, nil
}
