package proxy

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var apiSocketHandler = websocket.Upgrader{}

func APISocketHandler(w http.ResponseWriter, r *http.Request) {

	logger := logrus.WithField("socket-handler", "api")

	connection, err := apiSocketHandler.Upgrade(w, r, nil)
	if err != nil {
		logger.WithError(err).Error()
		return
	}

	defer connection.Close()

	for {
		_, message, err := connection.ReadMessage()

		if err != nil {
			logger.WithError(err).Warn("An error ")
			continue
		}

		logger.WithField("message", string(message)).Info()
	}
}
