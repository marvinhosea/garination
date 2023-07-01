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

func (d dealershipService) GetUserDealership(ctx context.Context, userID string) (*postgres.Dealership, error) {
	return d.dealershipRepo.GetUserDealership(ctx, userID)
}

func (d dealershipService) InsertDealership(ctx context.Context, arg postgres.InsertDealershipParams) (*postgres.Dealership, error) {
	insertDealership, err := d.dealershipRepo.InsertDealership(ctx, arg)
	if err != nil {
		return nil, err
	}

	// update user meta
	userMeta, err := d.userRepo.GetUserMeta(ctx, arg.OwnerID)
	if err != nil {
		// rollback insertDealership
		err = d.dealershipRepo.DeleteDealership(ctx, insertDealership.DealershipID)
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	userMeta.DealershipID.String = insertDealership.DealershipID
	userMeta.DealershipID.Valid = true

	updateDealershipArg := postgres.UpdateUserDealershipParams{
		UserID:       userMeta.UserID,
		DealershipID: userMeta.DealershipID,
	}

	_, err = d.userRepo.UpdateUserDealership(ctx, updateDealershipArg)
	if err != nil {
		// rollback insertDealership
		err = d.dealershipRepo.DeleteDealership(ctx, insertDealership.DealershipID)
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	return insertDealership, nil
}

func (d dealershipService) UpdateDealership(ctx context.Context, arg postgres.UpdateDealershipParams) (*postgres.Dealership, error) {
	return d.dealershipRepo.UpdateDealership(ctx, arg)
}

func (d dealershipService) GetDealershipByID(ctx context.Context, id string) (*postgres.Dealership, error) {
	return d.dealershipRepo.GetDealershipByID(ctx, id)
}

func (d dealershipService) DeleteDealership(ctx context.Context, id string) error {
	return d.dealershipRepo.DeleteDealership(ctx, id)
}

func NewDealershipService(dealershipRepo dealership.DealershipRepository, userRepo user.UserRepository) dealership.DealershipService {
	return &dealershipService{dealershipRepo: dealershipRepo, userRepo: userRepo}
}
