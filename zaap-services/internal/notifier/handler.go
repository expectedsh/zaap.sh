package notifier

import (
	"context"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/protocol"
	uuid "github.com/satori/go.uuid"
)

func (s *Server) HandleApplicationDeploymentRequested(ctx context.Context, message *protocol.ApplicationDeploymentRequest) error {
	application, err := s.applicationStore.Find(ctx, uuid.FromStringOrNil(message.ApplicationId))
	if err != nil {
		return err
	}
	return s.notifier.WhenApplicationDeploymentRequest(application)
}

func (s *Server) HandleApplicationDeleted(ctx context.Context, message *protocol.ApplicationDeleted) error {
	return s.notifier.WhenApplicationDeleted(message.Id, message.Name)
}
