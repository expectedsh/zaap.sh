package messaging

import "github.com/streadway/amqp"

type (
	ExchangeConfig interface {
		Name() string

		Declare(*amqp.Channel) error
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

func (d durableExchangeTopic) Declare(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(d.exchangeName, amqp.ExchangeTopic, true, false, false, false, nil)
}
