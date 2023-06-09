package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"garination.com/gateway/internal/core/auth/model"
	"garination.com/gateway/internal/core/auth/ports"
	"github.com/redis/go-redis/v9"
)

// DefaultCacheExpiry is the default expiry time for cache (1 hour)
const DefaultCacheExpiry = 3600

type redisRepo struct {
	client redis.Client
}

func (r redisRepo) GetUserMeta(ctx context.Context, userID string) (*model.UserMetum, error) {
	userMetaBytes, err := r.client.Get(ctx, userID).Bytes()
	if err != nil {
		return nil, err
	}

	var userMeta model.UserMetum

	// unmarshall  the json bytes into the userMeta struct
	err = json.Unmarshal(userMetaBytes, &userMeta)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal user meta: %w", err)
	}

	return &userMeta, nil
}

func (r redisRepo) InsertUserMeta(ctx context.Context, arg model.UserMetum) (*model.UserMetum, error) {
	// marshall the userMeta struct into json bytes
	userMetaBytes, err := json.Marshal(arg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user meta: %w", err)
	}

	// set the userMeta in redis
	err = r.client.Set(ctx, arg.UserID, userMetaBytes, DefaultCacheExpiry).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to set user meta: %w", err)
	}

	return &arg, nil
}

func NewRedisRepo(client redis.Client) ports.AuthRedisRepo {
	return &redisRepo{client: client}
}
