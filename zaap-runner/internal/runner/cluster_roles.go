package runner

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-runner/pkg/runnerpb"
	"github.com/sirupsen/logrus"
)

func (r *Runner) GetClusterRoles(_ context.Context, _ *runnerpb.GetClusterRolesRequest) (*runnerpb.GetClusterRolesReply, error) {
	logrus.Info("cluster roles requested")

	roles, err := r.client.ClusterRoleList()
	if err != nil {
		logrus.WithError(err).Error("failed to get cluster roles")
		return nil, err
	}

	return &runnerpb.GetClusterRolesReply{
		Roles: roles,
	}, nil
}
