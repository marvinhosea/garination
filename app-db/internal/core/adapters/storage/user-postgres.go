package storage

import (
	"context"
	"garination.com/db/internal/core/ports/user"
	"garination.com/db/internal/platform/postgres"
)

type userPostgresStorage struct {
	conn *postgres.Connection
}

func (u userPostgresStorage) UpdateUserDealership(ctx context.Context, arg postgres.UpdateUserDealershipParams) (postgres.UserMetum, error) {
	return u.conn.Queries.UpdateUserDealership(ctx, arg)
}

func (u userPostgresStorage) GetUserDealership(ctx context.Context, userID string) (postgres.Dealership, error) {
	return u.conn.Queries.GetUserDealership(ctx, userID)
}

func (u userPostgresStorage) GetUserMeta(ctx context.Context, userID string) (postgres.UserMetum, error) {
	return u.conn.Queries.GetUserMeta(ctx, userID)
}

func (u userPostgresStorage) InsertUserMeta(ctx context.Context, arg postgres.InsertUserMetaParams) (postgres.UserMetum, error) {
	if len(arg.DealershipID.String) == 0 {
		argsWithoutDealershipID := postgres.InsertUserMetaWithoutDealershipParams{
			UserMetaID:   arg.UserMetaID,
			UserID:       arg.UserID,
			FacebookUrl:  arg.FacebookUrl,
			TwitterUrl:   arg.TwitterUrl,
			InstagramUrl: arg.InstagramUrl,
			LinkedinUrl:  arg.LinkedinUrl,
			WebsiteUrl:   arg.WebsiteUrl,
		}

		return u.conn.Queries.InsertUserMetaWithoutDealership(ctx, argsWithoutDealershipID)
	}
	return u.conn.Queries.InsertUserMeta(ctx, arg)
}

func (u userPostgresStorage) UpdateUserMeta(ctx context.Context, arg postgres.UpdateUserMetaParams) (postgres.UserMetum, error) {
	return u.conn.Queries.UpdateUserMeta(ctx, arg)
}

func NewUserPostgresStorage(conn *postgres.Connection) user.PostgresStorage {
	return &userPostgresStorage{conn: conn}
}
