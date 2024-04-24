package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/lucabrx/wuhu/internal/data"
	"github.com/redis/go-redis/v9"
)

type UserCache struct {
	client *redis.Client
}

const userCacheKey = "user:"

func (c *UserCache) SetUserByToken(ctx context.Context, token string, user *data.User) error {
	key := userCacheKey + token
	value, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, key, value, 30*time.Minute).Err()
}

func (c *UserCache) GetUserByToken(ctx context.Context, token string) (*data.User, error) {
	key := userCacheKey + token
	result, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil // Cache miss
	} else if err != nil {
		return nil, err
	}

	var user data.User
	if err := json.Unmarshal([]byte(result), &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *UserCache) DelUserByToken(ctx context.Context, token string) error {
	key := userCacheKey + token
	return c.client.Del(ctx, key).Err()
}
