package daemon

import (
	"context"
	"io"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"

	"github.com/remicaumette/zaap.sh/zaap-services/pkg/backoff"
	"github.com/remicaumette/zaap.sh/zaap-services/pkg/configs"
	"github.com/remicaumette/zaap.sh/zaap-services/pkg/ws"
)

type DockerEventConsumer struct {
	eventsQueued []ws.Message
	connection   *websocket.Conn
	logger       *logrus.Entry

	ctx          context.Context
	dockerClient *client.Client
	factory      controllerFactoryFunc
}

func NewDockerEventConsumer(ctx context.Context, daemon configs.Daemon, dockerClient *client.Client) (*DockerEventConsumer, error) {
	factory := controllerWebsocketFactory(daemon)

	conn, err := factory()
	if err != nil {
		return nil, err
	}

	return &DockerEventConsumer{
		ctx:          ctx,
		dockerClient: dockerClient,
		eventsQueued: make([]ws.Message, 0),
		connection:   conn,
		factory:      factory,
		logger: logrus.WithField("consumer-type", "docker-event").
			WithField("consumer-name", "docker-event-listener"),
	}, nil
}

func (d *DockerEventConsumer) Listen() error {
	eventMessage, eventError := d.dockerClient.Events(context.Background(), types.EventsOptions{})

	for {
		select {
		case <-d.ctx.Done():
			d.logger.Info("closing consumer due to app termination")
			return nil
		case msg := <-eventMessage:
			message, err := ws.NewMessage(ws.MessageTypeDockerEvent, msg)
			if err != nil {
				d.logger.WithError(err).Warn("unable to create ws.Message")
				continue
			}

			d.eventsQueued = append(d.eventsQueued, *message)
			d.sendToDaemonProxy()

			break
		case err := <-eventError:
			if err == io.EOF {
				return nil
			}
		}
	}
}

func (d *DockerEventConsumer) sendToDaemonProxy() {

	for _, e := range d.eventsQueued {
		if err := d.connection.WriteJSON(e); err != nil {

			err := backoff.New("try connecting to daemon controller", func() error {
				conn, err := d.factory()
				if err != nil {
					return err
				}

				d.connection = conn
				return nil
			}, d.logger).
				WithInterval(time.Millisecond * 100).
				WithMaxAttempt(10).
				Run()

			if err != nil {

				// In this case 10 attempt has been tried to connect to the daemon controller.
				// Since the server is not available we are relying to the next docker event
				// to retry the connection.

				return
			} else {

				// In this case the connection is established so we can retry the WriteJSON operation
				// previously failed.

				d.sendToDaemonProxy()
				return
			}
		} else {
			d.eventsQueued = d.eventsQueued[1:]
		}
	}

}
