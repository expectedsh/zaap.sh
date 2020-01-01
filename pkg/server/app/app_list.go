package app

import (
  "context"
  "github.com/remicaumette/zaap.sh/pkg/protocol"
)

func (s *Service) ListApps(ctx context.Context, r *protocol.ListAppsRequest) (*protocol.ListAppsResponse, error) {
  return &protocol.ListAppsResponse{
    Apps: nil,
  }, nil
}
