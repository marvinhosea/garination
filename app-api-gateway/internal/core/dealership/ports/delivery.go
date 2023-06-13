package ports

import "github.com/gin-gonic/gin"

type DealerShiphandler interface {
	GetDealershipByID() gin.HandlerFunc
	GetDealershipByUserID() gin.HandlerFunc
	UpdateDealership() gin.HandlerFunc
	CreateDealership() gin.HandlerFunc
	DeleteDealership() gin.HandlerFunc
}
