package ws

import "github.com/gorilla/websocket"

func IsClosedError(err error) bool {
	code := err.(*websocket.CloseError).Code
	return code >= 1000 && code <= 1015
}
