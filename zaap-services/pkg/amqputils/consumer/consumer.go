package consumer

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type Handler interface {

	// Handle is called when a message is ready to be handled
	//
	// If ConsumeOption.AutoAck is set to false and no error is thrown
	// the message will be ack.
	Handle(amqp.Delivery) error

	// Close is called when the consumer is done (context cancellation)
	Close() error
}

type Consumer struct {
	ctx context.Context

	ExchangeOptions  ExchangeOptions
	QueueOptions     QueueOptions
	QueueBindOptions QueueBindOptions
	ConsumeOptions   ConsumeOptions

	logger *logrus.Entry
}

type ExchangeOptions struct {
	Name, Kind                            string
	Durable, AutoDelete, Internal, NoWait bool
	Args                                  amqp.Table
}

type QueueOptions struct {
	Name                                   string
	Durable, AutoDelete, Exclusive, NoWait bool
	Args                                   amqp.Table
}

type ConsumeOptions struct {
	AutoAck, Exclusive, NoLocal, NoWait bool
	Args                                amqp.Table
}

type QueueBindOptions struct {
	RoutineKey string
	NoWait     bool
	Args       amqp.Table
}

type OptionFn func(options *Consumer)

func RegisterAmqpConsumer(handler Handler, conn *amqp.Connection, name string, options ...OptionFn) error {
	consumer := &Consumer{
		ctx:    context.Background(),
		logger: logrus.WithField("consumer-type", "rabbitmq").WithField("consumer-name", name),

		ExchangeOptions: ExchangeOptions{
			Name:       name,
			Kind:       amqp.ExchangeDirect,
			Durable:    true,
			AutoDelete: false,
			Internal:   false,
			NoWait:     false,
			Args:       nil,
		},
		QueueOptions: QueueOptions{
			Name:       "",
			Durable:    true,
			AutoDelete: false,
			Exclusive:  true,
			NoWait:     false,
			Args:       nil,
		},
		QueueBindOptions: QueueBindOptions{
			RoutineKey: "",
			NoWait:     false,
			Args:       nil,
		},
		ConsumeOptions: ConsumeOptions{
			AutoAck:   false,
			Exclusive: false,
			NoLocal:   false,
			NoWait:    false,
			Args:      nil,
		},
	}

	for _, f := range options {
		f(consumer)
	}

	if consumer.QueueBindOptions.RoutineKey != "" {
		consumer.logger = consumer.logger.WithField("routine-key", consumer.QueueBindOptions.RoutineKey)
	}

	channel, err := conn.Channel()
	if err != nil {
		return err
	}

	if err := channel.ExchangeDeclare(
		name,
		consumer.ExchangeOptions.Kind,
		consumer.ExchangeOptions.Durable,
		consumer.ExchangeOptions.AutoDelete,
		consumer.ExchangeOptions.Internal,
		consumer.ExchangeOptions.NoWait,
		consumer.ExchangeOptions.Args,
	); err != nil {
		return err
	}

	queue, err := channel.QueueDeclare(
		consumer.QueueOptions.Name,
		consumer.QueueOptions.Durable,
		consumer.QueueOptions.AutoDelete,
		consumer.QueueOptions.Exclusive,
		consumer.QueueOptions.NoWait,
		nil,
	)

	if err != nil {
		return err
	}

	if err := channel.QueueBind(
		queue.Name,
		consumer.QueueBindOptions.RoutineKey,
		consumer.ExchangeOptions.Name,
		consumer.QueueBindOptions.NoWait,
		consumer.QueueBindOptions.Args,
	); err != nil {
		return err
	}

	deliveryChan, err := channel.Consume(
		queue.Name,
		consumer.QueueBindOptions.RoutineKey,
		consumer.ConsumeOptions.AutoAck,
		consumer.ConsumeOptions.Exclusive,
		consumer.ConsumeOptions.NoLocal,
		consumer.ConsumeOptions.NoWait,
		consumer.ConsumeOptions.Args,
	)

	if err != nil {
		return err
	}

	go consumer.consume(deliveryChan, handler)

	return nil
}

func (c *Consumer) consume(deliveryChan <-chan amqp.Delivery, handler Handler) {
	c.logger.Info("handling messages")
	for {
		select {
		case <-c.ctx.Done():
			c.logger.Info("closing consumer")
			if err := handler.Close(); err != nil {
				c.logger.WithError(err).Error("unable to close handler")
			}
			return
		case msg := <-deliveryChan:

			if msg.Body == nil {
				continue
			}

			before := time.Now()
			err := handler.Handle(msg)

			// logging
			tempLogger := c.logger.WithField("elapsed", time.Now().Sub(before).String())
			if err == nil {
				tempLogger.WithField("payload", msg).Info("message handled")
			} else {
				tempLogger.WithError(err).Error("error while handling message")
			}

			// ack logic business
			if !c.ConsumeOptions.AutoAck {
				if !c.ConsumeOptions.AutoAck && err == nil {
					if err := msg.Ack(false); err != nil {
						c.logger.WithError(err).Error("unable to ack")
					}
				} else {
					if err := msg.Nack(false, true); err != nil {
						c.logger.WithError(err).Error("unable to nack")
					}
				}
			}
		}
	}
}

func WithOptionExchangeName(name string) OptionFn {
	return func(options *Consumer) {
		options.ExchangeOptions.Name = name
	}
}
func WithOptionExchangeKind(kind string) OptionFn {
	return func(options *Consumer) {
		options.ExchangeOptions.Kind = kind
	}
}
func WithOptionExchangeDurable(durable bool) OptionFn {
	return func(options *Consumer) {
		options.ExchangeOptions.Durable = durable
	}
}
func WithOptionExchangeAutoDelete(autoDelete bool) OptionFn {
	return func(options *Consumer) {
		options.ExchangeOptions.AutoDelete = autoDelete
	}
}
func WithOptionExchangeInternal(internal bool) OptionFn {
	return func(options *Consumer) {
		options.ExchangeOptions.Internal = internal
	}
}
func WithOptionExchangeNoWait(noWait bool) OptionFn {
	return func(options *Consumer) {
		options.ExchangeOptions.NoWait = noWait
	}
}
func WithOptionExchangeArgs(args amqp.Table) OptionFn {
	return func(options *Consumer) {
		options.ExchangeOptions.Args = args
	}
}

func WithOptionQueueName(name string) OptionFn {
	return func(options *Consumer) {
		options.QueueOptions.Name = name
	}
}

func WithOptionQueueDurable(durable bool) OptionFn {
	return func(options *Consumer) {
		options.QueueOptions.Durable = durable
	}
}

func WithOptionQueueAutoDelete(autoDelete bool) OptionFn {
	return func(options *Consumer) {
		options.QueueOptions.AutoDelete = autoDelete
	}
}

func WithOptionQueueExclusive(exclusive bool) OptionFn {
	return func(options *Consumer) {
		options.QueueOptions.Exclusive = exclusive
	}
}

func WithOptionQueueNoWait(noWait bool) OptionFn {
	return func(options *Consumer) {
		options.QueueOptions.NoWait = noWait
	}
}

func WithOptionQueueArgs(args amqp.Table) OptionFn {
	return func(options *Consumer) {
		options.QueueOptions.Args = args
	}
}

func WithOptionConsumeAutoAck(autoAck bool) OptionFn {
	return func(options *Consumer) {
		options.ConsumeOptions.AutoAck = autoAck
	}
}

func WithOptionConsumeExclusive(exclusive bool) OptionFn {
	return func(options *Consumer) {
		options.ConsumeOptions.Exclusive = exclusive
	}
}

func WithOptionConsumeNoLocal(noLocal bool) OptionFn {
	return func(options *Consumer) {
		options.ConsumeOptions.NoLocal = noLocal
	}
}

func WithOptionConsumeNoWait(noWait bool) OptionFn {
	return func(options *Consumer) {
		options.ConsumeOptions.NoWait = noWait
	}
}

func WithOptionConsumeArgs(args amqp.Table) OptionFn {
	return func(options *Consumer) {
		options.ConsumeOptions.Args = args
	}
}

func WithOptionQueueBindRoutineKey(routineKey string) OptionFn {
	return func(options *Consumer) {
		options.QueueBindOptions.RoutineKey = routineKey
	}
}

func WithOptionQueueBindNoWait(noWait bool) OptionFn {
	return func(options *Consumer) {
		options.QueueBindOptions.NoWait = noWait
	}
}

func WithOptionQueueBindArgs(args amqp.Table) OptionFn {
	return func(options *Consumer) {
		options.QueueBindOptions.Args = args
	}
}

func WithOptionContext(ctx context.Context) OptionFn {
	return func(options *Consumer) {
		options.ctx = ctx
	}
}
