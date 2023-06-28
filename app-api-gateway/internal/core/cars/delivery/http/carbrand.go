package http

import (
	"garination.com/gateway/internal/core/cars/dto"
	"garination.com/gateway/internal/core/common/model"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
)

func (h handlerImpl) InsertCarBrand() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.CreateCarBrandRequest

		// bind request
		if err := ctx.ShouldBind(&request); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
			return
		}

		// validate request
		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
			return
		}

		// insert car brand
		carBrand, err := h.carUsecase.InsertCarBrand(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
			return
		}

		response.Data = carBrand
		response.Message = "success"
		response.Success = true
		ctx.JSON(200, response)
	}
}

func (h handlerImpl) UpdateCarBrand() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.UpdateCarBrandRequest

		// get car brand id from url params
		carBrandID := ctx.Param("brand_id")

		// bind request
		if err := ctx.ShouldBindJSON(&request); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
			return
		}

		request.BrandID = carBrandID

		// validate request
		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
			return
		}

		// update car brand
		carBrand, err := h.carUsecase.UpdateCarBrand(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
			return
		}

		response.Data = carBrand
		response.Message = "success"
		response.Success = true
		ctx.JSON(200, response)

	}
}

func (h handlerImpl) GetCarBrandByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.GetCarBrandByIDRequest

		// get car brand id from url params
		carBrandID := ctx.Param("brand_id")

		// bind request
		if err := ctx.ShouldBindQuery(&request); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
			return
		}

		request.BrandID = carBrandID

		// validate request
		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
			return
		}

		// get car brand
		carBrand, err := h.carUsecase.GetCarBrandByID(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
			return
		}

		response.Data = carBrand
		response.Message = "success"
		response.Success = true
		ctx.JSON(200, response)
	}
}

func (h handlerImpl) GetCarBrandsPaginated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.GetCarBrandsPaginatedRequest

		// bind request
		if err := ctx.ShouldBindQuery(&request); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
			return
		}

		// validate request
		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
			return
		}

		// get car brands
		carBrands, err := h.carUsecase.GetCarBrandsPaginated(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
			return
		}

		response.Data = carBrands
		response.Message = "success"
		response.Success = true
		ctx.JSON(200, response)
	}
}

func (h handlerImpl) DeleteCarBrand() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.DeleteCarBrandRequest

		// get car brand id from url params
		carBrandID := ctx.Param("brand_id")

		// bind request
		if err := ctx.ShouldBindQuery(&request); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
			return
		}

		request.BrandID = carBrandID

		// validate request
		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
			return
		}

		// delete car brand
		res, err := h.carUsecase.DeleteCarBrand(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
			return
		}

		response.Data = res
		response.Message = "success"
		response.Success = true
		ctx.JSON(200, response)
	}
}
