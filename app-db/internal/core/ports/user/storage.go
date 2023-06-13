package user

import (
	"context"
	"garination.com/db/internal/platform/postgres"
)

var _ UserPostgresStorage = &postgres.Queries{}

type UserPostgresStorage interface {
	GetUserMeta(ctx context.Context, userID string) (postgres.UserMetum, error)
	InsertUserMeta(ctx context.Context, arg postgres.InsertUserMetaParams) (postgres.UserMetum, error)
	UpdateUserMeta(ctx context.Context, arg postgres.UpdateUserMetaParams) (postgres.UserMetum, error)
}
