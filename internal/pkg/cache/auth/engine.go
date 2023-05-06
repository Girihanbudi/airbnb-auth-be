package auth

import (
	"encoding/json"
	"time"
)

func Set(key string, value interface{}, expiration time.Duration) error {
	exp := expiration / 60 * time.Minute
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = Cache.Set(key, p, exp).Err()
	return err
}

func Get(key string, destination interface{}) error {
	p, err := Cache.Get(key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(p, destination)
}

func Del(keys ...string) error {
	_, err := Cache.Del(keys...).Result()
	return err
}
