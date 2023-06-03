package casdoor

import (
	"car-marketplace-backend/config"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

type Casdoor interface {
	GetId(name string) string
	GetUser(name string) (*casdoorsdk.User, error)
	GetUrl(action string, queryMap map[string]string) string
	AddPermission(permission *casdoorsdk.Permission) (bool, error)
	GetPermission(name string) (*casdoorsdk.Permission, error)
	AddRole(role *casdoorsdk.Role) (bool, error)
	GetRoles() ([]*casdoorsdk.Role, error)
	GetRole(name string) (*casdoorsdk.Role, error)
	UpdateUser(user *casdoorsdk.User) (bool, error)
}

type casdoor struct{}

func (c casdoor) GetId(name string) string {
	return casdoorsdk.GetId(name)
}

func (c casdoor) GetUser(name string) (*casdoorsdk.User, error) {
	return casdoorsdk.GetUser(name)
}

func (c casdoor) GetUrl(action string, queryMap map[string]string) string {
	return casdoorsdk.GetUrl(action, queryMap)
}

func (c casdoor) AddPermission(permission *casdoorsdk.Permission) (bool, error) {
	return casdoorsdk.AddPermission(permission)
}

func (c casdoor) GetPermission(name string) (*casdoorsdk.Permission, error) {
	return casdoorsdk.GetPermission(name)
}

func (c casdoor) AddRole(role *casdoorsdk.Role) (bool, error) {
	return casdoorsdk.AddRole(role)
}

func (c casdoor) GetRoles() ([]*casdoorsdk.Role, error) {
	return casdoorsdk.GetRoles()
}

func (c casdoor) GetRole(name string) (*casdoorsdk.Role, error) {
	return casdoorsdk.GetRole(name)
}

func (c casdoor) UpdateUser(user *casdoorsdk.User) (bool, error) {
	return casdoorsdk.UpdateUser(user)
}

func NewCasdoor(config config.Casdoor) Casdoor {
	casdoorsdk.InitConfig(config.Endpoint, config.ClientId, config.ClientSecret, config.CertificateX509, config.OrganisationName, config.ApplicationName)
	return &casdoor{}
}
