package cache

import (
	"context"
	"garination.com/gateway/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg config.Redis) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: "",
		DB:       0,
	})

	if status := rdb.Ping(context.Background()); status.Err() != nil {
		return nil, status.Err()
	}
	return rdb, nil
}
