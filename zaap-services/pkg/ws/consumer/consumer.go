package consumer

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"

	"github.com/remicaumette/zaap.sh/zaap-services/pkg/ws"
)

type Handler interface {

	// Handle is called when a message is ready to be handled
	Handle(message ws.Message, conn *websocket.Conn) error

	// Close is called when the consumer is done (context cancellation)
	Close() error
}

func RegisterHttpWebsocketConsumer(name string, handlerFactory func() Handler) http.HandlerFunc {
	var daemonSocketHandler = websocket.Upgrader{}
	logger := logrus.WithField("consumer-type", "websocket").WithField("consumer-name", name)

	return func(w http.ResponseWriter, r *http.Request) {
		connection, err := daemonSocketHandler.Upgrade(w, r, nil)
		if err != nil {
			logger.WithError(err).Error()
			return
		}

		defer connection.Close()

		handler := handlerFactory()

		logger.Info("handling messages")
		for {
			_, message, err := connection.ReadMessage()

			if err != nil {
				if ws.IsClosedError(err) {
					logger.WithError(err).Info("closing consumer")
					if err := handler.Close(); err != nil {
						logger.WithError(err).Error("unable to close handler")
					}
					return
				}
				logger.WithError(err).Warn("an error occurred while reading a message")
				continue
			}

			wsMessage, err := ws.NewMessageFromBytes(message)
			if err != nil {
				logger.WithField("message", string(message)).WithError(err).
					Warn("unable to unmarshal message")
			}

			before := time.Now()
			if err := handler.Handle(*wsMessage, connection); err != nil {
				logger.WithField("elapsed", time.Now().Sub(before).String()).
					WithError(err).Error("handler return an error")
				return
			} else {
				logger.WithField("elapsed", time.Now().Sub(before).String()).
					Info("message handled")
			}
		}
	}
}
