package main

import (
  "github.com/jinzhu/gorm"
  "github.com/kelseyhightower/envconfig"
  "github.com/remicaumette/zaap.sh/pkg/server"
  "golang.org/x/oauth2"
  "golang.org/x/oauth2/github"
  "golang.org/x/oauth2/google"
  "net/http"
  "time"
)

type HttpConfig struct {
  Addr string `envconfig:"ADDR" default:":3000"`
}

type DBConfig struct {
  Addr            string        `envconfig:"ADDR" default:"postgres://zaap:zaap@localhost/zaap?sslmode=disable"`
  ConnMaxLifetime time.Duration `envconfig:"CONNMAXLIFETIME" default:"5m"`
  MaxIdleConns    int           `envconfig:"MAXIDLECONNS" default:"4"`
  MaxOpenConns    int           `envconfig:"MAXOPENCONNS" default:"100"`
}

type OAuthConfig struct {
  ClientID     string `envconfig:"CLIENT_ID"`
  ClientSecret string `envconfig:"CLIENT_SECRET"`
}

func newDB() (*gorm.DB, error) {
  config := &DBConfig{}
  if err := envconfig.Process("DB", config); err != nil {
    return nil, err
  }
  db, err := gorm.Open("postgres", config.Addr)
  if err != nil {
    return nil, err
  }
  db.DB().SetMaxIdleConns(config.MaxIdleConns)
  db.DB().SetMaxOpenConns(config.MaxOpenConns)
  db.DB().SetConnMaxLifetime(config.ConnMaxLifetime)
  return db, nil
}

func newHttpServer() (*http.Server, error) {
  config := &HttpConfig{}
  if err := envconfig.Process("HTTP", config); err != nil {
    return nil, err
  }
  return &http.Server{
    Addr: config.Addr,
  }, nil
}

func newOAuthConfig(prefix string, endpoint oauth2.Endpoint, scopes []string) (*oauth2.Config, error) {
  config := &OAuthConfig{}
  if err := envconfig.Process(prefix, config); err != nil {
    return nil, err
  }
  return &oauth2.Config{
    ClientID:     config.ClientID,
    ClientSecret: config.ClientSecret,
    Endpoint:     endpoint,
    Scopes:       scopes,
  }, nil
}

func newServerFromEnv() (*server.Server, error) {
  srv := &server.Server{}

  httpServer, err := newHttpServer()
  if err != nil {
    return nil, err
  }
  srv.HttpServer = httpServer

  db, err := newDB()
  if err != nil {
    return nil, err
  }
  srv.DB = db

  oauthGithubConfig, err := newOAuthConfig("GITHUB", github.Endpoint, []string{"user", "user:email"})
  if err != nil {
    return nil, err
  }
  srv.GithubOAuthConfig = oauthGithubConfig

  oauthGoogleConfig, err := newOAuthConfig("GOOGLE", google.Endpoint, []string{
    "https://www.googleapis.com/auth/userinfo.email",
    "https://www.googleapis.com/auth/userinfo.profile",
  })
  if err != nil {
    return nil, err
  }
  srv.GoogleOAuthConfig = oauthGoogleConfig

  return srv, nil
}
