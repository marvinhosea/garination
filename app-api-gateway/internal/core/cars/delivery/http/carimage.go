package http

import (
	"garination.com/gateway/internal/core/cars/dto"
	"garination.com/gateway/internal/core/common/model"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
)

func (h handlerImpl) CreateCarImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.CreateCarImageRequest

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

		// insert car image
		carImage, err := h.carUsecase.CreateCarImage(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
			return
		}

		response.Data = carImage
		response.Message = "success"
		response.Success = true
		ctx.JSON(200, response)
	}
}

func (h handlerImpl) UpdateCarImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.UpdateCarImageRequest

		// get car image id from url params
		carImageID := ctx.Param("image_id")

		// bind request
		if err := ctx.ShouldBind(&request); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
			return
		}

		request.ImageID = carImageID

		// validate request
		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
			return
		}

		// update car image
		carImage, err := h.carUsecase.UpdateCarImage(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
			return
		}

		response.Data = carImage
		response.Message = "success"
		response.Success = true
		ctx.JSON(200, response)
	}
}

func (h handlerImpl) GetCarImagePaginated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.GetCarImagesPaginatedRequest

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

		// get car images
		carImages, err := h.carUsecase.GetCarImagePaginated(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
			return
		}

		response.Data = carImages
		response.Message = "success"
		response.Success = true
		ctx.JSON(200, response)
	}
}

func (h handlerImpl) DeleteCarImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.DeleteCarImageRequest

		// get car image id from url params
		carImageID := ctx.Param("image_id")

		request.ImageID = carImageID

		// validate request
		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
			return
		}

		// delete car image
		res, err := h.carUsecase.DeleteCarImage(ctx, &request)
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
