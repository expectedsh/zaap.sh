package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	CloudflareToken  string `envconfig:"CLOUDFLARE_TOKEN"`
	CloudflareZoneId string `envconfig:"CLOUDFLARE_ZONE_ID"`
}

func FromEnv() (*Config, error) {
	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}
	return config, nil
}
