package storage

import (
	"context"
	"garination.com/db/internal/core/ports/dealership"
	"garination.com/db/internal/platform/postgres"
)

type dealershipPostgresStorage struct {
	conn *postgres.Connection
}

func (d dealershipPostgresStorage) GetUserDealership(ctx context.Context, userID string) (postgres.Dealership, error) {
	return d.conn.Queries.GetUserDealership(ctx, userID)
}

func (d dealershipPostgresStorage) InsertDealership(ctx context.Context, arg postgres.InsertDealershipParams) (postgres.Dealership, error) {
	return d.conn.Queries.InsertDealership(ctx, arg)
}

func (d dealershipPostgresStorage) UpdateDealership(ctx context.Context, arg postgres.UpdateDealershipParams) (postgres.Dealership, error) {
	return d.conn.Queries.UpdateDealership(ctx, arg)
}

func (d dealershipPostgresStorage) GetDealershipByID(ctx context.Context, id string) (postgres.Dealership, error) {
	return d.conn.Queries.GetDealershipById(ctx, id)
}

func (d dealershipPostgresStorage) DeleteDealership(ctx context.Context, id string) error {
	_, err := d.conn.Queries.DeleteDealership(ctx, id)
	return err
}

func NewDealershipPostgresStorage(conn *postgres.Connection) dealership.DealershipPostgresStorage {
	return &dealershipPostgresStorage{conn: conn}
}
