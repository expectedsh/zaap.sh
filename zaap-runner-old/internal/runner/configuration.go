package runner

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/protocol"
	"github.com/sirupsen/logrus"
)

func (r *Runner) GetConfiguration(ctx context.Context, req *protocol.GetConfigurationRequest) (*protocol.GetConfigurationResponse, error) {
	logrus.Info("configuration requested")
	return &protocol.GetConfigurationResponse{
		Type:        protocol.RunnerType_DOCKER_SWARM,
		ExternalIps: r.config.ExternalIps,
	}, nil
}
