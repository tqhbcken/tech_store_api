package database

import (
	"api_techstore/internal/config"

	"github.com/go-redis/redis/v8"
)

func InitRedis() (*redis.Client, error) {
	redisConfig := config.GetRedisConfig()

	
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password, 
		DB:       redisConfig.DB,       
	})

	// Test the connection
	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
