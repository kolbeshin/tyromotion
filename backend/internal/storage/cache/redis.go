package cache

import (
	"github.com/redis/go-redis/v9"
	"tyromotion/backend/internal/config"
)

func NewRedisClient(config *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})
	return client
}
