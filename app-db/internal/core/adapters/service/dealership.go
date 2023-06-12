package service

import (
	"context"
	"garination.com/db/internal/core/ports/dealership"
	"garination.com/db/internal/core/ports/user"
	"garination.com/db/internal/platform/postgres"
)

type dealershipService struct {
	dealershipRepo dealership.DealershipRepository
	userRepo       user.UserRepository
}

func (d dealershipService) GetUserDealership(ctx context.Context, userID string) (postgres.Dealership, error) {
	return d.dealershipRepo.GetUserDealership(ctx, userID)
}

func (d dealershipService) InsertDealership(ctx context.Context, arg postgres.InsertDealershipParams) (postgres.Dealership, error) {
	dealership, err := d.dealershipRepo.InsertDealership(ctx, arg)
	if err != nil {
		return postgres.Dealership{}, err
	}

	// update user meta
	userMeta, err := d.userRepo.GetUserMeta(ctx, arg.OwnerID)
	if err != nil {
		// rollback dealership
		err = d.dealershipRepo.DeleteDealership(ctx, dealership.DealershipID)
		if err != nil {
			return postgres.Dealership{}, err
		}
		return postgres.Dealership{}, err
	}

	userMeta.DealershipID.String = dealership.DealershipID
	userMeta.DealershipID.Valid = true

	_, err = d.userRepo.UpdateUserDealership(ctx, userMeta.UserID, dealership.DealershipID)
	if err != nil {
		// rollback dealership
		err = d.dealershipRepo.DeleteDealership(ctx, dealership.DealershipID)
		if err != nil {
			return postgres.Dealership{}, err
		}
		return postgres.Dealership{}, err
	}

	return dealership, nil
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

func NewDealershipService(dealershipRepo dealership.DealershipRepository, userRepo user.UserRepository) dealership.DealershipService {
	return &dealershipService{dealershipRepo: dealershipRepo, userRepo: userRepo}
}
