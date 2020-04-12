package daemon

import (
	"net/url"

	"github.com/gorilla/websocket"

	"github.com/remicaumette/zaap.sh/pkg/configs"
)

type ProxyFactoryFunc func() (*websocket.Conn, error)

func ProxyFactory(daemonConfig configs.Daemon) ProxyFactoryFunc {
	daemonProxyUrl := url.URL{Scheme: "ws", Host: daemonConfig.DaemonProxyAddress, Path: "/daemon"}

	return func() (*websocket.Conn, error) {
		connection, _, err := websocket.DefaultDialer.Dial(daemonProxyUrl.String(), nil)
		if err != nil {
			return nil, err
		}
		return connection, nil
	}
}
