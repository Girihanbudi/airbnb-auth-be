package config

import (
	facebook "airbnb-auth-be/internal/pkg/oauth/facebook/config"
	google "airbnb-auth-be/internal/pkg/oauth/google/config"
)

type Config struct {
	RedirectUrl string          `mapstructure:"redirecturl"`
	Google      google.Config   `mapstructure:"google"`
	Facebook    facebook.Config `mapstructure:"facebook"`
}
