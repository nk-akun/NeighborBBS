package util

import (
	"fmt"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

// OpenRedis create a connection of redis
func OpenRedis(host string, port string, password string) error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0, // use default DB
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// GetRedis return the pointer of db
func GetRedis() *redis.Client {
	return rdb
}
