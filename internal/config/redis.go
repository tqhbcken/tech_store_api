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
	return RedisConfig{
		Addr:     os.Getenv("REDIS_ADDRESS"),
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
