package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Config struct {
	Region                 string `env:"REGION" envDefault:"us-west-2"`
	UserPoolId             string `env:"USER_POOL_ID" envDefault:""`
	UserPoolWebClientId    string `env:"USER_POOL_WEB_CLIENT_ID" envDefault:""`
	AuthenticationFlowType string `env:"USER_PASSWORD_AUTH" envDefault:"USER_PASSWORD_AUTH"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return &cfg, nil
}
