package controller

import (
	"encoding/json"
	"net/http"

	"github.com/docker/docker/api/types/events"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"

	"github.com/remicaumette/zaap.sh/zaap-services/internal/controller/react"
	"github.com/remicaumette/zaap.sh/zaap-services/pkg/ws"
)

var daemonSocketHandler = websocket.Upgrader{}

func DaemonSocketHandler(w http.ResponseWriter, r *http.Request) {
	logger := logrus.WithField("socket-handler", "daemon")

	connection, err := daemonSocketHandler.Upgrade(w, r, nil)
	if err != nil {
		logger.WithError(err).Error()
		return
	}

	defer connection.Close()

	for {
		_, message, err := connection.ReadMessage()

		if err != nil {
			if ws.IsClosedError(err) {
				// in this case the connection is closed and the websocket client is not reachable
				break
			}
			logger.WithError(err).Warn("an error occured while reading a message")
			continue
		}

		event := events.Message{}
		if err := json.Unmarshal(message, &event); err != nil {
			logger.WithError(err).Warn("a message could not be unmarshal to docker event type")
			continue
		}

		// since the daemon send all docker events we must ensure that the event need is reactable
		if react.IsReactiveTo(event) {
			react.On(event)
		}
	}
}
