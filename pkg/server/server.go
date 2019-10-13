package server

import (
  "context"
  "net/http"
)

type Server struct {
  httpServer *http.Server
}

func New(config *Config) *Server {
  return &Server{
    httpServer: &http.Server{

    },
  }
}

func (s *Server) Start() error {
  return nil
}

func (s *Server) Stop(ctx context.Context) error {
  return s.httpServer.Shutdown(ctx)
}
