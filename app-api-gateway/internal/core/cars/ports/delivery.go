package ports

import "github.com/gin-gonic/gin"

type Handler interface {
	InsertCarBrand() gin.HandlerFunc
	UpdateCarBrand() gin.HandlerFunc
	GetCarBrandByID() gin.HandlerFunc
	GetCarBrandsPaginated() gin.HandlerFunc
	DeleteCarBrand() gin.HandlerFunc

	InsertCarExtraFeature() gin.HandlerFunc
	UpdateCarExtraFeature() gin.HandlerFunc
	GetCarExtraFeaturePaginated() gin.HandlerFunc
	DeleteCarExtraFeature() gin.HandlerFunc

	CreateCarImage() gin.HandlerFunc
	UpdateCarImage() gin.HandlerFunc
	GetCarImagePaginated() gin.HandlerFunc
	DeleteCarImage() gin.HandlerFunc

	InsertCar() gin.HandlerFunc
	UpdateCar() gin.HandlerFunc
	GetOneCar() gin.HandlerFunc
	GetCarsPaginated() gin.HandlerFunc
	GetCarsByDealershipIDPaginated() gin.HandlerFunc
	GetCarsByBrandIDPaginated() gin.HandlerFunc
	GetCarsByDealerIDPaginated() gin.HandlerFunc
	SearchCarsPaginated() gin.HandlerFunc
	GetCarByField() gin.HandlerFunc
	GetCarByDealerCount() gin.HandlerFunc
	GetCarByDealershipCount() gin.HandlerFunc
	GetCarByBrandCount() gin.HandlerFunc
	DeleteCar() gin.HandlerFunc
}
