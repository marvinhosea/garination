package service

import (
	"context"
	"garination.com/db/internal/core/model"
	"garination.com/db/internal/core/ports/user"
)

type userService struct {
	repo user.UserRepository
}

func (u userService) GetUserMeta(ctx context.Context, userID string) (*model.UserMetum, error) {
	return u.repo.GetUserMeta(ctx, userID)
}

func (u userService) InsertUserMeta(ctx context.Context, arg model.UserMetum) (*model.UserMetum, error) {
	return u.repo.InsertUserMeta(ctx, arg)
}

func (u userService) UpdateUserMeta(ctx context.Context, arg model.UserMetum) (*model.UserMetum, error) {
	return u.repo.UpdateUserMeta(ctx, arg)
}

func NewUserService(repo user.UserRepository) user.UserService {
	return &userService{repo: repo}
}
