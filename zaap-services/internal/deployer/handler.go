package deployer

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/protocol"
	"github.com/sirupsen/logrus"
)

func (s *Server) DeploymentHandler(ctx context.Context, message *protocol.ApplicationDeploymentRequest) error {
	logrus.Info(message)
	return nil
}
