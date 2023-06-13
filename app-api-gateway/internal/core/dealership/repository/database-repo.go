package repository

import (
	"context"
	"garination.com/gateway/internal/core/dealership/ports"
	"garination.com/gateway/internal/platform/grpc/app-db/proto"
)

type dealershipDatabaseRepo struct {
	client proto.DatabaseServiceClient
}

func (d dealershipDatabaseRepo) GetDealershipByID(ctx context.Context, id string) (*proto.Dealership, error) {
	req := &proto.GetDealershipByIDRequest{
		DealershipId: id,
	}

	res, err := d.client.GetDealershipByID(ctx, req)
	if err != nil {
		return nil, err
	}

	return &proto.Dealership{
		DealershipId: res.Dealership.DealershipId,
		OwnerId:      res.Dealership.OwnerId,
		Name:         res.Dealership.Name,
		DisplayName:  res.Dealership.DisplayName,
		Address:      res.Dealership.Address,
		City:         res.Dealership.City,
		State:        res.Dealership.State,
		Zip:          res.Dealership.Zip,
		Phone:        res.Dealership.Phone,
		Email:        res.Dealership.Email,
		Website:      res.Dealership.Website,
		FacebookUrl:  res.Dealership.FacebookUrl,
		TwitterUrl:   res.Dealership.TwitterUrl,
		InstagramUrl: res.Dealership.InstagramUrl,
		LinkedinUrl:  res.Dealership.LinkedinUrl,
		LogoUrl:      res.Dealership.LogoUrl,
		CoverUrl:     res.Dealership.CoverUrl,
		Description:  res.Dealership.Description,
		CreatedAt:    res.Dealership.CreatedAt,
		UpdatedAt:    res.Dealership.UpdatedAt,
	}, nil
}

func (d dealershipDatabaseRepo) GetDealershipByUserID(ctx context.Context, userID string) (*proto.Dealership, error) {
	req := &proto.GetUserDealershipRequest{
		UserId: userID,
	}

	res, err := d.client.GetDealershipByUserId(ctx, req)
	if err != nil {
		return nil, err
	}

	return &proto.Dealership{
		DealershipId: res.Dealership.DealershipId,
		OwnerId:      res.Dealership.OwnerId,
		Name:         res.Dealership.Name,
		DisplayName:  res.Dealership.DisplayName,
		Address:      res.Dealership.Address,
		City:         res.Dealership.City,
		State:        res.Dealership.State,
		Zip:          res.Dealership.Zip,
		Phone:        res.Dealership.Phone,
		Email:        res.Dealership.Email,
		Website:      res.Dealership.Website,
		FacebookUrl:  res.Dealership.FacebookUrl,
		TwitterUrl:   res.Dealership.TwitterUrl,
		InstagramUrl: res.Dealership.InstagramUrl,
		LinkedinUrl:  res.Dealership.LinkedinUrl,
		LogoUrl:      res.Dealership.LogoUrl,
		CoverUrl:     res.Dealership.CoverUrl,
		Description:  res.Dealership.Description,
		CreatedAt:    res.Dealership.CreatedAt,
		UpdatedAt:    res.Dealership.UpdatedAt,
	}, nil
}

func (d dealershipDatabaseRepo) UpdateDealership(ctx context.Context, dealership *proto.Dealership) (*proto.Dealership, error) {
	req := &proto.UpdateDealershipRequest{
		Dealership: dealership,
	}

	res, err := d.client.UpdateDealership(ctx, req)
	if err != nil {
		return nil, err
	}

	return &proto.Dealership{
		DealershipId: res.Dealership.DealershipId,
		OwnerId:      res.Dealership.OwnerId,
		Name:         res.Dealership.Name,
		DisplayName:  res.Dealership.DisplayName,
		Address:      res.Dealership.Address,
		City:         res.Dealership.City,
		State:        res.Dealership.State,
		Zip:          res.Dealership.Zip,
		Phone:        res.Dealership.Phone,
		Email:        res.Dealership.Email,
		Website:      res.Dealership.Website,
		FacebookUrl:  res.Dealership.FacebookUrl,
		TwitterUrl:   res.Dealership.TwitterUrl,
		InstagramUrl: res.Dealership.InstagramUrl,
		LinkedinUrl:  res.Dealership.LinkedinUrl,
		LogoUrl:      res.Dealership.LogoUrl,
		CoverUrl:     res.Dealership.CoverUrl,
		Description:  res.Dealership.Description,
		CreatedAt:    res.Dealership.CreatedAt,
		UpdatedAt:    res.Dealership.UpdatedAt,
	}, nil
}

func (d dealershipDatabaseRepo) CreateDealership(ctx context.Context, dealership *proto.Dealership) (*proto.Dealership, error) {
	req := &proto.InsertDealershipRequest{
		Dealership: dealership,
	}

	res, err := d.client.InsertDealership(ctx, req)
	if err != nil {
		return nil, err
	}

	return &proto.Dealership{
		DealershipId: res.Dealership.DealershipId,
		OwnerId:      res.Dealership.OwnerId,
		Name:         res.Dealership.Name,
		DisplayName:  res.Dealership.DisplayName,
		Address:      res.Dealership.Address,
		City:         res.Dealership.City,
		State:        res.Dealership.State,
		Zip:          res.Dealership.Zip,
		Phone:        res.Dealership.Phone,
		Email:        res.Dealership.Email,
		Website:      res.Dealership.Website,
		FacebookUrl:  res.Dealership.FacebookUrl,
		TwitterUrl:   res.Dealership.TwitterUrl,
		InstagramUrl: res.Dealership.InstagramUrl,
		LinkedinUrl:  res.Dealership.LinkedinUrl,
		LogoUrl:      res.Dealership.LogoUrl,
		CoverUrl:     res.Dealership.CoverUrl,
		Description:  res.Dealership.Description,
		CreatedAt:    res.Dealership.CreatedAt,
		UpdatedAt:    res.Dealership.UpdatedAt,
	}, nil
}

func (d dealershipDatabaseRepo) DeleteDealership(ctx context.Context, dealershipId *proto.Dealership) error {
	req := &proto.DeleteDealershipRequest{
		DealershipId: dealershipId.DealershipId,
	}

	_, err := d.client.DeleteDealership(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func NewDealershipDatabaseRepo(client proto.DatabaseServiceClient) ports.DealershipDatabaseRepo {
	return &dealershipDatabaseRepo{client: client}
}
