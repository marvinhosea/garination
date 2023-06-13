package dto

import (
	"errors"
	"time"
)

type AuthRefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (r AuthRefreshTokenRequest) Validate() error {
	if r.RefreshToken == "" {
		return errors.New("refresh_token is required")
	}
	return nil
}

type AuthRefreshTokenResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	TokenType    string    `json:"token_type"`
	Expiry       time.Time `json:"expiry"`
}
