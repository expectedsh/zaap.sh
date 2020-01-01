package app

import (
  "context"
  "github.com/remicaumette/zaap.sh/pkg/protocol"
)

func (s *Service) GetApp(ctx context.Context, r *protocol.GetAppRequest) (*protocol.App, error) {
  return &protocol.App{}, nil
}
