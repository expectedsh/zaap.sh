package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/docker/docker/api/types/events"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"

	"github.com/remicaumette/zaap.sh/zaap-services/internal/controller/react"
	"github.com/remicaumette/zaap.sh/zaap-services/pkg/models/scheduler"
	"github.com/remicaumette/zaap.sh/zaap-services/pkg/ws"
	"github.com/remicaumette/zaap.sh/zaap-services/pkg/ws/consumer"
)

type daemonWebsocketConsumer struct {
	deploymentConsumerCtx          context.Context
	deploymentConsumerCtxCanceller context.CancelFunc
	hasDeploymentConsumer          bool
	conn                           *amqp.Connection
	logger                         *logrus.Entry
}

func RegisterDaemonWebsocketConsumer(connection *amqp.Connection) http.HandlerFunc {
	ctx, cancelFunc := context.WithCancel(context.Background())
	const name = "daemon-websocket"
	factory := func() consumer.Handler {
		return &daemonWebsocketConsumer{
			deploymentConsumerCtx:          ctx,
			deploymentConsumerCtxCanceller: cancelFunc,
			hasDeploymentConsumer:          false,
			conn:                           connection,
			logger:                         logrus.WithField("consumer", name),
		}
	}

	return consumer.RegisterServerWebsocketConsumer(name, factory)
}

func (d *daemonWebsocketConsumer) Handle(message ws.Message, conn *websocket.Conn) error {
	switch message.MessageType {
	case ws.MessageTypeSchedulerToken:
		if !d.hasDeploymentConsumer {
			token := scheduler.Token{}
			if err := json.Unmarshal(message.Payload, &token); err != nil {
				return errors.Wrap(err, "unable to unmarshal in scheduler.Token")
			}

			if err := registerDeploymentConsumer(token.Token, conn, d.deploymentConsumerCtx, d.conn); err != nil {
				return errors.Wrap(err, "unable to register a deploymentConsumer")
			}

			d.logger = d.logger.WithField("associated-routine-key", token.Token)
			d.hasDeploymentConsumer = true
		}
		break
	case ws.MessageTypeDockerEvent:
		event := events.Message{}
		if err := json.Unmarshal(message.Payload, &event); err != nil {
			return errors.Wrap(err, "unable to unmarshal in  events.Message")
		}

		// since the daemon sends all docker events, we need to make sure we need to check on those received
		if react.IsReactiveTo(event) {
			react.On(event)
		}
	}

	return nil
}

func (d *daemonWebsocketConsumer) Close() error {
	if d.hasDeploymentConsumer {
		d.deploymentConsumerCtxCanceller()
	}

	return nil
}
