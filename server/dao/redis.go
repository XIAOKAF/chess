package dao

import (
	"time"
)

func Set(key string, val string, expiration time.Duration) error {
	err := RDB.Set(key, val, expiration*time.Minute)
	return err.Err()
}

func Get(key string) (string, error) {
	result, err := RDB.Get(key).Result()
	return result, err
}

func HashSet(name string, key string, value string) error {
	result := RDB.HSet(name, key, value)
	return result.Err()
}
