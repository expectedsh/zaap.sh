package app

import (
  "context"
  "github.com/remicaumette/zaap.sh/pkg/protocol"
)

type Service struct {}

func (s *Service) ListApps(ctx context.Context, r *protocol.ListAppsRequest) (*protocol.ListAppsResponse, error) {
  return &protocol.ListAppsResponse{
    Apps: nil,
  }, nil
}

func (s *Service) GetApp(ctx context.Context, r *protocol.GetAppRequest) (*protocol.App, error) {
  //for _, app := range apps {
  //  if app.Name == r.Id {
  //    return app, nil
  //  }
  //}
  return &protocol.App{}, nil
}
