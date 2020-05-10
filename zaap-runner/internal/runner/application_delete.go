package runner

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/runnerpb"
	"github.com/sirupsen/logrus"
)

func (r *Runner) DeleteApplication(_ context.Context, req *runnerpb.DeleteApplicationRequest) (*runnerpb.DeleteApplicationReply, error) {
	log := logrus.WithField("application-id", req.Id).WithField("application-name", req.Name)
	log.Info("deletion requested")

	if err := r.client.DeploymentDelete(req.Name); err != nil {
		log.WithError(err).Error("failed to delete deployment")
	}

	if err := r.client.ServiceDelete(req.Name); err != nil {
		log.WithError(err).Error("failed to delete service")
	}

	if err := r.client.IngressDelete(req.Name); err != nil {
		log.WithError(err).Error("failed to delete ingress")
	}

	if err := r.client.ClusterRoleBindingDeleteAll(req.Id, req.Name); err != nil {
		log.WithError(err).Error("failed to delete cluster role binding")
	}

	if err := r.client.ServiceAccountDelete(req.Name); err != nil {
		log.WithError(err).Error("failed to delete service account")
	}

	return &runnerpb.DeleteApplicationReply{}, nil
}
