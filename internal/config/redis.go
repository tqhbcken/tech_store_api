package config

import (
	"os"
	"strconv"
)

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

func GetRedisConfig() RedisConfig {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "localhost" // Default host
	}
	port := os.Getenv("REDIS_PORT")
	if port == "" {
		port = "6379" // Default port
	}

	return RedisConfig{
		Addr:     host + ":" + port,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       getRedisDB(),
	}
}

func getRedisDB() int {
	db := os.Getenv("REDIS_DB")
	if db == "" {
		return 0
	}
	dbNum, err := strconv.Atoi(db)
	if err != nil {
		return 0
	}
	return dbNum
}
