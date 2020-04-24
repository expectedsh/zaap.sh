package messaging

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"reflect"
)

type (
	SubscriberHandler func(ctx context.Context, message proto.Message) error

	Subscriber struct {
		conn            *amqp.Connection
		exchangeName    string
		queueName       string
		messageRegistry map[string]reflect.Type
		Handler         SubscriberHandler
	}
)

func NewSubscriber(conn *amqp.Connection, exchangeName string, queueName string) *Subscriber {
	return &Subscriber{
		conn:            conn,
		exchangeName:    exchangeName,
		queueName:       queueName,
		messageRegistry: make(map[string]reflect.Type),
	}
}

func (s *Subscriber) RegisterMessage(message *proto.Message) {
	messageType := reflect.TypeOf(message).Elem()
	s.messageRegistry[messageType.Name()] = messageType
}

func (s *Subscriber) Subscribe(ctx context.Context) error {
	ch, err := s.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(s.exchangeName, amqp.ExchangeTopic, true, false, false, false, nil)
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(s.queueName, false, false, false, false, nil)
	if err != nil {
		return err
	}

	for messageType, _ := range s.messageRegistry {
		err = ch.QueueBind(q.Name, messageType, s.exchangeName, false, nil)
		if err != nil {
			return err
		}
	}

	err = ch.Qos(1, 0, false)
	if err != nil {
		return err
	}

	messages, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case message := <-messages:
			messageType := s.messageRegistry[message.RoutingKey]
			if messageType == nil {
				_ = message.Nack(false, true)
				continue
			}

			payload := reflect.New(messageType).Interface().(proto.Message)
			if err = proto.Unmarshal(message.Body, payload); err != nil {
				logrus.WithError(err).Warn("could not unmarshal message")
				_ = message.Nack(false, true)
				continue
			}

			if err = s.Handler(ctx, payload); err != nil {
				_ = message.Nack(false, true)
			} else {
				_ = message.Ack(false)
			}
		}
	}
}
