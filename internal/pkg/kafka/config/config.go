package config

import (
	consumer "airbnb-auth-be/internal/pkg/kafka/consumer/config"
)

type Config struct {
	ClientId string          `mapstructure:"clientid"`
	Brokers  []string        `mapstructure:"brokers"`
	Version  string          `mapstructure:"version"`
	Consumer consumer.Config `mapstructure:"consumer"`
}
