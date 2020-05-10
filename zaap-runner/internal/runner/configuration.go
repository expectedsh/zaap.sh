package runner

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/runnerpb"
	"github.com/sirupsen/logrus"
)

func (r *Runner) GetConfiguration(ctx context.Context, req *runnerpb.GetConfigurationRequest) (*runnerpb.GetConfigurationResponse, error) {
	logrus.Info("configuration requested")
	return &runnerpb.GetConfigurationResponse{
		ExternalIps: r.config.ExternalIps,
	}, nil
}
