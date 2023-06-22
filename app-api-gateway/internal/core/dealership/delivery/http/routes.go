package http

import (
	"garination.com/gateway/internal/core/dealership/ports"
	"garination.com/gateway/internal/core/middleware"
	"github.com/gin-gonic/gin"
)

func MapDealershipRoutes(group *gin.RouterGroup, h ports.DealerShiphandler, mw middleware.MiddlewareManager) {
	group.GET("/:id", h.GetDealershipByID())
	group.GET("/user/:id", mw.Auth, h.GetDealershipByUserID())
	group.PUT("/:id", mw.Auth, h.UpdateDealership())
	group.POST("/", mw.Auth, h.CreateDealership())
	group.DELETE("/:id", mw.Auth, h.DeleteDealership())
}
