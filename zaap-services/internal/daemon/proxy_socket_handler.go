package daemon

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"

	"github.com/remicaumette/zaap.sh/zaap-services/pkg/ws"
)

func ProxySocketHandler(connection *websocket.Conn) {

	logger := logrus.WithField("socket-handler", "proxy")

	go func() {
		for {
			_, message, err := connection.ReadMessage()

			if err != nil {
				if ws.IsClosedError(err) {
					break
				}
				logger.WithError(err).Warn()
			}

			logger.WithField("message", string(message))
		}
	}()
}
