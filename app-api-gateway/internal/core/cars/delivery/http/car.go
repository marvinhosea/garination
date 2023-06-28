package http

import (
	"garination.com/gateway/internal/core/cars/dto"
	"garination.com/gateway/internal/core/common/model"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
)

func (h handlerImpl) InsertCar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.InsertCarRequest

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

		// insert car
		car, err := h.carUsecase.InsertCar(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
			return
		}

		response.Data = car
		response.Message = "success"
		response.Success = true
		ctx.JSON(200, response)
	}
}

func (h handlerImpl) UpdateCar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.UpdateCarRequest

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

		// update car
		car, err := h.carUsecase.UpdateCar(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
			return
		}

		response.Data = car
	}
}

func (h handlerImpl) GetOneCar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.GetOneCarRequest

		// get car id from url params
		request.CarID = ctx.Param("car_id")

		// validate request
		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
			return
		}

		// get car
		car, err := h.carUsecase.GetOneCar(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
			return
		}

		response.Data = car
		response.Message = "success"
		response.Success = true
		ctx.JSON(200, response)
	}
}

func (h handlerImpl) GetCarsPaginated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.GetCarsPaginatedRequest

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

		// get cars
		cars, err := h.carUsecase.GetCarsPaginated(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
			return
		}

		response.Data = cars
		response.Message = "success"
		response.Success = true
		ctx.JSON(200, response)
	}
}

func (h handlerImpl) GetCarsByDealershipIDPaginated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.GetCarsByDealershipIDPaginatedRequest

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

		// get cars
		cars, err := h.carUsecase.GetCarsByDealershipIDPaginated(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
			return
		}

		response.Data = cars
		response.Message = "success"
		response.Success = true
		ctx.JSON(200, response)
	}
}

func (h handlerImpl) GetCarsByBrandIDPaginated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.GetCarsByBrandIDPaginatedRequest

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

		// get cars
		cars, err := h.carUsecase.GetCarsByBrandIDPaginated(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
			return
		}

		response.Data = cars
		response.Message = "success"
		response.Success = true
	}
}

func (h handlerImpl) GetCarsByDealerIDPaginated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.GetCarsByDealerIDPaginatedRequest

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

		// get cars
		cars, err := h.carUsecase.GetCarsByDealerIDPaginated(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
			return
		}

		response.Data = cars
		response.Message = "success"
		response.Success = true
		ctx.JSON(200, response)
	}
}

func (h handlerImpl) SearchCarsPaginated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.SearchCarsPaginatedRequest

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

		// get cars
		cars, err := h.carUsecase.SearchCarsPaginated(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
			return
		}

		response.Data = cars
		response.Message = "success"
		response.Success = true
		ctx.JSON(200, response)
	}
}

func (h handlerImpl) GetCarByField() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.GetCarsByFieldRequest

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

		// get cars
		cars, err := h.carUsecase.GetCarsByField(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
			return
		}

		response.Data = cars
		response.Message = "success"
		response.Success = true
		ctx.JSON(200, response)
	}
}

func (h handlerImpl) GetCarByDealerCount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.GetCarByDealerCountRequest

		// get from params
		request.DealerID = ctx.Param("dealer_id")

		// validate request
		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
			return
		}

		// get cars
		cars, err := h.carUsecase.GetCarByDealerCount(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
			return
		}

		response.Data = cars
		response.Message = "success"
		response.Success = true
		ctx.JSON(200, response)
	}
}

func (h handlerImpl) GetCarByDealershipCount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.GetCarByDealershipCountRequest

		// get from params
		request.DealershipID = ctx.Param("dealership_id")

		// validate request
		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
			return
		}

		// get cars
		cars, err := h.carUsecase.GetCarByDealershipCount(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
			return
		}

		response.Data = cars
		response.Message = "success"
		response.Success = true
		ctx.JSON(200, response)
	}
}

func (h handlerImpl) GetCarByBrandCount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.GetCarByBrandCountRequest

		// get from params
		request.BrandID = ctx.Param("brand_id")

		// validate request
		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
			return
		}

		// get cars
		cars, err := h.carUsecase.GetCarByBrandCount(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
			return
		}

		response.Data = cars
		response.Message = "success"
		response.Success = true
		ctx.JSON(200, response)
	}
}

func (h handlerImpl) DeleteCar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.DeleteCarRequest

		// get from params
		request.CarID = ctx.Param("car_id")

		// validate request
		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
			return
		}

		// delete car
		res, err := h.carUsecase.DeleteCar(ctx, &request)
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
