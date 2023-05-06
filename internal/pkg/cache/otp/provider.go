package otp

import (
	"airbnb-auth-be/internal/pkg/env"
	"airbnb-auth-be/internal/pkg/log"
	"fmt"

	"github.com/go-redis/redis"
)

const Instance string = "OTP Cache"

// global auth cache declaration
var Cache *redis.Client

func InitOtpCache() {
	client := redis.NewClient(&redis.Options{
		Addr:     env.CONFIG.Cache.Otp.Host,
		Password: env.CONFIG.Cache.Otp.Password,
		DB:       env.CONFIG.Cache.Otp.Db,
	})

	log.Event(Instance, fmt.Sprintf("connected to %s", env.CONFIG.Cache.Otp.Host))

	Cache = client
}
