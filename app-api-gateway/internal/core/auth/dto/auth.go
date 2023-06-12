package dto

import (
	"errors"
	"regexp"
)

const urlRegex = `^(http|https):\/\/[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`

// AuthLoginRequest ...
type AuthLoginRequest struct {
	RedirectURL string `json:"redirect_url"` // redirect_url is the url to redirect after login
	Origin      string `json:"origin"`       // origin is the url of the client
	Provider    string `json:"provider"`     // provider is the name of the provider
}

func (a AuthLoginRequest) Validate() error {
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

// AuthLoginResponse ...
type AuthLoginResponse struct {
	RedirectURL  string `json:"redirect_url"`   // redirect_url is the url to redirect after login
	LoginPageURL string `json:"login_page_url"` // login_page_url is the url of the login page
}

// AuthRegisterRequest ...
type AuthRegisterRequest struct {
	RedirectURL string `json:"redirect_url"` // redirect_url is the url to redirect after login
	Origin      string `json:"origin"`       // origin is the url of the client
	Provider    string `json:"provider"`     // provider is the name of the provider
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

// AuthRegisterResponse ...
type AuthRegisterResponse struct {
	RedirectURL  string `json:"redirect_url"`   // redirect_url is the url to redirect after login
	LoginPageURL string `json:"login_page_url"` // login_page_url is the url of the login page
}

// AuthLogoutRequest ...
type AuthLogoutRequest struct {
}

func (r AuthLogoutRequest) Validate() error {
	return nil
}

// AuthLogoutResponse ...
type AuthLogoutResponse struct {
}
