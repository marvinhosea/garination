package http

import (
	"garination.com/gateway/internal/core/cars/ports"
	"garination.com/gateway/internal/core/middleware"
	"github.com/gin-gonic/gin"
)

func MapCarRoutes(group *gin.RouterGroup, carHttphandler ports.Handler, mw middleware.MiddlewareManager) {
	group.GET("/", carHttphandler.GetCarsPaginated())
	group.POST("/", mw.Auth, carHttphandler.InsertCar())
	group.PUT("/:car_id", mw.Auth, carHttphandler.UpdateCar())
	group.GET("/:car_id", carHttphandler.GetOneCar())
	group.DELETE("/:car_id", mw.Auth, carHttphandler.DeleteCar())
	group.GET("/search", carHttphandler.SearchCarsPaginated())
	group.GET("/filter", carHttphandler.GetCarByField())

	group.POST("/brand", mw.Auth, carHttphandler.InsertCarBrand())
	group.GET("/brand", carHttphandler.GetCarBrandsPaginated())
	group.PUT("/brand/:brand_id", mw.Auth, carHttphandler.UpdateCarBrand())
	group.GET("/brand/:brand_id", carHttphandler.GetCarBrandByID())
	group.DELETE("/brand/:brand_id", mw.Auth, carHttphandler.DeleteCarBrand())

	group.POST("/:car_id/extra", mw.Auth, carHttphandler.InsertCarExtraFeature())
	group.GET("/:car_id/extra", carHttphandler.GetCarExtraFeaturePaginated())
	group.PUT("/extra/:extra_id", mw.Auth, carHttphandler.UpdateCarExtraFeature())
	group.DELETE("/extra/:extra_id", mw.Auth, carHttphandler.DeleteCarExtraFeature())

	group.POST("/:car_id/image", mw.Auth, carHttphandler.CreateCarImage())
	group.GET("/:car_id/image", carHttphandler.GetCarImagePaginated())
	group.PUT("/image/:image_id", mw.Auth, carHttphandler.UpdateCarImage())
	group.DELETE("/image/:image_id", mw.Auth, carHttphandler.DeleteCarImage())

	group.GET("/by/dealership/:dealership_id", carHttphandler.GetCarsByDealershipIDPaginated())
	group.GET("/by/brand/:brand_id", carHttphandler.GetCarsByBrandIDPaginated())
	group.GET("/by/dealer/:dealer_id", carHttphandler.GetCarsByDealerIDPaginated())
	group.GET("/by/dealer/count", carHttphandler.GetCarByDealerCount())
	group.GET("/by/dealership/count", carHttphandler.GetCarByDealershipCount())
	group.GET("/by/brand/count", carHttphandler.GetCarByBrandCount())

}
