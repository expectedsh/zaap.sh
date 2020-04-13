package controller

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"

	"github.com/remicaumette/zaap.sh/zaap-services/pkg/utils/amqputils/consumer"
	"github.com/remicaumette/zaap.sh/zaap-services/pkg/utils/ws"
)

type deploymentConsumer struct {
	websocket *websocket.Conn
}

func (s *Server) registerDeploymentConsumer(
	schedulerToken string,
	websocket *websocket.Conn,
	ctx context.Context) error {

	deploymentConsumer := deploymentConsumer{websocket: websocket}

	options := []consumer.OptionFn{
		consumer.WithOptionQueueBindRoutineKey(schedulerToken),
		consumer.WithOptionContext(ctx),
	}
	if err := consumer.RegisterAmqpConsumer(&deploymentConsumer, s.amqpConnection, "deployments", options...); err != nil {
		return err
	}

	return nil
}

func (d *deploymentConsumer) Handle(delivery amqp.Delivery) error {
	wsMessage, err := ws.NewMessageRaw(ws.MessageTypeApplicationDeployment, delivery.Body)
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
