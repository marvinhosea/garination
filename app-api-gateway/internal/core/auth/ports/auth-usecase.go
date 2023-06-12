package ports

import (
	"context"
	"garination.com/gateway/internal/core/auth/dto"
)

type AuthUsecase interface {
	InitiateLogin(ctx context.Context, req *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error)
	InitiateRegister(ctx context.Context, req *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error)
	Logout(ctx context.Context, req *dto.AuthLogoutRequest) (*dto.AuthLogoutResponse, error)
}
