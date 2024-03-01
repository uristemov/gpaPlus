package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/uristemov/repeatPro/internal/config"
)

func NewRedisClient(cfg *config.Config) (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port), Password: "", DB: 0})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	return client, nil
}
