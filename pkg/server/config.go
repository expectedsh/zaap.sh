package server

import (
  "github.com/kelseyhightower/envconfig"
  "time"
)

type Config struct {
  Addr string `envconfig:"ADDR" default:":3000"`
  DB   struct {
    Addr            string        `envconfig:"ADDR" default:"postgres://zaap:zaap@localhost/zaap?sslmode=disable"`
    ConnMaxLifetime time.Duration `envconfig:"CONNMAXLIFETIME" default:"5m"`
    MaxIdleConns    int           `envconfig:"MAXIDLECONNS" default:"4"`
    MaxOpenConns    int           `envconfig:"MAXOPENCONNS" default:"100"`
  }
  Github struct {
    ClientID     string `envconfig:"CLIENT_ID"`
    ClientSecret string `envconfig:"CLIENT_SECRET"`
  }
  Google struct {
    ClientID     string `envconfig:"CLIENT_ID"`
    ClientSecret string `envconfig:"CLIENT_SECRET"`
  }
}

func NewConfig() *Config {
  return &Config{}
}

func NewConfigFromEnviron() (*Config, error) {
  config := NewConfig()
  if err := envconfig.Process("", config); err != nil {
    return nil, err
  }
  return config, nil
}
