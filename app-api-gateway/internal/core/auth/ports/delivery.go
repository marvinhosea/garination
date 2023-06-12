package ports

import (
	"github.com/gin-gonic/gin"
)

type AuthHttpHandler interface {
	InitiateLogin() gin.HandlerFunc
	InitiateRegister() gin.HandlerFunc
	Logout() gin.HandlerFunc
}
