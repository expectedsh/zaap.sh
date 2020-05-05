package runner

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/protocol"
	"github.com/sirupsen/logrus"
)

func (r *Runner) DeployApplication(_ context.Context, req *protocol.DeployApplicationRequest) (*protocol.DeployApplicationResponse, error) {
	log := logrus.WithField("application", req.Application.Id)
	log.Info("deployment requested")
	currentApp, err := r.client.DeploymentGet(req.Application)
	if err != nil {
		return nil, err
	}

	if currentApp == nil {
		log.Info("application does not exists, creating")
		err = r.client.DeploymentCreate(req.Application)
	} else {
		log.Info("application already exists, updating")
		err = r.client.DeploymentUpdate(req.Application)
	}

	if err != nil {
		return nil, err
	}
	return &protocol.DeployApplicationResponse{}, nil
}
