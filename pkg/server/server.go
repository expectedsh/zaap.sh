package server

import (
  "context"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/remicaumette/zaap.sh/pkg/util/httpx"
  "golang.org/x/oauth2"
  "net/http"
)

type Server struct {
  Secret            []byte
  HttpServer        *http.Server
  DB                *gorm.DB
  GoogleOAuthConfig *oauth2.Config
  GithubOAuthConfig *oauth2.Config
}

func (s *Server) Start() error {
  handler := httpx.NewHandler()
  handler.Get("/oauth/github", s.OAuthGithubRoute)
  handler.Get("/oauth/github/callback", s.OAuthGithubCallbackRoute)

  s.HttpServer.Handler = handler
  return s.HttpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
  return s.HttpServer.Shutdown(ctx)
}
