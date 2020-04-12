package proxy

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types/events"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"

	"github.com/remicaumette/zaap.sh/pkg/ws"
)

var daemonSocketHandler = websocket.Upgrader{}

func DaemonSocketHandler(queue *DockerEventQueue) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

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

			// since the daemon send all docker event we ensure that the message type + action is one of desired
			// to be retrieve from the api
			if canBeSent(event, logger) {
				if err := queue.SendDockerEvent(message); err != nil {
					logger.WithError(err).Warn("a message could not be send to docker event queue")
					continue
				}
			}
		}
	}
}

func canBeSent(message events.Message, logger *logrus.Entry) bool {
	canBeSentLogger := logger.
		WithField("filtering-message", fmt.Sprintf("type: %s, action: %s", message.Type, message.Action))

	if message.Type == "container" {
		switch message.Action {
		case "create", "delete", "die", "stop", "start":
			canBeSentLogger.WithField("accepted", true).Info()
			return true
		}
	}

	canBeSentLogger.WithField("accepted", false).Info()
	return false
}
