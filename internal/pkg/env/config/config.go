package config

import (
	cache "airbnb-auth-be/internal/pkg/cache/config"
	credential "airbnb-auth-be/internal/pkg/credential/config"
	elastic "airbnb-auth-be/internal/pkg/elasticsearch/config"
	gorm "airbnb-auth-be/internal/pkg/gorm/config"
	httpserver "airbnb-auth-be/internal/pkg/http/server/config"
	jwt "airbnb-auth-be/internal/pkg/jwt/config"
	kafka "airbnb-auth-be/internal/pkg/kafka/config"
	oauth "airbnb-auth-be/internal/pkg/oauth/config"
	svcuser "airbnb-auth-be/internal/pkg/svcuser/config"
)

type Config struct {
	Origins    []string          `mapstructure:"origins"`
	Stage      string            `mapstructure:"stage"`
	Domain     string            `mapstructure:"domain"`
	Creds      credential.Config `mapstructure:"creds"`
	HttpServer httpserver.Config `mapstructure:"httpserver"`
	DB         gorm.Config       `mapstructure:"db"`
	Oauth      oauth.Config      `mapstructure:"oauth"`
	Jwt        jwt.Config        `mapstructure:"jwt"`
	Cache      cache.Config      `mapstructure:"cache"`
	Kafka      kafka.Config      `mapstructure:"kafka"`
	SvcUser    svcuser.Config    `mapstructure:"svcuser"`
	Elastic    elastic.Config    `mapstructure:"elastic"`
}
