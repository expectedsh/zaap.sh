package daemon

import (
	"net/url"

	"github.com/gorilla/websocket"

	"github.com/remicaumette/zaap.sh/zaap-services/pkg/configs"
)

type controllerFactoryFunc func() (*websocket.Conn, error)

func controllerWebsocketFactory(daemonConfig configs.Daemon) controllerFactoryFunc {
	daemonProxyUrl := url.URL{Scheme: "ws", Host: daemonConfig.DaemonProxyAddress, Path: "/"}

	return func() (*websocket.Conn, error) {
		connection, _, err := websocket.DefaultDialer.Dial(daemonProxyUrl.String(), nil)
		if err != nil {
			return nil, err
		}

		return connection, nil
	}
}
