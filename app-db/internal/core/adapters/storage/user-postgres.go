package storage

import (
	"context"
	"garination.com/db/internal/core/ports"
	"garination.com/db/internal/platform/postgres"
)

type userPostgresStorage struct {
	conn *postgres.Connection
}

func (u userPostgresStorage) GetUserDealership(ctx context.Context, userID string) (postgres.Dealership, error) {
	return u.conn.Queries.GetUserDealership(ctx, userID)
}

func (u userPostgresStorage) GetUserMeta(ctx context.Context, userID string) (postgres.UserMetum, error) {
	return u.conn.Queries.GetUserMeta(ctx, userID)
}

func (u userPostgresStorage) InsertDealership(ctx context.Context, arg postgres.InsertDealershipParams) (postgres.Dealership, error) {
	return u.conn.Queries.InsertDealership(ctx, arg)
}

func (u userPostgresStorage) InsertUserMeta(ctx context.Context, arg postgres.InsertUserMetaParams) (postgres.UserMetum, error) {
	return u.conn.Queries.InsertUserMeta(ctx, arg)
}

func (u userPostgresStorage) UpdateUserMeta(ctx context.Context, arg postgres.UpdateUserMetaParams) (postgres.UserMetum, error) {
	return u.conn.Queries.UpdateUserMeta(ctx, arg)
}

func NewUserPostgresStorage(conn *postgres.Connection) ports.UserPostgresStorage {
	return &userPostgresStorage{conn: conn}
}
