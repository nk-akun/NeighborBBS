package util

import "github.com/go-redis/redis"

var rdb *redis.Client

// OpenRedis create a connection of redis
func OpenRedis() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
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
