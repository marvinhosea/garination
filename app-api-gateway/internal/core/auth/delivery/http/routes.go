package http

import (
	"garination.com/gateway/internal/core/auth/ports"
	"garination.com/gateway/internal/core/middleware"
	"github.com/gin-gonic/gin"
)

func MapAuthRoutes(group *gin.RouterGroup, authHttpHandler ports.AuthHttpHandler, mw middleware.MiddlewareManager) {
	group.POST("/login", authHttpHandler.InitiateLogin())
	group.POST("/login/verify", authHttpHandler.LoginCallback())
	group.POST("/register", authHttpHandler.InitiateRegister())
	group.POST("/register/verify", authHttpHandler.RegisterCallback())
	group.GET("/user/:user_id", authHttpHandler.GetUserMeta())
	group.PUT("/user/:user_id", authHttpHandler.UpdateUserMeta())
	group.GET("/logout", authHttpHandler.Logout())
}
