package cache

import (
	"context"
	"fmt"
	"time"
)

type TokenCacher interface {
	GetToken(ctx context.Context, key string) (string, error)
	SetToken(ctx context.Context, key, refreshId string, expirationTime time.Duration) error
	DeleteToken(ctx context.Context, key string) error
}

func (c *Cache) GetToken(ctx context.Context, key string) (string, error) {
	key = fmt.Sprintf("refresh: %s", key)
	value, err := c.redisCli.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (c *Cache) SetToken(ctx context.Context, key, refreshId string, expirationTime time.Duration) error {
	key = fmt.Sprintf("refresh: %s", key)
	return c.redisCli.Set(ctx, key, refreshId, expirationTime).Err()
}

func (c *Cache) DeleteToken(ctx context.Context, key string) error {
	key = fmt.Sprintf("refresh: %s", key)
	return c.redisCli.Del(ctx, key).Err()
}
