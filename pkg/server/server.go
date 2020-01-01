package server

import (
  "github.com/remicaumette/zaap.sh/pkg/protocol"
  "github.com/remicaumette/zaap.sh/pkg/server/app"
  "google.golang.org/grpc"
  "net"
)

type Server struct {
  AppService *app.Service
}

func New() *Server {
  return &Server{
    AppService: &app.Service{},
  }
}

func (s *Server) Start() error {
  lis, err := net.Listen("tcp", ":5200")
  if err != nil {
    return err
  }
  srv := grpc.NewServer()
  protocol.RegisterAppServiceServer(srv, s.AppService)
  return srv.Serve(lis)
}
