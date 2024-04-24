package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/lucabrx/wuhu/internal/data"
	"github.com/redis/go-redis/v9"
)

type WorkspaceCache struct {
	client *redis.Client
}

const (
	workspaceSingleCacheKey = "workspace:"
	myWorkspacesCacheKey    = "my-workspaces:"
	workspaceUsers          = "workspace-users:"
)

func (c *WorkspaceCache) SetWorkspace(ctx context.Context, userId int64, workspace *data.Workspace) error {
	key := workspaceSingleCacheKey + workspace.Slug + ":" + fmt.Sprint(userId)
	value, err := json.Marshal(workspace)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, key, value, 30*time.Minute).Err()
}

func (c *WorkspaceCache) GetWorkspace(ctx context.Context, userId int64, slug string) (*data.Workspace, error) {
	key := workspaceSingleCacheKey + slug + ":" + fmt.Sprint(userId)
	result, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var workspace data.Workspace
	if err := json.Unmarshal([]byte(result), &workspace); err != nil {
		return nil, err
	}
	return &workspace, nil
}

func (c *WorkspaceCache) DelWorkspace(ctx context.Context, userId int64, slug string) error {
	key := workspaceSingleCacheKey + slug + ":" + fmt.Sprint(userId)
	return c.client.Del(ctx, key).Err()
}

func (c *WorkspaceCache) SetMyWorkspaces(ctx context.Context, userId int64, workspaces []*data.Workspace) error {
	key := myWorkspacesCacheKey + fmt.Sprint(userId)
	value, err := json.Marshal(workspaces)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, key, value, 30*time.Minute).Err()
}

func (c *WorkspaceCache) GetMyWorkspaces(ctx context.Context, userId int64) ([]*data.Workspace, error) {
	key := myWorkspacesCacheKey + fmt.Sprint(userId)
	result, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var workspaces []*data.Workspace
	if err := json.Unmarshal([]byte(result), &workspaces); err != nil {
		return nil, err
	}
	return workspaces, nil
}

func (c *WorkspaceCache) DelMyWorkspaces(ctx context.Context, userId int64) error {
	key := myWorkspacesCacheKey + fmt.Sprint(userId)
	return c.client.Del(ctx, key).Err()
}

func (c *WorkspaceCache) SetWorkspaceUsers(ctx context.Context, workspaceId int64, users []*data.User) error {
	key := workspaceUsers + fmt.Sprint(workspaceId)
	value, err := json.Marshal(users)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, key, value, 30*time.Minute).Err()
}

func (c *WorkspaceCache) GetWorkspaceUsers(ctx context.Context, workspaceId int64) ([]*data.User, error) {
	key := workspaceUsers + fmt.Sprint(workspaceId)
	result, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var users []*data.User
	if err := json.Unmarshal([]byte(result), &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (c *WorkspaceCache) DelWorkspaceUsers(ctx context.Context, workspaceId int64) error {
	key := workspaceUsers + fmt.Sprint(workspaceId)
	return c.client.Del(ctx, key).Err()
}
