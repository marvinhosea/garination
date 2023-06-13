package usecase

import (
	"context"
	"errors"
	"garination.com/gateway/internal/core/auth/dto"
	"garination.com/gateway/internal/core/auth/model"
	"garination.com/gateway/internal/core/auth/ports"
)

type authUsecase struct {
	authRedisRepo ports.AuthRedisRepo
	casdoorRepo   ports.AuthCasdoorRepo
	authDBRepo    ports.AuthDbRepo
}

func (a authUsecase) RefreshToken(ctx context.Context, req *dto.AuthRefreshTokenRequest) (*dto.AuthRefreshTokenResponse, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	refreshToken, err := a.casdoorRepo.RefreshOAuthToken(req.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &dto.AuthRefreshTokenResponse{
		AccessToken:  refreshToken.AccessToken,
		RefreshToken: refreshToken.RefreshToken,
		TokenType:    "Bearer",
		Expiry:       refreshToken.Expiry,
	}, nil
}

func (a authUsecase) GetUserMeta(ctx context.Context, req *dto.AuthGetUserMetaRequest) (*dto.AuthGetUserMetaResponse, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	userMeta, err := a.authDBRepo.GetUserMeta(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	return &dto.AuthGetUserMetaResponse{
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

func (a authUsecase) UpdateUserMeta(ctx context.Context, req *dto.AuthUpdateUserMetaRequest) (*dto.AuthUpdateUserMetaResponse, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	userMeta := model.UserMeta{
		UserID:       req.UserID,
		FacebookUrl:  req.FacebookUrl,
		TwitterUrl:   req.TwitterUrl,
		InstagramUrl: req.InstagramUrl,
		LinkedinUrl:  req.LinkedinUrl,
		WebsiteUrl:   req.WebsiteUrl,
		DealershipID: req.DealershipID,
	}

	userMetaRes, err := a.authDBRepo.UpdateUserMeta(ctx, userMeta)
	if err != nil {
		return nil, err
	}

	return &dto.AuthUpdateUserMetaResponse{
		UserMetaID:   userMetaRes.UserMetaID,
		UserID:       userMetaRes.UserID,
		FacebookUrl:  userMetaRes.FacebookUrl,
		TwitterUrl:   userMetaRes.TwitterUrl,
		InstagramUrl: userMetaRes.InstagramUrl,
		LinkedinUrl:  userMetaRes.LinkedinUrl,
		WebsiteUrl:   userMetaRes.WebsiteUrl,
		DealershipID: userMetaRes.DealershipID,
	}, nil
}

func (a authUsecase) LoginCallback(ctx context.Context, req *dto.AuthLoginCallbackRequest) (*dto.AuthLoginCallbackResponse, error) {
	token, err := a.casdoorRepo.GetOAuthToken(req.Code, req.State)
	if err != nil {
		return nil, err
	}

	// verify token
	claims, err := a.casdoorRepo.ParseJwtToken(token.AccessToken)

	if err != nil {
		return nil, err
	}

	// get user meta
	userMeta, _ := a.authDBRepo.GetUserMeta(ctx, claims.GetId())
	claims.AccessToken = token.AccessToken

	return &dto.AuthLoginCallbackResponse{
		User: &model.User{
			UserID:      claims.Id,
			Email:       claims.Email,
			PhoneNumber: claims.Phone,
			Address:     claims.Address,
		},
		UserAdditionalInfo: userMeta,
		AccessToken:        claims.AccessToken,
		TokenType:          token.TokenType,
		RefreshToken:       token.RefreshToken,
		Expiry:             token.Expiry,
	}, nil

}

func (a authUsecase) RegisterCallback(ctx context.Context, req *dto.AuthRegisterCallbackRequest) (*dto.AuthRegisterCallbackResponse, error) {
	token, err := a.casdoorRepo.GetOAuthToken(req.Code, req.State)
	if err != nil {
		return nil, err
	}

	// verify token
	claims, err := a.casdoorRepo.ParseJwtToken(token.AccessToken)

	if err != nil {
		return nil, err
	}

	// get user meta
	userMeta, _ := a.authDBRepo.GetUserMeta(ctx, claims.GetId())

	claims.AccessToken = token.AccessToken

	return &dto.AuthRegisterCallbackResponse{
		User: &model.User{
			UserID:      claims.Id,
			Email:       claims.Email,
			PhoneNumber: claims.Phone,
			Address:     claims.Address,
		},
		UserAdditionalInfo: userMeta,
		AccessToken:        claims.AccessToken,
		TokenType:          token.TokenType,
		RefreshToken:       token.RefreshToken,
		Expiry:             token.Expiry,
	}, nil
}

func (a authUsecase) InitiateLogin(_ context.Context, req *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error) {
	res := a.casdoorRepo.GetSigninUrl(req.RedirectURL)
	return &dto.AuthLoginResponse{
		LoginPageURL: res,
		RedirectURL:  req.RedirectURL,
	}, nil
}

func (a authUsecase) InitiateRegister(_ context.Context, req *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error) {
	res := a.casdoorRepo.GetSignupUrl(true, req.RedirectURL)
	return &dto.AuthRegisterResponse{
		RegisterPageUrl: res,
		RedirectURL:     req.RedirectURL,
	}, nil
}

func (a authUsecase) Logout(ctx context.Context, req *dto.AuthLogoutRequest) (*dto.AuthLogoutResponse, error) {
	res, err := a.casdoorRepo.DeleteToken(req.Name)
	if err != nil {
		return nil, err
	}

	if !res {
		return nil, errors.New("we tried to logout but failed")
	}

	return &dto.AuthLogoutResponse{
		State: "success",
	}, nil
}

func NewAuthUsecase(authRedisRepo ports.AuthRedisRepo, casdoorRepo ports.AuthCasdoorRepo, authDBRepo ports.AuthDbRepo) ports.AuthUsecase {
	return &authUsecase{
		authRedisRepo: authRedisRepo,
		casdoorRepo:   casdoorRepo,
		authDBRepo:    authDBRepo,
	}
}
