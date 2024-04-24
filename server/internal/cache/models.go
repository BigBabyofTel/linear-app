package cache

import "github.com/redis/go-redis/v9"

type RedisCache struct {
	User      *UserCache
	Workspace *WorkspaceCache
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{
		User:      &UserCache{client: client},
		Workspace: &WorkspaceCache{client: client},
	}
}
