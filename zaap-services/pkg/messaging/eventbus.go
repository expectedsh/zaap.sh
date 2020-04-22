package messaging

import (
	"context"
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"reflect"
)

var ErrUnknownMessageType = errors.New("unknown message type")

type EventBus struct {
	exchangeName    string
	conn            *amqp.Connection
	messageRegistry map[string]reflect.Type
}

type EventListener struct {
	Messages <-chan proto.Message
	Errors   <-chan error
}

func NewEventBus(exchangeName string, conn *amqp.Connection) *EventBus {
	return &EventBus{
		exchangeName:    exchangeName,
		conn:            conn,
		messageRegistry: make(map[string]reflect.Type),
	}
}

func (b *EventBus) RegisterMessage(name string, msg proto.Message) {
	b.messageRegistry[name] = reflect.TypeOf(msg).Elem()
}

func (b *EventBus) Emit(msg proto.Message) error {
	messageType := reflect.TypeOf(msg).Elem().Name()
	if b.messageRegistry[messageType] == nil {
		return ErrUnknownMessageType
	}

	ch, err := b.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(b.exchangeName, amqp.ExchangeFanout, true, false, false, false, nil)
	if err != nil {
		return err
	}

	payload, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	return ch.Publish(b.exchangeName, "", false, false, amqp.Publishing{
		Headers: amqp.Table{
			"X-Message-Type": messageType,
		},
		Body: payload,
	})
}

func (b *EventBus) Listen(ctx context.Context) (*EventListener, error) {
	initialized := false

	ch, err := b.conn.Channel()
	if err != nil {
		return nil, err
	}
	defer func() {
		if initialized == false {
			_ = ch.Close()
		}
	}()

	err = ch.ExchangeDeclare(b.exchangeName, amqp.ExchangeFanout, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		return nil, err
	}

	err = ch.QueueBind(q.Name, "", b.exchangeName, false, nil)
	if err != nil {
		return nil, err
	}

	messages, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	events := make(chan proto.Message)
	errors := make(chan error)
	initialized = true

	go func() {
		defer func() {
			close(events)
			_ = ch.Close()
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-messages:
				messageTypeName := msg.Headers["X-Message-Type"].(string)
				messageType := b.messageRegistry[messageTypeName]
				if messageType == nil {
					errors <- ErrUnknownMessageType
					continue
				}

				payload := reflect.New(messageType).Interface().(proto.Message)
				if err := proto.Unmarshal(msg.Body, payload); err != nil {
					logrus.WithError(err).Warn("failed to unmarshal payload")
					errors <- err
					continue
				}

				events <- payload
			}
		}
	}()

	return &EventListener{
		Messages: events,
		Errors:   errors,
	}, nil
}
