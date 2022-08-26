package redis

import "github.com/go-redis/redis"

func NewRedisClient() *redis.Client {
	var redisHost = "localhost:6379"
	var redisPassword = ""
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       0,
	})
	return client
}

