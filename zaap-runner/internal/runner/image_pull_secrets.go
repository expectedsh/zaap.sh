package runner

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/runnerpb"
	"github.com/sirupsen/logrus"
)

func (r *Runner) GetImagePullSecrets(_ context.Context, _ *runnerpb.GetImagePullSecretsRequest) (*runnerpb.GetImagePullSecretsReply, error) {
	logrus.Info("image pull secrets requested")

	secrets, err := r.client.SecretImagePullList()
	if err != nil {
		logrus.WithError(err).Error("failed to get image pull secrets")
		return nil, err
	}

	return &runnerpb.GetImagePullSecretsReply{
		Secrets: secrets,
	}, nil
}
