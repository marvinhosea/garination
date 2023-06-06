package ports

import (
	"context"
	"garination.com/db/internal/platform/postgres"
)

var _ UserPostgresStorage = &postgres.Queries{}

type UserPostgresStorage interface {
	GetUserDealership(ctx context.Context, userID string) (postgres.Dealership, error)
	GetUserMeta(ctx context.Context, userID string) (postgres.UserMetum, error)
	InsertDealership(ctx context.Context, arg postgres.InsertDealershipParams) (postgres.Dealership, error)
	InsertUserMeta(ctx context.Context, arg postgres.InsertUserMetaParams) (postgres.UserMetum, error)
	UpdateUserMeta(ctx context.Context, arg postgres.UpdateUserMetaParams) (postgres.UserMetum, error)
}
