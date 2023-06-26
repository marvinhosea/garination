package ports

import (
	"context"
	"garination.com/gateway/internal/core/auth/model"
)

type AuthDbRepo interface {
	GetUserMeta(ctx context.Context, userID string) (*model.UserMeta, error)
	InsertUserMeta(ctx context.Context, arg model.UserMeta) (*model.UserMeta, error)
	UpdateUserMeta(ctx context.Context, arg model.UserMeta) (*model.UserMeta, error)
	ChangeUserDealership(ctx context.Context, arg model.UserMeta) (string, error)
}
