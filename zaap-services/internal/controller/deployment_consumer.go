package controller

import (
	"context"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"

	"github.com/remicaumette/zaap.sh/zaap-services/pkg/amqputils/consumer"
	"github.com/remicaumette/zaap.sh/zaap-services/pkg/ws"
)

type deploymentConsumer struct {
	websocket *websocket.Conn
}

func registerDeploymentConsumer(
	schedulerToken string,
	websocket *websocket.Conn,
	ctx context.Context,
	connection *amqp.Connection) error {

	deploymentConsumer := deploymentConsumer{websocket: websocket}

	options := []consumer.OptionFn{
		consumer.WithOptionQueueBindRoutineKey(fmt.Sprintf("deployment-consumer-%s", schedulerToken)),
		consumer.WithOptionContext(ctx),
	}
	if err := consumer.RegisterAmqpConsumer(&deploymentConsumer, connection, "deployment", options...); err != nil {
		return err
	}

	return nil
}

func (d *deploymentConsumer) Handle(delivery amqp.Delivery) error {
	fmt.Println(string(delivery.Body))
	wsMessage, err := ws.NewMessageRaw(ws.MessageTypeDeployment, delivery.Body)
	if err != nil {
		return errors.Wrap(err, "unable to create ws.Message from  amqp.Delivery")
	}

	if err := d.websocket.WriteJSON(wsMessage); err != nil {
		return err
	}

	return nil
}

func (d *deploymentConsumer) Close() error {
	return nil
}
