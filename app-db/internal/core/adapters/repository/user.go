package repository

import (
	"context"
	"garination.com/db/internal/core/model"
	"garination.com/db/internal/core/ports/user"
	"garination.com/db/internal/platform/postgres"
	"github.com/jackc/pgx/v5/pgtype"
)

type userRepo struct {
	userPostgresStorage user.PostgresStorage
}

func (u userRepo) UpdateUserDealership(ctx context.Context, userId, dealershipId string) (*model.UserMetum, error) {
	request := postgres.UpdateUserDealershipParams{
		UserID: userId,
		DealershipID: pgtype.Text{
			String: dealershipId,
			Valid:  true,
		},
	}
	userMeta, err := u.userPostgresStorage.UpdateUserDealership(ctx, request)
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
		UserID:       arg.UserID,
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

func NewUserRepo(userPostgresStorage user.PostgresStorage) user.UserRepository {
	return &userRepo{
		userPostgresStorage: userPostgresStorage,
	}
}
