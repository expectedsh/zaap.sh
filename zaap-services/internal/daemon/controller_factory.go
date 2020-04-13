package daemon

import (
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"

	"github.com/remicaumette/zaap.sh/zaap-services/pkg/configs"
	"github.com/remicaumette/zaap.sh/zaap-services/pkg/models/scheduler"
	"github.com/remicaumette/zaap.sh/zaap-services/pkg/ws"
)

type controllerFactoryFunc func() (*websocket.Conn, error)

func controllerFactory(daemonConfig configs.Daemon) controllerFactoryFunc {
	daemonProxyUrl := url.URL{Scheme: "ws", Host: daemonConfig.DaemonProxyAddress, Path: "/"}

	return func() (*websocket.Conn, error) {
		connection, _, err := websocket.DefaultDialer.Dial(daemonProxyUrl.String(), nil)
		if err != nil {
			return nil, err
		}

		message, err := ws.NewMessage(ws.MessageTypeSchedulerToken, scheduler.Token{Token: daemonConfig.SchedulerToken})
		if err != nil {
			return nil, errors.Wrap(err, "sending SchedulerToken: unable to marshal a new ws.Message")
		}

		if err := connection.WriteJSON(message); err != nil {
			return nil, errors.Wrap(err, "sending SchedulerToken: unable to writeJSON")
		}

		return connection, nil
	}
}
