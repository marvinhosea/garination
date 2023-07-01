package http

import (
	"garination.com/gateway/internal/core/cars/dto"
	"garination.com/gateway/internal/core/common/model"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
)

func (h handlerImpl) InsertCarExtraFeature() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.CreateCarExtraFeatureRequest

		// bind request
		if err := ctx.ShouldBindJSON(&request); err != nil {
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

		// insert car image
		carExtraFeature, err := h.carUsecase.InsertCarExtraFeature(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
			return
		}

		response.Data = carExtraFeature
		response.Message = "success"
		response.Success = true
		ctx.JSON(200, response)

	}
}

func (h handlerImpl) UpdateCarExtraFeature() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.UpdateCarExtraFeatureRequest

		// get car image id from url params
		carExtraFeatureID := ctx.Param("extra_id")

		// bind request
		if err := ctx.ShouldBindJSON(&request); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
			return
		}

		request.ExtraFeatureID = carExtraFeatureID

		// validate request
		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
			return
		}

		res, err := h.carUsecase.UpdateCarExtraFeature(ctx, &request)
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

func (h handlerImpl) GetCarExtraFeaturePaginated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.GetCarExtraFeaturesPaginatedRequest

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

		res, err := h.carUsecase.GetCarExtraFeaturePaginated(ctx, &request)
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

func (h handlerImpl) DeleteCarExtraFeature() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.DeleteCarExtraFeatureRequest

		// get car image id from url params
		carExtraFeatureID := ctx.Param("extra_id")

		request.ExtraFeatureID = carExtraFeatureID

		// validate request
		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
			return
		}

		res, err := h.carUsecase.DeleteCarExtraFeature(ctx, &request)
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
