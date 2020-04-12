package proxy

import (
	"github.com/streadway/amqp"
)

const (
	dockerEventExchangeName = "docker-event"
	dockerEventKey          = "docker-event"
	dockerEventExchangeType = amqp.ExchangeDirect
	dockerEventDeliveryMode = amqp.Persistent
)

type DockerEventQueue struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewDockerEventQueue(conn *amqp.Connection) (*DockerEventQueue, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	if err := channel.ExchangeDeclare(
		dockerEventExchangeName,
		dockerEventExchangeType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return nil, err
	}
	return &DockerEventQueue{conn: conn, channel: channel}, nil
}

func (d DockerEventQueue) SendDockerEvent(bytes []byte) error {
	if err := d.channel.Publish(dockerEventExchangeName, dockerEventKey, false, false, amqp.Publishing{
		Headers:         amqp.Table{},
		ContentType:     "application/json",
		ContentEncoding: "",
		DeliveryMode:    dockerEventDeliveryMode,
		Body:            bytes,
	}); err != nil {
		return err
	}

	return nil
}
