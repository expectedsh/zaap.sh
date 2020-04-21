package service

import (
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/protocol"
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
)

const (
	DeploymentQueueName = "deployments"
	ApplicationEventsExchange = "application_events"
)

type applicationService struct {
	amqpConn *amqp.Connection
}

func NewApplicationService(amqpConn *amqp.Connection) core.ApplicationService {
	return &applicationService{amqpConn}
}

func (s applicationService) Deploy(application *core.Application) error {
	ch, err := s.amqpConn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(DeploymentQueueName, true, false, false, false, nil)
	if err != nil {
		return err
	}

	payload, err := proto.Marshal(&protocol.DeploymentRequest{
		ApplicationId: application.ID.String(),
		DeploymentId: application.CurrentDeploymentID.String(),
	})
	if err != nil {
		return nil
	}

	return ch.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Body: payload,
	})
}

func (s applicationService) NotifyDeletion(application *core.Application) error {
	ch, err := s.amqpConn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(ApplicationEventsExchange, amqp.ExchangeFanout, true, false, false, false, nil)
	if err != nil  {
		return err
	}

	payload, err := proto.Marshal(&protocol.ApplicationDeleted{
		ApplicationId: application.ID.String(),
	})
	if err != nil {
		return err
	}

	return ch.Publish(ApplicationEventsExchange, "", false, false, amqp.Publishing{
		Body: payload,
	})
}
