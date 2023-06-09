package ports

import (
	"context"
	"garination.com/gateway/internal/core/auth/model"
)

type AuthRedisRepo interface {
	GetUserMeta(ctx context.Context, userID string) (*model.UserMetum, error)
	InsertUserMeta(ctx context.Context, arg model.UserMetum) (*model.UserMetum, error)
}
