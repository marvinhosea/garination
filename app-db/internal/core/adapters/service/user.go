package service

import (
	"context"
	"garination.com/db/internal/core/model"
	"garination.com/db/internal/core/ports/user"
	"github.com/google/uuid"
)

type userService struct {
	repo user.UserRepository
}

func (u userService) GetUserMeta(ctx context.Context, userID string) (*model.UserMetum, error) {
	return u.repo.GetUserMeta(ctx, userID)
}

func (u userService) InsertUserMeta(ctx context.Context, arg model.UserMetum) (*model.UserMetum, error) {
	if arg.UserMetaID == "" {
		arg.UserMetaID = uuid.NewString()
	}
	return u.repo.InsertUserMeta(ctx, arg)
}

func (u userService) UpdateUserMeta(ctx context.Context, arg model.UserMetum) (*model.UserMetum, error) {
	// get user meta
	userMeta, err := u.repo.GetUserMeta(ctx, arg.UserID)
	if err != nil {
		return nil, err
	}

	arg.UserMetaID = userMeta.UserMetaID
	return u.repo.UpdateUserMeta(ctx, arg)
}

func NewUserService(repo user.UserRepository) user.UserService {
	return &userService{repo: repo}
}
