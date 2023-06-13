package dto

import (
	"errors"
	"garination.com/gateway/internal/core/auth/model"
	"regexp"
	"time"
)

// AuthLoginRequest ...
type AuthLoginRequest struct {
	RedirectURL string `json:"redirect_url"` // redirect_url is the url to redirect after login
}

// AuthLoginResponse ...
type AuthLoginResponse struct {
	RedirectURL  string `json:"redirect_url"`   // redirect_url is the url to redirect after login
	LoginPageURL string `json:"login_page_url"` // login_page_url is the url of the login page
}

// AuthLoginCallbackRequest ...
type AuthLoginCallbackRequest struct {
	Code  string `json:"code"`  // code is the code returned from the provider
	State string `json:"state"` // state is the state returned from the provider
}

// AuthLoginCallbackResponse ...
type AuthLoginCallbackResponse struct {
	User               *model.User     `json:"user_info"`
	UserAdditionalInfo *model.UserMeta `json:"user_additional_info"`
	AccessToken        string          `json:"access_token"`
	TokenType          string          `json:"token_type,omitempty"`
	RefreshToken       string          `json:"refresh_token,omitempty"`
	Expiry             time.Time       `json:"expiry,omitempty"`
}

func (a AuthLoginRequest) Validate() error {
	var issues error
	if a.RedirectURL == "" {
		issues = errors.Join(issues, errors.New("redirect_url is required"))
	}
	// validate redirect_url is a valid url
	re := regexp.MustCompile(urlRegex)
	if !re.MatchString(a.RedirectURL) {
		issues = errors.Join(issues, errors.New("redirect_url is not a valid url"))
	}

	return issues
}

func (a AuthLoginCallbackRequest) Validate() error {
	var issues error
	if a.Code == "" {
		issues = errors.Join(issues, errors.New("code is required"))
	}

	if a.State == "" {
		issues = errors.Join(issues, errors.New("state is required"))
	}

	return issues
}
