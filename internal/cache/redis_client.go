package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisClient wraps the Redis client
type RedisClient struct {
	client *redis.Client
}

// NewRedisClient creates a new RedisClient
func NewRedisClient(client *redis.Client) *RedisClient {
	return &RedisClient{client: client}
}

// SetToken save accessUUID
func (r *RedisClient) SetToken(ctx context.Context, uuid, userID string, ttl time.Duration) error {
	return r.client.Set(ctx, uuid, userID, ttl).Err()
}

func (r *RedisClient) DeleteToken(ctx context.Context, uuid string) error {
	return r.client.Del(ctx, uuid).Err()
}

func (r *RedisClient) IsValidToken(ctx context.Context, uuid string) (bool, error) {
	exists, err := r.client.Exists(ctx, uuid).Result()
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}
