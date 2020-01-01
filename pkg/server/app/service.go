package app

import (
  "github.com/docker/docker/client"
)

type Service struct {
  Docker *client.Client
}

func NewService(docker *client.Client) *Service {
  return &Service{
    Docker: docker,
  }
}
