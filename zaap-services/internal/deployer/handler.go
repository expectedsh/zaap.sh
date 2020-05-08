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
	log := logrus.WithField("application-id", message.Id).WithField("deployment-id", message.DeploymentId)
	log.Info("handling application creation")

	applicationId := uuid.FromStringOrNil(message.Id)

	if err := s.createDnsEntry(ctx, applicationId); err != nil {
		log.WithError(err).Error("could not create dns entries")
	}

	return s.deployApplication(ctx, applicationId, uuid.FromStringOrNil(message.DeploymentId))
}

func (s *Server) ApplicationUpdatedHandler(ctx context.Context, message *protocol.ApplicationUpdated) error {
	logrus.
		WithField("application-id", message.Id).
		WithField("deployment-id", message.DeploymentId).
		Info("handling application update")
	return s.deployApplication(ctx, uuid.FromStringOrNil(message.Id), uuid.FromStringOrNil(message.DeploymentId))
}

func (s *Server) ApplicationDeletedHandler(ctx context.Context, message *protocol.ApplicationDeleted) error {
	log := logrus.WithField("application-id", message.Id)
	log.Info("handling application deletion")

	applicationId := uuid.FromStringOrNil(message.Id)

	if err := s.deleteDnsEntry(message.DefaultDomain); err != nil {
		log.WithError(err).Error("could not delete dns entries")
	}

	return s.deleteApplication(ctx, applicationId, message.Name, uuid.FromStringOrNil(message.RunnerId))
}
