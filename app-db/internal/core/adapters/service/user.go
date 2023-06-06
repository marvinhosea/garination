package service

import (
	"context"
	"garination.com/db/internal/core/model"
	"garination.com/db/internal/core/ports"
)

type userService struct {
	repo ports.UserRepository
}

func (u userService) GetUserDealership(ctx context.Context, userID string) (*model.Dealership, error) {
	return u.repo.GetUserDealership(ctx, userID)
}

func (u userService) GetUserMeta(ctx context.Context, userID string) (*model.UserMetum, error) {
	return u.repo.GetUserMeta(ctx, userID)
}

func (u userService) InsertDealership(ctx context.Context, arg model.Dealership) (*model.Dealership, error) {
	return u.repo.InsertDealership(ctx, arg)
}

func (u userService) InsertUserMeta(ctx context.Context, arg model.UserMetum) (*model.UserMetum, error) {
	return u.repo.InsertUserMeta(ctx, arg)
}

func (u userService) UpdateUserMeta(ctx context.Context, arg model.UserMetum) (*model.UserMetum, error) {
	return u.repo.UpdateUserMeta(ctx, arg)
}

func NewUserService(repo ports.UserRepository) ports.UserService {
	return &userService{repo: repo}
}
