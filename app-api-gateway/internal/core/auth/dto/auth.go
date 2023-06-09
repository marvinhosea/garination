package dto

// AuthLoginRequest ...
type AuthLoginRequest struct {
	RedirectURL string `json:"redirect_url"` // redirect_url is the url to redirect after login
	Origin      string `json:"origin"`       // origin is the url of the client
	Provider    string `json:"provider"`     // provider is the name of the provider
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

// AuthRegisterResponse ...
type AuthRegisterResponse struct {
	RedirectURL  string `json:"redirect_url"`   // redirect_url is the url to redirect after login
	LoginPageURL string `json:"login_page_url"` // login_page_url is the url of the login page
}

// AuthLogoutRequest ...
type AuthLogoutRequest struct {
}

// AuthLogoutResponse ...
type AuthLogoutResponse struct {
}
