package client

import (
  "github.com/remicaumette/zaap.sh/protocol"
  "google.golang.org/grpc"
)

type Client struct {
  Conn       *grpc.ClientConn
  AppService protocol.AppServiceClient
}

func New(addr string) (*Client, error) {
  conn, err := grpc.Dial(addr, grpc.WithInsecure())
  if err != nil {
    return nil, err
  }
  return &Client{
    Conn:       conn,
    AppService: protocol.NewAppServiceClient(conn),
  }, nil
}
