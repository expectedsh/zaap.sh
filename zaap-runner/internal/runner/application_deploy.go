package runner

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/protocol"
	"github.com/sirupsen/logrus"
)

func (r *Runner) DeployApplication(_ context.Context, req *protocol.DeployApplicationRequest) (*protocol.DeployApplicationResponse, error) {
	log := logrus.WithField("application-id", req.Application.Id).WithField("application-name", req.Application.Name)
	log.Info("deployment requested")

	if err := r.client.DeploymentCreateOrUpdate(req.Application); err != nil {
		log.WithError(err).Error("failed to create/update deployment")
		return nil, err
	}

	if err := r.client.ServiceCreateOrUpdate(req.Application); err != nil {
		log.WithError(err).Error("failed to create/update service")
		return nil, err
	}

	if err := r.client.IngressCreateOrUpdate(req.Application); err != nil {
		log.WithError(err).Error("failed to create/update ingress")
		return nil, err
	}

	return &protocol.DeployApplicationResponse{}, nil
}
