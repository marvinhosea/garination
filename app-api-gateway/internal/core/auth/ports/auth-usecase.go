package ports

import (
	"context"
	"garination.com/gateway/internal/core/auth/dto"
)

type AuthUsecase interface {
	InitiateLogin(ctx context.Context, req *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error)
	InitiateRegister(ctx context.Context, req *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error)
	LoginCallback(ctx context.Context, req *dto.AuthLoginCallbackRequest) (*dto.AuthLoginCallbackResponse, error)
	GetUserMeta(ctx context.Context, req *dto.AuthGetUserMetaRequest) (*dto.AuthGetUserMetaResponse, error)
	ChangeUserDealership(ctx context.Context, req *dto.AuthChangeUserDealershipRequest) (*dto.AuthChangeUserDealershipResponse, error)
	UpdateUserMeta(ctx context.Context, req *dto.AuthUpdateUserMetaRequest) (*dto.AuthUpdateUserMetaResponse, error)
	RegisterCallback(ctx context.Context, req *dto.AuthRegisterCallbackRequest) (*dto.AuthRegisterCallbackResponse, error)
	RefreshToken(ctx context.Context, req *dto.AuthRefreshTokenRequest) (*dto.AuthRefreshTokenResponse, error)
	Logout(ctx context.Context, req *dto.AuthLogoutRequest) (*dto.AuthLogoutResponse, error)
}
