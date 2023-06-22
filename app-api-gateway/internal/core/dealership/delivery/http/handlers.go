package http

import (
	"garination.com/gateway/internal/core/common"
	"garination.com/gateway/internal/core/dealership/dto"
	"garination.com/gateway/internal/core/dealership/ports"
	"github.com/gin-gonic/gin"
)

type dealershipHandler struct {
	dealershipUsecase ports.DealerShipUseCase
}

func (d dealershipHandler) GetDealershipByID() gin.HandlerFunc {
	return func(context *gin.Context) {
		var response common.HttpReponse
		var request = dto.GetDealershipRequest{
			ID: context.Param("id"),
		}

		// validate request
		if err := request.Validate(); err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "it'd be nice if you could provide a valid request"
			context.JSON(400, response)
			return
		}

		dealership, err := d.dealershipUsecase.GetDealershipByID(context, &request)
		if err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "something went wrong"
			context.JSON(400, response)
			return
		}

		response.Data = dealership
		response.Message = "success"
		response.Success = true
		context.JSON(200, response)
	}
}

func (d dealershipHandler) GetDealershipByUserID() gin.HandlerFunc {
	return func(context *gin.Context) {
		var response common.HttpReponse
		var request *dto.GetDealershipRequest

		// get user id from url params
		userID := context.Param("id")

		// userReq
		request = &dto.GetDealershipRequest{
			ID: userID,
		}

		// validate request
		if err := request.Validate(); err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "it'd be nice if you could provide a valid request"
			context.JSON(400, response)
			return
		}

		dealership, err := d.dealershipUsecase.GetDealershipByUserID(context, request)
		if err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "something went wrong"
			context.JSON(400, response)
			return
		}

		response.Data = dealership
		response.Message = "success"
		response.Success = true
		context.JSON(200, response)
	}
}

func (d dealershipHandler) UpdateDealership() gin.HandlerFunc {
	return func(context *gin.Context) {
		var response common.HttpReponse
		var request dto.UpdateDealershipRequest

		if err := context.ShouldBindJSON(&request); err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "it'd be nice if you could provide a valid request"
			context.JSON(400, response)
			return
		}

		// validate request
		if err := request.Validate(); err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "it'd be nice if you could provide a valid request"
			context.JSON(400, response)
			return
		}

		dealership, err := d.dealershipUsecase.UpdateDealership(context, &request)
		if err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "something went wrong"
			context.JSON(400, response)
			return
		}

		response.Data = dealership
		response.Message = "success"
		response.Success = true
		context.JSON(200, response)
	}
}

func (d dealershipHandler) CreateDealership() gin.HandlerFunc {
	return func(context *gin.Context) {
		var response common.HttpReponse
		var request dto.CreateDealershipRequest

		if err := context.ShouldBindJSON(&request); err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "it'd be nice if you could provide a valid request"
			context.JSON(400, response)
			return
		}

		// validate request
		if err := request.Validate(); err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "it'd be nice if you could provide a valid request"
			context.JSON(400, response)
			return
		}

		dealership, err := d.dealershipUsecase.CreateDealership(context, &request)
		if err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "something went wrong"
			context.JSON(400, response)
			return
		}

		response.Data = dealership
		response.Message = "success"
		response.Success = true
		context.JSON(200, response)
	}
}

func (d dealershipHandler) DeleteDealership() gin.HandlerFunc {
	return func(context *gin.Context) {
		var response common.HttpReponse
		var request dto.DeleteDealershipRequest

		if err := context.ShouldBindJSON(&request); err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "it'd be nice if you could provide a valid request"
			context.JSON(400, response)
			return
		}

		// validate request
		if err := request.Validate(); err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "it'd be nice if you could provide a valid request"
			context.JSON(400, response)
			return
		}

		dealership, err := d.dealershipUsecase.DeleteDealership(context, &request)
		if err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "something went wrong"
			context.JSON(400, response)
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
