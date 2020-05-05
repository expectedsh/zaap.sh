package runner

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/protocol"
	"github.com/sirupsen/logrus"
)

func (r *Runner) DeleteApplication(_ context.Context, req *protocol.DeleteApplicationRequest) (*protocol.DeleteApplicationResponse, error) {
	log := logrus.WithField("application", req.Id)
	log.Info("deletion requested")
	if err := r.client.DeploymentDelete(&protocol.Application{Id: req.Id}); err != nil {
		return nil, err
	}
	return &protocol.DeleteApplicationResponse{}, nil
}
