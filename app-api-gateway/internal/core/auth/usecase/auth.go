package usecase

import (
	"context"
	"garination.com/gateway/internal/core/auth/dto"
	"garination.com/gateway/internal/core/auth/ports"
)

type authUsecase struct {
	authRedisRepo ports.AuthRedisRepo
	casdoorRepo   ports.AuthCasdoorRepo
	authDBRepo    ports.AuthDbRepo
}

func (a authUsecase) InitiateLogin(ctx context.Context, req *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error) {
	res := a.casdoorRepo.GetSigninUrl(req.RedirectURL)
	return &dto.AuthLoginResponse{
		RedirectURL: res,
	}, nil
}

func (a authUsecase) InitiateRegister(ctx context.Context, req *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error) {
	res := a.casdoorRepo.GetSignupUrl(true, req.RedirectURL)
	return &dto.AuthRegisterResponse{
		RedirectURL: res,
	}, nil
}

func (a authUsecase) Logout(ctx context.Context, req *dto.AuthLogoutRequest) (*dto.AuthLogoutResponse, error) {
	panic("implement me")
}

func NewAuthUsecase(authRedisRepo ports.AuthRedisRepo, casdoorRepo ports.AuthCasdoorRepo, authDBRepo ports.AuthDbRepo) ports.AuthUsecase {
	return &authUsecase{
		authRedisRepo: authRedisRepo,
		casdoorRepo:   casdoorRepo,
		authDBRepo:    authDBRepo,
	}
}
