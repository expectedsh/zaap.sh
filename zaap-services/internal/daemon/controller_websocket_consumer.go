package daemon

import (
	"context"
	"encoding/json"
	"net/url"
	"time"

	"github.com/docker/docker/client"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"

	"github.com/remicaumette/zaap.sh/zaap-services/pkg/core"
	"github.com/remicaumette/zaap.sh/zaap-services/pkg/utils/ws"
	"github.com/remicaumette/zaap.sh/zaap-services/pkg/utils/ws/consumer"
)

type controllerConsumer struct {
	daemon *daemon

	hasSentSchedulerToken bool
}

func RegisterControllerConsumer(ctx context.Context, schedulerToken string, controllerWsUrl url.URL, client *client.Client) error {
	factory := func() consumer.Handler {
		return &controllerConsumer{daemon: newDaemon(client, schedulerToken)}
	}

	err := consumer.RegisterClientWebsocketConsumer(ctx, "controller", 1*time.Second, controllerWsUrl, factory)
	if err != nil {
		return err
	}

	return nil
}

func (c *controllerConsumer) Handle(message ws.Message, connection *websocket.Conn) error {
	switch message.MessageType {
	case ws.MessageTypeApplicationDeployment:
		payload := core.DeploymentPayload{}
		if err := json.Unmarshal(message.Payload, &payload); err != nil {
			return errors.Wrap(err, "unable to unmarshal in application.DeploymentPayload")
		}
		if err := c.daemon.deployApplication(payload); err != nil {
			return err
		}
	}

	return nil
}

func (c *controllerConsumer) OnConnectionCreation(conn *websocket.Conn) error {
	if !c.hasSentSchedulerToken {
		message, err := ws.NewMessage(ws.MessageTypeSchedulerToken, core.Token{Token: c.daemon.schedulerToken})
		if err != nil {
			return err
		}

		if err := conn.WriteJSON(message); err != nil {
			return err
		} else {
			c.hasSentSchedulerToken = true
		}
	}

	return nil
}

func (c *controllerConsumer) Close() error {
	return nil
}
