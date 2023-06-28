package casdoor

import (
	"garination.com/gateway/config"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"golang.org/x/oauth2"
	"os"
	"strconv"
)

type PaginationUsersFilter struct {
	Owner    string
	Page     int
	PageSize int
}

type CasdoorClient interface {
	GetUser(username string) (*casdoorsdk.User, error)
	GetUserByEmail(email string) (*casdoorsdk.User, error)
	GetPermission(name string) (*casdoorsdk.Permission, error)
	GetSigninUrl(redirectUri string) string
	GetSignupUrl(enablePassword bool, redirectUri string) string
	GetPaginationUsers(filter PaginationUsersFilter) ([]*casdoorsdk.User, int, error)
	GetOAuthToken(code string, state string) (*oauth2.Token, error)
	ParseJwtToken(token string) (*casdoorsdk.Claims, error)
	RefreshOAuthToken(refreshToken string) (*oauth2.Token, error)
	DeleteToken(name string) (bool, error)
}

type casdoorClientImpl struct {
}

func (c casdoorClientImpl) RefreshOAuthToken(refreshToken string) (*oauth2.Token, error) {
	return casdoorsdk.RefreshOAuthToken(refreshToken)
}

func (c casdoorClientImpl) DeleteToken(name string) (bool, error) {
	return casdoorsdk.DeleteToken(name)
}

func (c casdoorClientImpl) GetOAuthToken(code string, state string) (*oauth2.Token, error) {
	return casdoorsdk.GetOAuthToken(code, state)
}

func (c casdoorClientImpl) ParseJwtToken(token string) (*casdoorsdk.Claims, error) {
	return casdoorsdk.ParseJwtToken(token)
}

func (c casdoorClientImpl) GetSigninUrl(redirectUri string) string {
	return casdoorsdk.GetSigninUrl(redirectUri)
}

func (c casdoorClientImpl) GetSignupUrl(enablePassword bool, redirectUri string) string {
	return casdoorsdk.GetSignupUrl(enablePassword, redirectUri)
}

func (c casdoorClientImpl) GetPaginationUsers(filter PaginationUsersFilter) ([]*casdoorsdk.User, int, error) {
	filterMap := map[string]string{}
	filterMap["p"] = strconv.Itoa(filter.Page)

	if filter.Owner != "" {
		filterMap["owner"] = filter.Owner
	}

	if filter.PageSize != 0 {
		filterMap["pageSize"] = strconv.Itoa(filter.PageSize)
	} else {
		filterMap["pageSize"] = strconv.Itoa(25)
	}

	return casdoorsdk.GetPaginationUsers(filter.Page, filter.PageSize, filterMap)
}

func (c casdoorClientImpl) GetUser(username string) (*casdoorsdk.User, error) {
	return casdoorsdk.GetUser(username)
}

func (c casdoorClientImpl) GetUserByEmail(email string) (*casdoorsdk.User, error) {
	return casdoorsdk.GetUserByEmail(email)
}

func (c casdoorClientImpl) GetPermission(name string) (*casdoorsdk.Permission, error) {
	return casdoorsdk.GetPermission(name)
}

func NewCasdoorClient(cfg config.Casdoor) (CasdoorClient, error) {
	// open certificate
	file, err := os.Open(cfg.CertificateX509)
	if err != nil {
		return nil, err
	}

	// close file
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	// read certificate
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// read certificate
	certificate := make([]byte, stat.Size())
	_, err = file.Read(certificate)
	if err != nil {
		return nil, err
	}
	casdoorsdk.InitConfig(cfg.Endpoint, cfg.ClientId, cfg.ClientSecret, string(certificate), cfg.OrganisationName, cfg.ApplicationName)
	return &casdoorClientImpl{}, nil
}
