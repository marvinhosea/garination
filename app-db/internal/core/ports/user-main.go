package ports

import (
	"context"
	"garination.com/db/internal/core/model"
)

type UserRepository interface {
	GetUserDealership(ctx context.Context, userID string) (*model.Dealership, error)
	GetUserMeta(ctx context.Context, userID string) (*model.UserMetum, error)
	InsertDealership(ctx context.Context, arg model.Dealership) (*model.Dealership, error)
	InsertUserMeta(ctx context.Context, arg model.UserMetum) (*model.UserMetum, error)
	UpdateUserMeta(ctx context.Context, arg model.UserMetum) (*model.UserMetum, error)
}

type UserService interface {
	GetUserDealership(ctx context.Context, userID string) (*model.Dealership, error)
	GetUserMeta(ctx context.Context, userID string) (*model.UserMetum, error)
	InsertDealership(ctx context.Context, arg model.Dealership) (*model.Dealership, error)
	InsertUserMeta(ctx context.Context, arg model.UserMetum) (*model.UserMetum, error)
	UpdateUserMeta(ctx context.Context, arg model.UserMetum) (*model.UserMetum, error)
}
