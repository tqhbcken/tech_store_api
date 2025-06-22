package jwt

import (
	"api_techstore/internal/config"
	"os"
	"time"
)

type JWTConfig struct {
	SecretKey     string
	AccessTokenDuration  time.Duration
    RefreshTokenDuration time.Duration
}

func NewJWTConfig() *JWTConfig {

	
	config.LoadEnvVar()

	secret := os.Getenv("JWT_SECRET")
	accessStr := os.Getenv("JWT_ACCESS_DURATION")
	refreshStr := os.Getenv("JWT_REFRESH_DURATION")


	accessDuration, err := time.ParseDuration(accessStr)
	if err != nil {
		accessDuration = 15 * time.Minute
	}
	refreshDuration, err := time.ParseDuration(refreshStr)
	if err != nil {
		refreshDuration = 7 * 24 * time.Hour
	}

	return &JWTConfig{
		SecretKey:           secret,
		AccessTokenDuration: accessDuration,
		RefreshTokenDuration: refreshDuration,
	}
}
