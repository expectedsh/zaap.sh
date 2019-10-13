package server

import "github.com/kelseyhightower/envconfig"

type Config struct {

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
