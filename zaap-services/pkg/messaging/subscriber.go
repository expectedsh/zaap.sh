package messaging

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
	"reflect"
)

type Subscriber struct {
	conn            *amqp.Connection
	exchangeConfig  ExchangeConfig
	queueConfig     QueueConfig
	messageTypes    map[string]reflect.Type
	messageHandlers map[string]interface{}
	ErrorHandler    func(error, interface{}) bool
}

var (
	contextType = reflect.TypeOf(context.TODO())
	errorType   = reflect.TypeOf((*error)(nil)).Elem()
)

func NewSubscriber(conn *amqp.Connection, exchangeConfig ExchangeConfig, queueConfig QueueConfig) *Subscriber {
	return &Subscriber{
		conn:            conn,
		exchangeConfig:  exchangeConfig,
		queueConfig:     queueConfig,
		messageTypes:    make(map[string]reflect.Type),
		messageHandlers: make(map[string]interface{}),
	}
}

func (s *Subscriber) RegisterHandler(v interface{}) {
	function := reflect.TypeOf(v)
	if function.Kind() != reflect.Func {
		panic("subscriber handler must be a function")
	}
	if function.NumIn() != 2 {
		panic("subscriber handler function must take 2 arguments")
	}
	if function.In(0).Kind() == contextType.Kind() {
		panic("subscriber handler function first argument must be the context")
	}
	if function.NumOut() != 1 || !function.Out(0).Implements(errorType) {
		panic("subscriber handler function must return an error type")
	}
	messageType := function.In(1).Elem()
	s.messageTypes[messageType.Name()] = messageType
	s.messageHandlers[messageType.Name()] = v
}

func (s *Subscriber) Subscribe(ctx context.Context) error {
	ch, err := s.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = s.exchangeConfig.Declare(ch)
	if err != nil {
		return err
	}

	queue, err := s.queueConfig.Declare(ch)
	if err != nil {
		return err
	}

	for messageType, _ := range s.messageHandlers {
		err = s.queueConfig.Bind(ch, *queue, messageType)
		if err != nil {
			return err
		}
	}

	messages, err := s.queueConfig.Consume(ch, *queue)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case message := <-messages:
			messageType := s.messageTypes[message.RoutingKey]
			if messageType == nil {
				_ = message.Nack(false, true)
				continue
			}

			payload := reflect.New(messageType).Interface().(proto.Message)
			if err = proto.Unmarshal(message.Body, payload); err != nil {
				_ = message.Nack(false, true)
				continue
			}

			handler := s.messageHandlers[message.RoutingKey]
			if handler == nil {
				_ = message.Nack(false, true)
				continue
			}

			values := reflect.ValueOf(handler).Call([]reflect.Value{
				reflect.ValueOf(ctx),
				reflect.ValueOf(payload),
			})

			if values[0].Interface() == nil {
				_ = message.Ack(false)
			} else {
				requeue := true
				if s.ErrorHandler != nil {
					requeue = s.ErrorHandler(err, payload)
				}
				_ = message.Nack(false, requeue)
			}
		}
	}
}
