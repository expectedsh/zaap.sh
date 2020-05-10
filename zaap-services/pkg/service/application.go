package service

import (
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/connector/rabbitmq"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/messaging"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/protocol"
)

var ApplicationEventsExchange = messaging.NewDurableExchangeTopic("application_events")

type applicationService struct {
	amqpConn  *rabbitmq.Connection
	publisher *messaging.Publisher
}

func NewApplicationService(amqpConn *rabbitmq.Connection) core.ApplicationService {
	return &applicationService{
		amqpConn:  amqpConn,
		publisher: messaging.NewPublisher(amqpConn, ApplicationEventsExchange),
	}
}

func (s applicationService) Deploy(application *core.Application) error {
	return s.publisher.Publish(messaging.DeliveryModeTransient, &protocol.ApplicationDeploymentRequested{
		Id:           application.ID.String(),
		DeploymentId: application.CurrentDeploymentID.String(),
	})
}

func (s applicationService) NotifyCreated(application *core.Application) error {
	return s.publisher.Publish(messaging.DeliveryModeTransient, &protocol.ApplicationCreated{
		Id:           application.ID.String(),
		DeploymentId: application.CurrentDeploymentID.String(),
	})
}

func (s applicationService) NotifyUpdated(application *core.Application) error {
	return s.publisher.Publish(messaging.DeliveryModeTransient, &protocol.ApplicationUpdated{
		Id:           application.ID.String(),
		DeploymentId: application.CurrentDeploymentID.String(),
	})
}

func (s applicationService) NotifyDeleted(application *core.Application) error {
	return s.publisher.Publish(messaging.DeliveryModeTransient, &protocol.ApplicationDeleted{
		Id:            application.ID.String(),
		RunnerId:      application.RunnerID.String(),
		Name:          application.Name,
		DefaultDomain: application.DefaultDomain,
	})
}

func (s applicationService) NotifyStatusChanged(application *core.Application) error {
	return s.publisher.Publish(messaging.DeliveryModeTransient, &protocol.ApplicationStatusChanged{
		Id:     application.ID.String(),
		Status: application.Status.ToMessagingFormat(),
	})
}
