package repository

import (
	"garination.com/gateway/internal/core/auth/ports"
	"garination.com/gateway/internal/platform/casdoor"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"golang.org/x/oauth2"
)

type casdoorRepo struct {
	client casdoor.CasdoorClient
}

func (c casdoorRepo) GetOAuthToken(code string, state string) (*oauth2.Token, error) {
	return c.client.GetOAuthToken(code, state)
}

func (c casdoorRepo) ParseJwtToken(token string) (*casdoorsdk.Claims, error) {
	return c.client.ParseJwtToken(token)
}

func (c casdoorRepo) GetUser(username string) (*casdoorsdk.User, error) {
	return c.client.GetUser(username)
}

func (c casdoorRepo) GetUserByEmail(email string) (*casdoorsdk.User, error) {
	return c.client.GetUserByEmail(email)
}

func (c casdoorRepo) GetPermission(name string) (*casdoorsdk.Permission, error) {
	return c.client.GetPermission(name)
}

func (c casdoorRepo) GetSigninUrl(redirectUri string) string {
	return c.client.GetSigninUrl(redirectUri)
}

func (c casdoorRepo) GetSignupUrl(enablePassword bool, redirectUri string) string {
	return c.client.GetSignupUrl(enablePassword, redirectUri)
}

func (c casdoorRepo) GetPaginationUsers(filter casdoor.PaginationUsersFilter) ([]*casdoorsdk.User, int, error) {
	return c.client.GetPaginationUsers(filter)
}

func NewCasdoorRepo(client casdoor.CasdoorClient) ports.AuthCasdoorRepo {
	return &casdoorRepo{client: client}
}
