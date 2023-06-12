package http

import (
	"garination.com/gateway/internal/core/auth/ports"
	"garination.com/gateway/internal/core/middleware"
	"github.com/gin-gonic/gin"
)

func MapAuthRoutes(group *gin.RouterGroup, authHttpHandler ports.AuthHttpHandler, mw middleware.MiddlewareManager) {
	group.POST("/login", authHttpHandler.InitiateLogin())
	group.POST("/register", authHttpHandler.InitiateRegister())
	group.GET("/logout", authHttpHandler.Logout())
}
