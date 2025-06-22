package config

import (
	"os"	
	"strconv"
)

// RedisConfig holds the configuration for Redis
type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

// load environment variables 
func GetRedisConfig() RedisConfig {
	return RedisConfig{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       getRedisDB(),
	}
}

// getRedisDB retrieves the Redis DB number from environment variables or defaults to 0
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
