package repository

import (
	"context"
	"garination.com/db/internal/core/model"
	"garination.com/db/internal/core/ports"
	"garination.com/db/internal/platform/postgres"
)

type userRepo struct {
	userPostgresStorage ports.UserPostgresStorage
}

func (u userRepo) GetUserDealership(ctx context.Context, userID string) (*model.Dealership, error) {
	dealership, err := u.userPostgresStorage.GetUserDealership(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &model.Dealership{
		DealershipID: dealership.DealershipID,
		OwnerID:      dealership.OwnerID,
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
		CreatedAt:    dealership.CreatedAt,
		UpdatedAt:    dealership.UpdatedAt,
	}, nil
}

func (u userRepo) GetUserMeta(ctx context.Context, userID string) (*model.UserMetum, error) {
	userMeta, err := u.userPostgresStorage.GetUserMeta(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &model.UserMetum{
		UserMetaID:   userMeta.UserMetaID,
		UserID:       userMeta.UserID,
		FacebookUrl:  userMeta.FacebookUrl,
		TwitterUrl:   userMeta.TwitterUrl,
		InstagramUrl: userMeta.InstagramUrl,
		LinkedinUrl:  userMeta.LinkedinUrl,
		WebsiteUrl:   userMeta.WebsiteUrl,
		DealershipID: userMeta.DealershipID,
	}, nil
}

func (u userRepo) InsertDealership(ctx context.Context, arg model.Dealership) (*model.Dealership, error) {
	request := postgres.InsertDealershipParams{
		DealershipID: "",
		OwnerID:      arg.OwnerID,
		Name:         arg.Name,
		DisplayName:  arg.DisplayName,
		Address:      arg.Address,
		City:         arg.City,
		State:        arg.State,
		Zip:          arg.Zip,
		Phone:        arg.Phone,
		Email:        arg.Email,
		Website:      arg.Website,
		FacebookUrl:  arg.FacebookUrl,
		TwitterUrl:   arg.TwitterUrl,
		InstagramUrl: arg.InstagramUrl,
		LinkedinUrl:  arg.LinkedinUrl,
		LogoUrl:      arg.LogoUrl,
		CoverUrl:     arg.CoverUrl,
		Description:  arg.Description,
		CreatedAt:    arg.CreatedAt,
		UpdatedAt:    arg.UpdatedAt,
	}
	dealership, err := u.userPostgresStorage.InsertDealership(ctx, request)
	if err != nil {
		return nil, err
	}

	return &model.Dealership{
		DealershipID: dealership.DealershipID,
	}, nil
}

func (u userRepo) InsertUserMeta(ctx context.Context, arg model.UserMetum) (*model.UserMetum, error) {
	request := postgres.InsertUserMetaParams{
		UserMetaID:   "",
		UserID:       arg.UserID,
		FacebookUrl:  arg.FacebookUrl,
		TwitterUrl:   arg.TwitterUrl,
		InstagramUrl: arg.InstagramUrl,
		LinkedinUrl:  arg.LinkedinUrl,
		WebsiteUrl:   arg.WebsiteUrl,
		DealershipID: arg.DealershipID,
	}
	userMeta, err := u.userPostgresStorage.InsertUserMeta(ctx, request)
	if err != nil {
		return nil, err
	}

	return &model.UserMetum{
		UserMetaID:   userMeta.UserMetaID,
		UserID:       userMeta.UserID,
		FacebookUrl:  userMeta.FacebookUrl,
		TwitterUrl:   userMeta.TwitterUrl,
		InstagramUrl: userMeta.InstagramUrl,
		LinkedinUrl:  userMeta.LinkedinUrl,
		WebsiteUrl:   userMeta.WebsiteUrl,
		DealershipID: userMeta.DealershipID,
	}, nil
}

func (u userRepo) UpdateUserMeta(ctx context.Context, arg model.UserMetum) (*model.UserMetum, error) {
	request := postgres.UpdateUserMetaParams{
		UserMetaID:   arg.UserMetaID,
		FacebookUrl:  arg.FacebookUrl,
		TwitterUrl:   arg.TwitterUrl,
		InstagramUrl: arg.InstagramUrl,
		LinkedinUrl:  arg.LinkedinUrl,
		WebsiteUrl:   arg.WebsiteUrl,
	}
	userMeta, err := u.userPostgresStorage.UpdateUserMeta(ctx, request)
	if err != nil {
		return nil, err
	}

	return &model.UserMetum{
		UserMetaID:   userMeta.UserMetaID,
		UserID:       userMeta.UserID,
		FacebookUrl:  userMeta.FacebookUrl,
		TwitterUrl:   userMeta.TwitterUrl,
		InstagramUrl: userMeta.InstagramUrl,
		LinkedinUrl:  userMeta.LinkedinUrl,
		WebsiteUrl:   userMeta.WebsiteUrl,
		DealershipID: userMeta.DealershipID,
	}, nil
}

func NewUserRepo(userPostgresStorage ports.UserPostgresStorage) ports.UserRepository {
	return &userRepo{
		userPostgresStorage: userPostgresStorage,
	}
}
