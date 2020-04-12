package daemon

import (
	"context"
	"io"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/client"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"

	"github.com/remicaumette/zaap.sh/zaap-services/pkg/backoff"
	"github.com/remicaumette/zaap.sh/zaap-services/pkg/configs"
)

type DockerEventListener struct {
	eventsQueued []events.Message
	connection   *websocket.Conn
	logger       *logrus.Entry

	factory ProxyFactoryFunc
}

func NewDockerEventListener(daemon configs.Daemon) (*DockerEventListener, error) {
	factory := ProxyFactory(daemon)

	conn, err := factory()
	if err != nil {
		return nil, err
	}

	return &DockerEventListener{
		eventsQueued: make([]events.Message, 0),
		connection:   conn,
		factory:      factory,
		logger:       logrus.WithField("listener", "docker-event-listener"),
	}, nil
}

func (d *DockerEventListener) Listen() error {

	cli, err := client.NewEnvClient()

	if err != nil {
		return err
	}

	eventMessage, eventError := cli.Events(context.Background(), types.EventsOptions{})

	for {
		select {
		case msg := <-eventMessage:
			d.eventsQueued = append(d.eventsQueued, msg)

			d.SendToDaemonProxy()
			break
		case err := <-eventError:
			if err == io.EOF {
				return nil
			}
		}
	}
}

func (d *DockerEventListener) SendToDaemonProxy() {

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

				d.SendToDaemonProxy()
				return
			}
		} else {
			d.eventsQueued = d.eventsQueued[1:]
		}
	}

}
