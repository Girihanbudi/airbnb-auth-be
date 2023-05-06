package auth

import (
	"airbnb-auth-be/internal/pkg/env"
	"airbnb-auth-be/internal/pkg/log"
	"fmt"

	"github.com/go-redis/redis"
)

const Instance string = "Auth Cache"

// global auth cache declaration
var Cache *redis.Client

func InitAuthCache() {
	client := redis.NewClient(&redis.Options{
		Addr:     env.CONFIG.Cache.Auth.Host,
		Password: env.CONFIG.Cache.Auth.Password,
		DB:       env.CONFIG.Cache.Auth.Db,
	})

	log.Event(Instance, fmt.Sprintf("connected to %s", env.CONFIG.Cache.Auth.Host))

	Cache = client
}
