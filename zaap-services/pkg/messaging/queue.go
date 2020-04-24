package messaging

import "github.com/streadway/amqp"

type (
	QueueConfiguration interface {
		Initialise(*amqp.Channel) (*amqp.Queue, error)

		Bind(*amqp.Channel, amqp.Queue, string) error

		Consume(*amqp.Channel, amqp.Queue) (<-chan amqp.Delivery, error)
	}

	simpleWorkingQueue struct {
		exchangeName string
		queueName    string
	}
)

func NewSimpleWorkingQueue(exchangeName string, queueName string) QueueConfiguration {
	return &simpleWorkingQueue{
		exchangeName: exchangeName,
		queueName:    queueName,
	}
}

func (s simpleWorkingQueue) Initialise(ch *amqp.Channel) (*amqp.Queue, error) {
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

func (s simpleWorkingQueue) Bind(ch *amqp.Channel, q amqp.Queue, routingKey string) error {
	return ch.QueueBind(q.Name, routingKey, s.exchangeName, false, nil)
}

func (s simpleWorkingQueue) Consume(ch *amqp.Channel, q amqp.Queue) (<-chan amqp.Delivery, error) {
	return ch.Consume(q.Name, "", false, false, false, false, nil)
}
