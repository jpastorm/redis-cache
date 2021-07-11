package connection

import "github.com/go-redis/redis/v8"

func NewRedis () *redis.Client{
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return redisClient
}