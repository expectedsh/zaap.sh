package core

import (
  "github.com/gorilla/websocket"
  "github.com/sirupsen/logrus"
  "net/http"
)

type WebsocketServer struct {
  upgrader websocket.Upgrader
  log      *logrus.Entry
}

func NewWebsocketServer() *WebsocketServer {
  return &WebsocketServer{
    upgrader: websocket.Upgrader{},
    log:      logrus.WithField("app", "websocket"),
  }
}

func (s *WebsocketServer) Handler(w http.ResponseWriter, r *http.Request) {
  conn, err := s.upgrader.Upgrade(w, r, nil)
  if err != nil {
    s.log.WithError(err).Error("failed to upgrade connection")
    return
  }
  defer conn.Close()
  for {
    mtype, message, err := conn.ReadMessage()
    if err != nil {
      s.log.WithError(err).Error("failed to read message")
      break
    }
    s.log.WithField("type", mtype).Info(message)
  }
}

func (s *WebsocketServer) Start() {
  httpServer := http.Server{
    Addr:    ":8090",
    Handler: http.HandlerFunc(s.Handler),
  }
  s.log.Infof("listening on %v", httpServer.Addr)
  if err := httpServer.ListenAndServe(); err != nil {
    panic(err)
  }
}
