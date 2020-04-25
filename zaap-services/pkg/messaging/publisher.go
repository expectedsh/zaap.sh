package messaging

import (
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
	"reflect"
)

type (
	Publisher struct {
		conn           *amqp.Connection
		exchangeConfig ExchangeConfig
	}

	DeliveryMode uint8
)

const (
	DeliveryModeTransient  DeliveryMode = 1
	DeliveryModePersistent              = 2
)

func NewPublisher(conn *amqp.Connection, exchangeConfig ExchangeConfig) *Publisher {
	return &Publisher{
		conn:           conn,
		exchangeConfig: exchangeConfig,
	}
}

func (s Publisher) Publish(deliveryMode DeliveryMode, message proto.Message) error {
	ch, err := s.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = s.exchangeConfig.Declare(ch)
	if err != nil {
		return err
	}

	messageType := reflect.TypeOf(message).Elem().Name()

	body, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	return ch.Publish(s.exchangeConfig.Name(), messageType, false, false, amqp.Publishing{
		DeliveryMode: uint8(deliveryMode),
		Body:         body,
	})
}
