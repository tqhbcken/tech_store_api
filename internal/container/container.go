package container

import (
	"api_techstore/internal/database"
	"api_techstore/pkg/jwt"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Container struct {
	DB        *gorm.DB
	Redis     *redis.Client
	JWTConfig *jwt.JWTConfig
	Logger    *logrus.Logger
}

func NewContainer() (*Container, error){
	db, err := database.InitDB()
	if err != nil {
		return nil, err
	}

	redisClient, err := database.InitRedis()
	if err != nil {
		return nil, err
	}

	jwtCfg := jwt.NewJWTConfig()

	logger := logrus.New()

	return &Container{
		DB:        db.DB,
		Redis:     redisClient,
		JWTConfig: jwtCfg,
		Logger:    logger,
	}, nil
}