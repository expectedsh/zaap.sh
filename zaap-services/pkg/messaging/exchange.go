package messaging

import (
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/connector/rabbitmq"
	"github.com/streadway/amqp"
)

type (
	ExchangeConfig interface {
		Name() string

		Declare(*rabbitmq.Channel) error
	}

	durableExchangeTopic struct {
		exchangeName string
	}
)

func NewDurableExchangeTopic(exchangeName string) ExchangeConfig {
	return &durableExchangeTopic{
		exchangeName: exchangeName,
	}
}

func (d durableExchangeTopic) Name() string {
	return d.exchangeName
}

func (d durableExchangeTopic) Declare(ch *rabbitmq.Channel) error {
	return ch.ExchangeDeclare(d.exchangeName, amqp.ExchangeTopic, true, false, false, false, nil)
}
