package http

import (
	"garination.com/gateway/internal/core/middleware"
	"garination.com/gateway/internal/core/spares/ports"
	"github.com/gin-gonic/gin"
)

func MapSpareRoutes(group *gin.RouterGroup, spareHttHandler ports.Handler, mw middleware.MiddlewareManager) {
	group.GET("/", spareHttHandler.ListSparePartsPaged())
	group.POST("/", mw.Auth, spareHttHandler.InsertSparePart())
	group.PUT("/:spare_id", mw.Auth, spareHttHandler.UpdateSparePart())
	group.GET("/:spare_id", spareHttHandler.GetSparePartImageByID())
	group.DELETE("/:spare_id", mw.Auth, spareHttHandler.DeleteSparePart())
	group.GET("/search", spareHttHandler.SearchSparePartsPaged())
	group.GET("/filter", spareHttHandler.FilterSparePartByBrandPaginated())

	group.POST("/image", mw.Auth, spareHttHandler.InsertSparePartImage())
	group.GET("/image", spareHttHandler.ListSparePartImagesPaged())
	group.GET("/:spare_id/image", spareHttHandler.ListSparePartImagesBySparePartPaged())
	group.PUT("/image/:image_id", mw.Auth, spareHttHandler.UpdateSparePartImage())
	group.DELETE("/image/:image_id", mw.Auth)

	group.GET("/by/brand/:brand_id", spareHttHandler.FilterSparePartByBrandPaginated())
	group.GET("/by/category/:category_id", spareHttHandler.FilterSparePartByCategory())
	group.GET("/by/car/:car_id", spareHttHandler.FilterSparePartByCarModel())
	group.GET("/by/dealer/:dealer_id", spareHttHandler.ListSparePartsByDealerPaged())
	group.GET("/by/dealership/:dealership_id", spareHttHandler.ListSparePartsByDealershipPaged())
}
