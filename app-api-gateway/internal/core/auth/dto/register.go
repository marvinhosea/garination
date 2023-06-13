package dto

import (
	"errors"
	"garination.com/gateway/internal/core/auth/model"
	"regexp"
	"time"
)

// AuthRegisterRequest ...
type AuthRegisterRequest struct {
	RedirectURL string `json:"redirect_url"` // redirect_url is the url to redirect after login
	Origin      string `json:"origin"`       // origin is the url of the client
	Provider    string `json:"provider"`     // provider is the name of the provider
}

// AuthRegisterResponse ...
type AuthRegisterResponse struct {
	RedirectURL  string `json:"redirect_url"`   // redirect_url is the url to redirect after login
	LoginPageURL string `json:"login_page_url"` // login_page_url is the url of the login page
}

// AuthRegisterCallbackRequest ...
type AuthRegisterCallbackRequest struct {
	Code  string `json:"code"`  // code is the code returned by the provider
	State string `json:"state"` // state is the state returned by the provider
}

// AuthRegisterCallbackResponse ...
type AuthRegisterCallbackResponse struct {
	User               *model.User     `json:"user_info"`
	UserAdditionalInfo *model.UserMeta `json:"user_additional_info"`
	AccessToken        string          `json:"access_token"`
	TokenType          string          `json:"token_type,omitempty"`
	RefreshToken       string          `json:"refresh_token,omitempty"`
	Expiry             time.Time       `json:"expiry,omitempty"`
}

func (a AuthRegisterRequest) Validate() error {
	var issues error
	if a.RedirectURL == "" {
		issues = errors.Join(issues, errors.New("redirect_url is required"))
	}

	if a.Origin == "" {
		issues = errors.Join(issues, errors.New("origin is required"))
	}

	if a.Provider == "" {
		issues = errors.Join(issues, errors.New("provider is required"))
	}

	// validate provider (facebook, google, apple only)
	if a.Provider != "facebook" && a.Provider != "google" && a.Provider != "apple" {
		issues = errors.Join(issues, errors.New("provider is not supported"))
	}

	// validate redirect_url is a valid url
	re := regexp.MustCompile(urlRegex)
	if !re.MatchString(a.RedirectURL) {
		issues = errors.Join(issues, errors.New("redirect_url is not a valid url"))
	}

	// validate origin is a valid url
	if !re.MatchString(a.Origin) {
		issues = errors.Join(issues, errors.New("origin is not a valid url"))
	}

	return issues
}

func (a AuthRegisterCallbackRequest) Validate() error {
	var issues error
	if a.Code == "" {
		issues = errors.Join(issues, errors.New("code is required"))
	}

	if a.State == "" {
		issues = errors.Join(issues, errors.New("state is required"))
	}

	return issues
}
