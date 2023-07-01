package repository

import (
	"context"
	"garination.com/db/internal/core/ports/user"
	"garination.com/db/internal/platform/postgres"
)

type userRepo struct {
	conn *postgres.Connection
}

func (u userRepo) InsertUserMeta(ctx context.Context, arg postgres.InsertUserMetaParams) (*postgres.UserMetum, error) {
	res, err := u.conn.Queries.InsertUserMeta(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (u userRepo) InsertUserMetaWithoutDealershipParams(ctx context.Context, arg postgres.InsertUserMetaWithoutDealershipParams) (*postgres.UserMetum, error) {
	res, err := u.conn.Queries.InsertUserMetaWithoutDealership(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (u userRepo) GetUserMeta(ctx context.Context, userID string) (*postgres.UserMetum, error) {
	res, err := u.conn.Queries.GetUserMeta(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (u userRepo) UpdateUserMeta(ctx context.Context, arg postgres.UpdateUserMetaParams) (*postgres.UserMetum, error) {
	res, err := u.conn.Queries.UpdateUserMeta(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (u userRepo) UpdateUserDealership(ctx context.Context, args postgres.UpdateUserDealershipParams) (*postgres.UserMetum, error) {
	res, err := u.conn.Queries.UpdateUserDealership(ctx, args)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func NewUserRepo(connection *postgres.Connection) user.UserRepository {
	return &userRepo{
		conn: connection,
	}
}
