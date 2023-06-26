package service

import (
	"context"
	"garination.com/db/internal/core/ports/user"
	"garination.com/db/internal/platform/postgres"
	"github.com/google/uuid"
)

type userService struct {
	repo user.UserRepository
}

func (u userService) UpdateUserMeta(ctx context.Context, arg postgres.UpdateUserMetaParams) (*postgres.UserMetum, error) {
	// get user meta
	userMeta, err := u.repo.GetUserMeta(ctx, arg.UserID)
	if err != nil {
		return nil, err
	}

	arg.UserMetaID = userMeta.UserMetaID
	return u.repo.UpdateUserMeta(ctx, arg)
}

func (u userService) UpdateUserDealership(ctx context.Context, args postgres.UpdateUserDealershipParams) (*postgres.UserMetum, error) {
	return u.repo.UpdateUserDealership(ctx, args)
}

func (u userService) GetUserMeta(ctx context.Context, userID string) (*postgres.UserMetum, error) {
	return u.repo.GetUserMeta(ctx, userID)
}

func (u userService) InsertUserMeta(ctx context.Context, arg postgres.UserMetum) (*postgres.UserMetum, error) {
	if arg.UserMetaID == "" {
		arg.UserMetaID = uuid.NewString()
	}

	if len(arg.DealershipID.String) == 0 {
		insertUserMetaWithoutDealershipParams := postgres.InsertUserMetaWithoutDealershipParams{
			UserMetaID:   arg.UserMetaID,
			UserID:       arg.UserID,
			FacebookUrl:  arg.FacebookUrl,
			TwitterUrl:   arg.TwitterUrl,
			InstagramUrl: arg.InstagramUrl,
			LinkedinUrl:  arg.LinkedinUrl,
			WebsiteUrl:   arg.WebsiteUrl,
		}
		return u.repo.InsertUserMetaWithoutDealershipParams(ctx, insertUserMetaWithoutDealershipParams)
	} else {
		insertUserMetaParams := postgres.InsertUserMetaParams{
			UserMetaID:   arg.UserMetaID,
			UserID:       arg.UserID,
			DealershipID: arg.DealershipID,
			FacebookUrl:  arg.FacebookUrl,
			TwitterUrl:   arg.TwitterUrl,
			InstagramUrl: arg.InstagramUrl,
			LinkedinUrl:  arg.LinkedinUrl,
			WebsiteUrl:   arg.WebsiteUrl,
		}
		return u.repo.InsertUserMeta(ctx, insertUserMetaParams)
	}
}
func NewUserService(repo user.UserRepository) user.UserService {
	return &userService{repo: repo}
}
