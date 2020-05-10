package messaging

import (
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/connector/rabbitmq"
	"github.com/streadway/amqp"
)

type (
	QueueConfig interface {
		Name() string

		Declare(*rabbitmq.Channel) (*amqp.Queue, error)

		Bind(*rabbitmq.Channel, amqp.Queue, string) error

		Consume(*rabbitmq.Channel, amqp.Queue) (<-chan amqp.Delivery, error)
	}

	simpleWorkingQueue struct {
		exchangeName string
		queueName    string
	}
)

func NewSimpleWorkingQueue(exchangeName string, queueName string) QueueConfig {
	return &simpleWorkingQueue{
		exchangeName: exchangeName,
		queueName:    queueName,
	}
}

func (s simpleWorkingQueue) Name() string {
	return s.queueName
}

func (s simpleWorkingQueue) Declare(ch *rabbitmq.Channel) (*amqp.Queue, error) {
	err := ch.ExchangeDeclare(s.exchangeName, amqp.ExchangeTopic, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(s.queueName, false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	err = ch.Qos(1, 0, false)
	if err != nil {
		return nil, err
	}

	return &q, nil
}

func (s simpleWorkingQueue) Bind(ch *rabbitmq.Channel, q amqp.Queue, routingKey string) error {
	return ch.QueueBind(q.Name, routingKey, s.exchangeName, false, nil)
}

func (s simpleWorkingQueue) Consume(ch *rabbitmq.Channel, q amqp.Queue) (<-chan amqp.Delivery, error) {
	return ch.Consume(q.Name, "", false, false, false, false, nil)
}
