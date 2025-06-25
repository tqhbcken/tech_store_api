package database

import (
	"api_techstore/internal/config"
	"log"

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
	if err := client.Ping(client.Context()).Err(); err != nil {
		log.Println("Failed to connect to Redis:", err)
		return nil, err
	}

	log.Println("Connected to Redis at", redisConfig.Addr)

	return client, nil
}
