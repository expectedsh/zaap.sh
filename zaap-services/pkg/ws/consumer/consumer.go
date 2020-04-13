package consumer

import (
	"context"
	"net/http"
	"net/url"
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

func RegisterServerWebsocketConsumer(name string, handlerFactory func() Handler) http.HandlerFunc {
	var daemonSocketHandler = websocket.Upgrader{}
	logger := logrus.WithField("consumer-type", "server-websocket").WithField("consumer-name", name)

	return func(w http.ResponseWriter, r *http.Request) {
		connection, err := daemonSocketHandler.Upgrade(w, r, nil)
		if err != nil {
			logger.WithError(err).Error()
			return
		}

		err = handleConnection(r.Context(), handlerFactory, logger, connection)
	}
}

func RegisterClientWebsocketConsumer(
	ctx context.Context,
	name string,
	retriesInterval time.Duration,
	connectionUrl url.URL,
	handlerFactory func() Handler) error {

	logger := logrus.WithField("consumer-type", "client-websocket").WithField("consumer-name", name)

	connection, _, err := websocket.DefaultDialer.DialContext(ctx, connectionUrl.String(), nil)
	if err != nil {
		return err
	}

	for {
		if err := handleConnection(ctx, handlerFactory, logger, connection); err == context.Canceled {
			return nil
		}

		// reconnection logic
		for {
			logger.Info("trying to connect to ", connectionUrl.String())

			connection, _, err = websocket.DefaultDialer.Dial(connectionUrl.String(), nil)
			if err != nil {
				time.Sleep(retriesInterval)
				continue
			}

			break
		}
	}
}

func handleConnection(
	ctx context.Context,
	handlerFactory func() Handler,
	logger *logrus.Entry,
	connection *websocket.Conn) error {

	handler := handlerFactory()

	logger.Info("handling messages")

	messages, errors := readMessageAsChannel(connection, ctx)

	defer func() {
		if err := connection.Close(); err != nil {
			logger.WithError(err).Error("unable to close connection")
		}
	}()

	for {
		select {
		case <-ctx.Done():
			logger.Info("closing consumer due to app termination")
			if err := handler.Close(); err != nil {
				logger.WithError(err).Error("unable to close handler")
			}

			return ctx.Err()
		case err := <-errors:
			if ws.IsClosedError(err) {
				logger.WithError(err).Info("closing consumer due to error")
				if err := handler.Close(); err != nil {
					logger.WithError(err).Error("unable to close handler")
				}

				return err
			}
			logger.WithError(err).Warn("an error occurred while reading a message")
			continue
		case message := <-messages:

			wsMessage, err := ws.NewMessageFromBytes(message)
			if err != nil {
				logger.WithField("message", string(message)).WithError(err).
					Warn("unable to unmarshal message")
			}

			before := time.Now()
			if err := handler.Handle(*wsMessage, connection); err != nil {
				logger.WithField("elapsed", time.Now().Sub(before).String()).
					WithError(err).Error("handler return an error")
				continue
			} else {
				logger.WithField("elapsed", time.Now().Sub(before).String()).
					Info("message handled")
			}
		}
	}
}

func readMessageAsChannel(connection *websocket.Conn, ctx context.Context) (<-chan []byte, <-chan error) {
	msgCh, errCh := make(chan []byte), make(chan error)
	go func() {
		for {
			_, message, err := connection.ReadMessage()
			if err != nil {
				errCh <- err
			} else {
				msgCh <- message
			}

			if err := ctx.Err(); err != nil {
				return
			}
		}
	}()

	return msgCh, errCh
}
