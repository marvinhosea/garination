package http

import (
	"garination.com/gateway/internal/core/common/model"
	"garination.com/gateway/internal/core/dealership/dto"
	"garination.com/gateway/internal/core/dealership/ports"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
)

type dealershipHandler struct {
	dealershipUsecase ports.DealerShipUseCase
}

func (d dealershipHandler) GetDealershipByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request = dto.GetDealershipRequest{
			ID: ctx.Param("id"),
		}

		// validate request
		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
			return
		}

		dealership, err := d.dealershipUsecase.GetDealershipByID(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
			return
		}

		response.Data = dealership
		response.Message = "success"
		response.Success = true
		ctx.JSON(200, response)
	}
}

func (d dealershipHandler) GetDealershipByUserID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request *dto.GetDealershipRequest

		// get user id from url params
		userID := ctx.Param("id")

		// userReq
		request = &dto.GetDealershipRequest{
			ID: userID,
		}

		// validate request
		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
			return
		}

		dealership, err := d.dealershipUsecase.GetDealershipByUserID(ctx, request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
			return
		}

		response.Data = dealership
		response.Message = "success"
		response.Success = true
		ctx.JSON(200, response)
	}
}

func (d dealershipHandler) UpdateDealership() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.UpdateDealershipRequest

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

		dealership, err := d.dealershipUsecase.UpdateDealership(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
			return
		}

		response.Data = dealership
		response.Message = "success"
		response.Success = true
		ctx.JSON(200, response)
	}
}

func (d dealershipHandler) CreateDealership() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.CreateDealershipRequest

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

		dealership, err := d.dealershipUsecase.CreateDealership(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
			return
		}

		response.Data = dealership
		response.Message = "success"
		response.Success = true
		ctx.JSON(200, response)
	}
}

func (d dealershipHandler) DeleteDealership() gin.HandlerFunc {
	return func(context *gin.Context) {
		var response model.HttpReponse
		var request dto.DeleteDealershipRequest

		if err := context.ShouldBindJSON(&request); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			context.JSON(code, res)
			return
		}

		// validate request
		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			context.JSON(code, res)
			return
		}

		dealership, err := d.dealershipUsecase.DeleteDealership(context, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			context.JSON(code, res)
			return
		}

		response.Data = dealership
		response.Message = "success"
		response.Success = true
		context.JSON(200, response)
	}
}

func NewDealershipHandler(dealershipUsecase ports.DealerShipUseCase) ports.DealerShiphandler {
	return &dealershipHandler{
		dealershipUsecase: dealershipUsecase,
	}
}
