package repository

import (
	"context"
	"garination.com/db/internal/core/ports/dealership"
	"garination.com/db/internal/platform/postgres"
)

type dealershipRepo struct {
	conn *postgres.Connection
}

func (d dealershipRepo) GetUserDealership(ctx context.Context, userID string) (*postgres.Dealership, error) {
	res, err := d.conn.Queries.GetUserDealership(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (d dealershipRepo) InsertDealership(ctx context.Context, arg postgres.InsertDealershipParams) (*postgres.Dealership, error) {
	res, err := d.conn.Queries.InsertDealership(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (d dealershipRepo) UpdateDealership(ctx context.Context, arg postgres.UpdateDealershipParams) (*postgres.Dealership, error) {
	res, err := d.conn.Queries.UpdateDealership(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (d dealershipRepo) GetDealershipByID(ctx context.Context, id string) (*postgres.Dealership, error) {
	res, err := d.conn.Queries.GetDealershipById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (d dealershipRepo) DeleteDealership(ctx context.Context, id string) error {
	_, err := d.conn.Queries.DeleteDealership(ctx, id)
	return err
}

func NewDealershipRepo(conn *postgres.Connection) dealership.DealershipRepository {
	return &dealershipRepo{conn: conn}
}
