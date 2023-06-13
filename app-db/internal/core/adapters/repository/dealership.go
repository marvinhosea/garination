package repository

import (
	"context"
	"garination.com/db/internal/core/ports/dealership"
	"garination.com/db/internal/platform/postgres"
)

type dealershipRepo struct {
	dealershipPostgresStorage dealership.DealershipPostgresStorage
}

func (d dealershipRepo) GetUserDealership(ctx context.Context, userID string) (postgres.Dealership, error) {
	return d.dealershipPostgresStorage.GetUserDealership(ctx, userID)
}

func (d dealershipRepo) InsertDealership(ctx context.Context, arg postgres.InsertDealershipParams) (postgres.Dealership, error) {
	return d.dealershipPostgresStorage.InsertDealership(ctx, arg)
}

func (d dealershipRepo) UpdateDealership(ctx context.Context, arg postgres.UpdateDealershipParams) (postgres.Dealership, error) {
	return d.dealershipPostgresStorage.UpdateDealership(ctx, arg)
}

func (d dealershipRepo) GetDealershipByID(ctx context.Context, id string) (postgres.Dealership, error) {
	return d.dealershipPostgresStorage.GetDealershipByID(ctx, id)
}

func (d dealershipRepo) DeleteDealership(ctx context.Context, id string) error {
	return d.dealershipPostgresStorage.DeleteDealership(ctx, id)
}

func NewDealershipRepo(dealershipPostgresStorage dealership.DealershipPostgresStorage) dealership.DealershipRepository {
	return &dealershipRepo{dealershipPostgresStorage: dealershipPostgresStorage}
}
