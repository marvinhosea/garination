package http

import (
	"garination.com/gateway/internal/core/auth/dto"
	"garination.com/gateway/internal/core/auth/ports"
	"garination.com/gateway/internal/core/common"
	"github.com/gin-gonic/gin"
)

type handler struct {
	authUseCase ports.AuthUsecase
}

func (h handler) InitiateLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		panic("implement me")

		var response common.HttpReponse

		if ctx.IsAborted() {
			return
		}

		var loginRequest dto.AuthLoginRequest

		if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "Invalid request"
			ctx.JSON(400, response)
			return
		}

		// validate request
		if err := loginRequest.Validate(); err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "Errors in the data sent"
			ctx.JSON(400, response)
			return
		}

		// call usecase
		loginResponse, err := h.authUseCase.InitiateLogin(ctx, &loginRequest)
		if err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "Something went wrong"
			ctx.JSON(500, response)
			return
		}

		// update response
		response.Data = loginResponse
		response.Message = "Login request initiated"

		// return response
		ctx.JSON(200, loginResponse)

	}
}

func (h handler) InitiateRegister() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response common.HttpReponse

		if ctx.IsAborted() {
			return
		}

		var registerRequest dto.AuthRegisterRequest

		if err := ctx.ShouldBindJSON(&registerRequest); err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "Invalid request"
			ctx.JSON(400, response)
			return
		}

		// validate request
		if err := registerRequest.Validate(); err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "Errors in the data sent"
			ctx.JSON(400, response)
			return
		}

		// call usecase
		registerResponse, err := h.authUseCase.InitiateRegister(ctx, &registerRequest)
		if err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "Something went wrong"
			ctx.JSON(500, response)
			return
		}

		// update response
		response.Data = registerResponse
		response.Message = "Register request initiated"

		// return response
		ctx.JSON(200, registerResponse)

	}
}

func (h handler) Logout() gin.HandlerFunc {
	var response common.HttpReponse
	return func(ctx *gin.Context) {
		if ctx.IsAborted() {
			return
		}

		var logoutRequest dto.AuthLogoutRequest

		if err := ctx.ShouldBindJSON(&logoutRequest); err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "Invalid request"
			ctx.JSON(400, response)
			return
		}

		// validate request
		if err := logoutRequest.Validate(); err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "Errors in the data sent"
			ctx.JSON(400, response)
			return
		}

		// call usecase
		logoutResponse, err := h.authUseCase.Logout(ctx, &logoutRequest)
		if err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "Something went wrong"
			ctx.JSON(500, response)
			return
		}

		// update response
		response.Data = logoutResponse
		response.Message = "Logout request initiated"

		// return response
		ctx.JSON(200, logoutResponse)

	}
}

func NewHandler(authUseCase ports.AuthUsecase) ports.AuthHttpHandler {
	return &handler{
		authUseCase: authUseCase,
	}
}
