package usecase

import (
	"context"
	"errors"
	"garination.com/gateway/internal/core/dealership/dto"
	"garination.com/gateway/internal/core/dealership/ports"
	"garination.com/gateway/internal/platform/grpc/app-db/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type dealershipUsecase struct {
	dealershipDatabaseRepo ports.DealershipDatabaseRepo
}

func (d dealershipUsecase) GetDealershipByID(ctx context.Context, getDealershipRequest *dto.GetDealershipRequest) (*dto.GetDealershipResponse, error) {
	if getDealershipRequest == nil {
		return nil, errors.New("it'd be nice if you could provide a valid request")
	}

	// get dealership from database
	dealership, err := d.dealershipDatabaseRepo.GetDealershipByID(ctx, getDealershipRequest.ID)
	if err != nil {
		return nil, err
	}

	// convert dealership to response
	response := &dto.GetDealershipResponse{
		DealershipID: dealership.DealershipId,
		OwnerID:      dealership.OwnerId,
		Name:         dealership.Name,
		DisplayName:  dealership.DisplayName,
		Address:      dealership.Address,
		City:         dealership.City,
		State:        dealership.State,
		Zip:          dealership.Zip,
		Phone:        dealership.Phone,
		Email:        dealership.Email,
		Website:      dealership.Website,
		FacebookUrl:  dealership.FacebookUrl,
		TwitterUrl:   dealership.TwitterUrl,
		InstagramUrl: dealership.InstagramUrl,
		LinkedinUrl:  "",
		LogoUrl:      dealership.LogoUrl,
		CoverUrl:     dealership.CoverUrl,
		Description:  dealership.Description,
		CreatedAt:    dealership.CreatedAt.AsTime(),
		UpdatedAt:    dealership.UpdatedAt.AsTime(),
	}

	return response, nil
}

func (d dealershipUsecase) GetDealershipByUserID(ctx context.Context, getDealershipRequest *dto.GetDealershipRequest) (*dto.GetDealershipResponse, error) {
	if getDealershipRequest == nil {
		return nil, errors.New("it'd be nice if you could provide a valid request")
	}

	// get dealership from database
	dealership, err := d.dealershipDatabaseRepo.GetDealershipByUserID(ctx, getDealershipRequest.ID)
	if err != nil {
		return nil, err
	}

	// convert dealership to response
	response := &dto.GetDealershipResponse{
		DealershipID: dealership.DealershipId,
		OwnerID:      dealership.OwnerId,
		Name:         dealership.Name,
		DisplayName:  dealership.DisplayName,
		Address:      dealership.Address,
		City:         dealership.City,
		State:        dealership.State,
		Zip:          dealership.Zip,
		Phone:        dealership.Phone,
		Email:        dealership.Email,
		Website:      dealership.Website,
		FacebookUrl:  dealership.FacebookUrl,
		TwitterUrl:   dealership.TwitterUrl,
		InstagramUrl: dealership.InstagramUrl,
		LinkedinUrl:  dealership.LinkedinUrl,
		LogoUrl:      dealership.LogoUrl,
		CoverUrl:     dealership.CoverUrl,
		Description:  dealership.Description,
		CreatedAt:    dealership.CreatedAt.AsTime(),
		UpdatedAt:    dealership.UpdatedAt.AsTime(),
	}

	return response, nil
}

func (d dealershipUsecase) UpdateDealership(ctx context.Context, updateDealershipRequest *dto.UpdateDealershipRequest) (*dto.UpdateDealershipResponse, error) {

	if updateDealershipRequest == nil {
		return nil, errors.New("it'd be nice if you could provide a valid request")
	}

	dealershipReq := &proto.Dealership{
		DealershipId: updateDealershipRequest.DealershipID,
		Name:         updateDealershipRequest.Name,
		DisplayName:  updateDealershipRequest.DisplayName,
		Address:      updateDealershipRequest.Address,
		City:         updateDealershipRequest.City,
		State:        updateDealershipRequest.State,
		Zip:          updateDealershipRequest.Zip,
		Phone:        updateDealershipRequest.Phone,
		Email:        updateDealershipRequest.Email,
		Website:      updateDealershipRequest.Website,
		FacebookUrl:  updateDealershipRequest.FacebookUrl,
		TwitterUrl:   updateDealershipRequest.TwitterUrl,
		InstagramUrl: updateDealershipRequest.InstagramUrl,
		LinkedinUrl:  updateDealershipRequest.LinkedinUrl,
		LogoUrl:      updateDealershipRequest.LogoUrl,
		CoverUrl:     updateDealershipRequest.CoverUrl,
		Description:  updateDealershipRequest.Description,
		UpdatedAt:    timestamppb.Now(),
	}

	// get dealership from database
	dealership, err := d.dealershipDatabaseRepo.UpdateDealership(ctx, dealershipReq)
	if err != nil {
		return nil, err
	}

	// convert dealership to response
	response := &dto.UpdateDealershipResponse{
		DealershipID: dealership.DealershipId,
		OwnerID:      dealership.OwnerId,
		Name:         dealership.Name,
		DisplayName:  dealership.DisplayName,
		Address:      dealership.Address,
		City:         dealership.City,
		State:        dealership.State,
		Zip:          dealership.Zip,
		Phone:        dealership.Phone,
		Email:        dealership.Email,
		Website:      dealership.Website,
		FacebookUrl:  dealership.FacebookUrl,
		TwitterUrl:   dealership.TwitterUrl,
		InstagramUrl: dealership.InstagramUrl,
		LinkedinUrl:  dealership.LinkedinUrl,
		LogoUrl:      dealership.LogoUrl,
		CoverUrl:     dealership.CoverUrl,
		Description:  dealership.Description,
		CreatedAt:    dealership.CreatedAt.AsTime(),
	}

	return response, nil
}

func (d dealershipUsecase) CreateDealership(ctx context.Context, createDealershipRequest *dto.CreateDealershipRequest) (*dto.CreateDealershipResponse, error) {
	if createDealershipRequest == nil {
		return nil, errors.New("it'd be nice if you could provide a valid request")
	}

	dealershipReq := &proto.Dealership{
		OwnerId:      createDealershipRequest.OwnerID,
		Name:         createDealershipRequest.Name,
		DisplayName:  createDealershipRequest.DisplayName,
		Address:      createDealershipRequest.Address,
		City:         createDealershipRequest.City,
		State:        createDealershipRequest.State,
		Zip:          createDealershipRequest.Zip,
		Phone:        createDealershipRequest.Phone,
		Email:        createDealershipRequest.Email,
		Website:      createDealershipRequest.Website,
		FacebookUrl:  createDealershipRequest.FacebookUrl,
		TwitterUrl:   createDealershipRequest.TwitterUrl,
		InstagramUrl: createDealershipRequest.InstagramUrl,
		LinkedinUrl:  createDealershipRequest.LinkedinUrl,
		LogoUrl:      createDealershipRequest.LogoUrl,
		CoverUrl:     createDealershipRequest.CoverUrl,
		Description:  createDealershipRequest.Description,
		CreatedAt:    timestamppb.Now(),
		UpdatedAt:    timestamppb.Now(),
	}

	// get dealership from database
	dealership, err := d.dealershipDatabaseRepo.CreateDealership(ctx, dealershipReq)
	if err != nil {
		return nil, err
	}

	// convert dealership to response
	response := &dto.CreateDealershipResponse{
		DealershipID: dealership.DealershipId,
		OwnerID:      dealership.OwnerId,
		Name:         dealership.Name,
		DisplayName:  dealership.DisplayName,
		Address:      dealership.Address,
		City:         dealership.City,
		State:        dealership.State,
		Zip:          dealership.Zip,
		Phone:        dealership.Phone,
		Email:        dealership.Email,
		Website:      dealership.Website,
		FacebookUrl:  dealership.FacebookUrl,
		TwitterUrl:   dealership.TwitterUrl,
		InstagramUrl: dealership.InstagramUrl,
		LinkedinUrl:  dealership.LinkedinUrl,
		LogoUrl:      dealership.LogoUrl,
		CoverUrl:     dealership.CoverUrl,
		Description:  dealership.Description,
	}

	return response, nil
}

func (d dealershipUsecase) DeleteDealership(ctx context.Context, deleteDealershipRequest *dto.DeleteDealershipRequest) (*dto.DeleteLeadershipResponse, error) {
	if deleteDealershipRequest == nil {
		return nil, errors.New("it'd be nice if you could provide a valid request")
	}

	dealershipReq := &proto.Dealership{
		DealershipId: deleteDealershipRequest.DealershipID,
	}

	// get dealership from database
	err := d.dealershipDatabaseRepo.DeleteDealership(ctx, dealershipReq)
	if err != nil {
		return nil, err
	}

	return &dto.DeleteLeadershipResponse{}, nil
}

func NewDealershipUsecase(dealershipDatabaseRepo ports.DealershipDatabaseRepo) ports.DealerShipUseCase {
	return &dealershipUsecase{
		dealershipDatabaseRepo: dealershipDatabaseRepo,
	}
}
