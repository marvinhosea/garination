package ports

import (
	"github.com/gin-gonic/gin"
)

type AuthHttpHandler interface {
	InitiateLogin() gin.HandlerFunc
	LoginCallback() gin.HandlerFunc
	InitiateRegister() gin.HandlerFunc
	RegisterCallback() gin.HandlerFunc
	UpdateUserMeta() gin.HandlerFunc
	GetUserMeta() gin.HandlerFunc
	Logout() gin.HandlerFunc
}
