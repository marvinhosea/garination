package user

import (
	"context"
	"garination.com/db/internal/core/model"
)

type UserRepository interface {
	GetUserMeta(ctx context.Context, userID string) (*model.UserMetum, error)
	InsertUserMeta(ctx context.Context, arg model.UserMetum) (*model.UserMetum, error)
	UpdateUserMeta(ctx context.Context, arg model.UserMetum) (*model.UserMetum, error)
}
