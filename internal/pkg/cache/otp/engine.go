package otp

import "time"

func Set(key string, value interface{}, expiration time.Duration) error {
	exp := expiration / 60 * time.Minute
	_, err := Cache.Set(key, value, exp).Result()
	return err
}

func Get(key string) (string, error) {
	return Cache.Get(key).Result()
}
