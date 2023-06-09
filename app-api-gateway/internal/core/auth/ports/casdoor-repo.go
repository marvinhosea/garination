package ports

import (
	"garination.com/gateway/internal/platform/casdoor"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

type CasdoorRepo interface {
	GetUser(username string) (*casdoorsdk.User, error)
	GetUserByEmail(email string) (*casdoorsdk.User, error)
	GetPermission(name string) (*casdoorsdk.Permission, error)
	GetSigninUrl(redirectUri string) string
	GetSignupUrl(enablePassword bool, redirectUri string) string
	GetPaginationUsers(filter casdoor.PaginationUsersFilter) ([]*casdoorsdk.User, int, error)
}
