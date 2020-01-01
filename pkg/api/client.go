package api

import (
  "github.com/remicaumette/zaap.sh/pkg/protocol"
  "google.golang.org/grpc"
)

type Client struct {
  Conn       *grpc.ClientConn
  AppService protocol.AppServiceClient
}

func NewClient(addr string) (*Client, error) {
  conn, err := grpc.Dial(addr, grpc.WithInsecure())
  if err != nil {
    return nil, err
  }
  return &Client{
    Conn:       conn,
    AppService: protocol.NewAppServiceClient(conn),
  }, nil
}
