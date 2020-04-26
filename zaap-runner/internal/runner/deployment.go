package runner

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/docker"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/protocol"
	"github.com/sirupsen/logrus"
)

func (r *Runner) DeployApplication(ctx context.Context, req *protocol.DeployApplicationRequest) (*protocol.DeployApplicationResponse, error) {
	log := logrus.WithField("application", req.Application.Id)
	log.Info("deployment requested")
	currentApp, err := r.dockerClient.ServiceGetFromApplication(ctx, req.Application.Id)
	if err != nil {
		return nil, err
	}

	if currentApp == nil {
		log.Info("application does not exists, creating")
		err = r.dockerClient.ServiceCreate(ctx, docker.ToSwarmSpec(r.config.TraefikNetwork, req.Application))
	} else {
		log.Info("application already exists, updating")
		err = r.dockerClient.ServiceUpdate(ctx, docker.ToSwarmSpec(r.config.TraefikNetwork, req.Application), currentApp)
	}

	if err != nil {

		return nil, err
	}
	return &protocol.DeployApplicationResponse{}, nil
}

func (r *Runner) DeleteApplication(ctx context.Context, req *protocol.DeleteApplicationRequest) (*protocol.DeleteApplicationResponse, error) {
	log := logrus.WithField("application", req.Id)
	log.Info("deletion requested")
	currentApp, err := r.dockerClient.ServiceGetFromApplication(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	if currentApp == nil {
		log.Info("application does not exists")
		return nil, nil
	}
	if err := r.dockerClient.ServiceDelete(ctx, currentApp); err != nil {
		return nil, err
	}
	return &protocol.DeleteApplicationResponse{}, nil
}
