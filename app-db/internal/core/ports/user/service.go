package user

import (
	"context"
	"garination.com/db/internal/platform/postgres"
)

type UserService interface {
	GetUserMeta(ctx context.Context, userID string) (*postgres.UserMetum, error)
	InsertUserMeta(ctx context.Context, arg postgres.UserMetum) (*postgres.UserMetum, error)
	UpdateUserMeta(ctx context.Context, arg postgres.UpdateUserMetaParams) (*postgres.UserMetum, error)
	UpdateUserDealership(ctx context.Context, args postgres.UpdateUserDealershipParams) (*postgres.UserMetum, error)
}
