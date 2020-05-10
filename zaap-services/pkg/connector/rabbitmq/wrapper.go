package rabbitmq

import (
	"github.com/expected.sh/zaap.sh/zaap-services/pkg/backoff"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"sync/atomic"
)

type Connection struct {
	*amqp.Connection
}

type Channel struct {
	*amqp.Channel
	closed int32
}

func (c *Connection) Channel() (*Channel, error) {
	ch, err := c.Connection.Channel()
	if err != nil {
		return nil, err
	}

	channel := &Channel{Channel: ch}

	go func() {
		errors := ch.NotifyClose(make(chan *amqp.Error))
		for {
			amqpErr, ok := <-errors
			if !ok || channel.IsClosed() {
				_ = channel.Close()
				break
			}

			err = backoff.New("rabbitmq channel closed, reopening...", func() error {
				ch, err = c.Connection.Channel()
				if err != nil {
					return err
				}
				return nil
			}, logrus.WithError(amqpErr)).Run()
			if err != nil {
				logrus.WithError(err).Fatal("could not reopen channel")
				return
			}

			channel.Channel = ch
		}
	}()

	return channel, nil
}

func (ch *Channel) IsClosed() bool {
	return atomic.LoadInt32(&ch.closed) == 1
}

func (ch *Channel) Close() error {
	if ch.IsClosed() {
		return amqp.ErrClosed
	}

	atomic.StoreInt32(&ch.closed, 1)
	return ch.Channel.Close()
}

func (ch *Channel) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	deliveries := make(chan amqp.Delivery)

	go func() {
		for !ch.IsClosed() {
			var messages <-chan amqp.Delivery

			err := backoff.New("consuming message", func() error {
				d, err := ch.Channel.Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
				if err != nil {
					return err
				}
				messages = d
				return nil
			}, nil).Run()
			if err != nil {
				logrus.WithError(err).Fatal("could not consume message")
				return
			}

			for msg := range messages {
				deliveries <- msg
			}
		}
		close(deliveries)
	}()

	return deliveries, nil
}
