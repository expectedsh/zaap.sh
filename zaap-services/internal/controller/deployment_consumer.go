package controller

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type DeploymentConsumer struct {
	ctx        context.Context
	connection *amqp.Connection
	channel    *amqp.Channel
	deliveries <-chan amqp.Delivery
	logger     *logrus.Entry
}

func NewDeploymentQueueHandler(
	ctx context.Context,
	connection *amqp.Connection,
	schedulerToken string) (*DeploymentConsumer, error) {

	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	if err := channel.ExchangeDeclare(
		"deployment",
		amqp.ExchangeDirect,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return nil, err
	}

	queue, err := channel.QueueDeclare(
		"deployment",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	if err := channel.QueueBind(
		queue.Name,
		schedulerToken,
		"deployment",
		false,
		nil,
	); err != nil {
		return nil, err
	}

	deliveries, err := channel.Consume(
		queue.Name,
		"deployment-consumer-"+schedulerToken,
		false,
		false,
		false,
		false,
		nil,
	)

	return &DeploymentConsumer{
		connection: connection,
		channel:    channel,
		deliveries: deliveries,
		ctx:        ctx,
		logger:     logrus.WithField("amqp-consumer", "deployment-consumer-"+schedulerToken),
	}, nil
}

func (d *DeploymentConsumer) Consume() {
	d.logger.Info("consuming")
	for {
		cancelled := false
		select {
		case <-d.ctx.Done():
			d.logger.Info("terminated")
			if err := d.channel.Close(); err != nil {
				d.logger.Info("unable to close correctly")
			}
			cancelled = true
			break
		case msg := <-d.deliveries:
			d.logger.WithField("payload", string(msg.Body)).Info("consume message")
		}

		if cancelled {
			break
		}
	}
}
