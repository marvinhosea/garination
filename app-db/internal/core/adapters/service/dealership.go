package service

import (
	"context"
	"garination.com/db/internal/core/ports/dealership"
	"garination.com/db/internal/platform/postgres"
)

type dealershipService struct {
	dealershipRepo dealership.DealershipRepository
}

func (d dealershipService) GetUserDealership(ctx context.Context, userID string) (postgres.Dealership, error) {
	return d.dealershipRepo.GetUserDealership(ctx, userID)
}

func (d dealershipService) InsertDealership(ctx context.Context, arg postgres.InsertDealershipParams) (postgres.Dealership, error) {
	return d.dealershipRepo.InsertDealership(ctx, arg)
}

func (d dealershipService) UpdateDealership(ctx context.Context, arg postgres.UpdateDealershipParams) (postgres.Dealership, error) {
	return d.dealershipRepo.UpdateDealership(ctx, arg)
}

func (d dealershipService) GetDealershipByID(ctx context.Context, id string) (postgres.Dealership, error) {
	return d.dealershipRepo.GetDealershipByID(ctx, id)
}

func (d dealershipService) DeleteDealership(ctx context.Context, id string) error {
	return d.dealershipRepo.DeleteDealership(ctx, id)
}

func NewDealershipService(dealershipRepo dealership.DealershipRepository) dealership.DealershipService {
	return &dealershipService{dealershipRepo: dealershipRepo}
}
