package server

import (
  "context"
  "errors"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/remicaumette/zaap.sh/pkg/models"
  "github.com/remicaumette/zaap.sh/pkg/util/httpx"
  "github.com/remicaumette/zaap.sh/pkg/util/httpx/middleware"
  "golang.org/x/oauth2"
  "golang.org/x/oauth2/github"
  "golang.org/x/oauth2/google"
  "net/http"
)

type Server struct {
  config            *Config
  db                *gorm.DB
  httpServer        *http.Server
  googleOAuthConfig *oauth2.Config
  githubOAuthConfig *oauth2.Config
}

func New(config *Config) *Server {
  return &Server{config: config}
}

func (s *Server) setupDatabase() error {
  db, err := gorm.Open("postgres", s.config.DB.Addr)
  if err != nil {
    return err
  }
  db.DB().SetMaxIdleConns(s.config.DB.MaxIdleConns)
  db.DB().SetMaxOpenConns(s.config.DB.MaxOpenConns)
  db.DB().SetConnMaxLifetime(s.config.DB.ConnMaxLifetime)
  if err := db.AutoMigrate(&models.User{}).Error; err != nil {
    return err
  }
  s.db = db
  return nil
}

func (s *Server) Start() error {
  if s.httpServer != nil {
    return errors.New("server already started")
  }
  if err := s.setupDatabase(); err != nil {
    return err
  }

  s.githubOAuthConfig = &oauth2.Config{
    ClientID:     s.config.Github.ClientID,
    ClientSecret: s.config.Github.ClientSecret,
    Endpoint:     github.Endpoint,
    Scopes:       []string{"user", "user:email"},
  }
  s.googleOAuthConfig = &oauth2.Config{
    ClientID:     s.config.Google.ClientID,
    ClientSecret: s.config.Google.ClientSecret,
    Endpoint:     google.Endpoint,
    Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
  }

  handler := httpx.NewHandler()
  handler.Use(middleware.CanonicalLog)
  handler.Get("/oauth/github", s.OAuthGithubRoute)
  handler.Post("/oauth/github/callback", s.OAuthGithubCallbackRoute)
  v1 := handler.Group("/v1")
  {
    v1.Get("/", func(ctx *httpx.Context) {
      ctx.Json(http.StatusOK, "ok")
    })
  }

  s.httpServer = &http.Server{
    Addr:    s.config.Addr,
    Handler: handler,
  }
  return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
  return s.httpServer.Shutdown(ctx)
}
