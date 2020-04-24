package service

import (
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/core"
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/protocol"
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
	"reflect"
)

const ApplicationEventsExchange = "application_events"

type applicationService struct {
	amqpConn *amqp.Connection
}

func NewApplicationService(amqpConn *amqp.Connection) core.ApplicationService {
	return &applicationService{
		amqpConn: amqpConn,
	}
}

func (s applicationService) emitEvent(event proto.Message) error {
	messageType := reflect.TypeOf(event).Elem().Name()

	ch, err := s.amqpConn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(ApplicationEventsExchange, amqp.ExchangeTopic, true, false, false, false, nil)
	if err != nil {
		return err
	}

	payload, err := proto.Marshal(event)
	if err != nil {
		return err
	}

	return ch.Publish(ApplicationEventsExchange, messageType, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Body:         payload,
	})
}

func (s applicationService) Deploy(application *core.Application) error {
	return s.emitEvent(&protocol.ApplicationDeploymentRequest{
		ApplicationId: application.ID.String(),
		DeploymentId:  application.CurrentDeploymentID.String(),
	})
}

func (s applicationService) NotifyDeletion(application *core.Application) error {
	return s.emitEvent(&protocol.ApplicationDeleted{
		Id:   application.ID.String(),
		Name: application.Name,
	})
}
