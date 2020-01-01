package server

import (
  "github.com/remicaumette/zaap.sh/pkg/protocol"
  "github.com/remicaumette/zaap.sh/pkg/server/app"
  "github.com/sirupsen/logrus"
  "google.golang.org/grpc"
  "net"
)

type Server struct {
  AppService *app.Service
}

func New(appService *app.Service) *Server {
  return &Server{
    AppService: appService,
  }
}

func (s *Server) Start(addr string) error {
  lis, err := net.Listen("tcp", addr)
  if err != nil {
    return err
  }
  srv := grpc.NewServer()
  protocol.RegisterAppServiceServer(srv, s.AppService)
  logrus.Infof("listening on %v", addr)
  return srv.Serve(lis)
}
