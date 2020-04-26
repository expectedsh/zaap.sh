package deployer

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/protocol"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

func (s *Server) DeploymentHandler(ctx context.Context, message *protocol.ApplicationDeploymentRequest) error {
	logrus.
		WithField("application-id", message.Id).
		WithField("deployment-id", message.DeploymentId).
		Info("handling application deployment")
	return s.deployApplication(ctx, uuid.FromStringOrNil(message.Id), uuid.FromStringOrNil(message.DeploymentId))
}

func (s *Server) ApplicationCreatedHandler(ctx context.Context, message *protocol.ApplicationCreated) error {
	logrus.
		WithField("application-id", message.Id).
		WithField("deployment-id", message.DeploymentId).
		Info("handling application creation")
	return s.deployApplication(ctx, uuid.FromStringOrNil(message.Id), uuid.FromStringOrNil(message.DeploymentId))
}

func (s *Server) ApplicationUpdatedHandler(ctx context.Context, message *protocol.ApplicationUpdated) error {
	logrus.
		WithField("application-id", message.Id).
		WithField("deployment-id", message.DeploymentId).
		Info("handling application update")
	return s.deployApplication(ctx, uuid.FromStringOrNil(message.Id), uuid.FromStringOrNil(message.DeploymentId))
}

func (s *Server) ApplicationDeletedHandler(ctx context.Context, message *protocol.ApplicationDeleted) error {
	logrus.WithField("application-id", message.Id).Info("handling application deletion")
	return s.deleteApplication(ctx, uuid.FromStringOrNil(message.Id), uuid.FromStringOrNil(message.RunnerId))
}
