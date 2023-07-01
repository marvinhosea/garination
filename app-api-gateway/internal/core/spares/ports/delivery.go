package ports

import "github.com/gin-gonic/gin"

type Handler interface {
	DeleteSparePart() gin.HandlerFunc
	FilterSparePartByBrandPaginated() gin.HandlerFunc
	FilterSparePartByCategory() gin.HandlerFunc
	FilterSparePartByCarModel() gin.HandlerFunc
	GetSparePartImageByID() gin.HandlerFunc
	GetSpareByID() gin.HandlerFunc
	InsertSparePart() gin.HandlerFunc
	InsertSparePartImage() gin.HandlerFunc
	ListSparePartImagesBySparePartPaged() gin.HandlerFunc
	ListSparePartImagesPaged() gin.HandlerFunc
	ListSparePartsByDealerPaged() gin.HandlerFunc
	ListSparePartsByDealershipPaged() gin.HandlerFunc
	ListSparePartsPaged() gin.HandlerFunc
	SearchSparePartsPaged() gin.HandlerFunc
	UpdateSparePart() gin.HandlerFunc
	UpdateSparePartImage() gin.HandlerFunc
}
