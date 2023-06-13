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

func (h handler) RefreshToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		var response common.HttpReponse
		if context.IsAborted() {
			return
		}

		var refreshTokenRequest dto.AuthRefreshTokenRequest
		if err := context.ShouldBindJSON(&refreshTokenRequest); err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "we received an invalid json body"
			context.JSON(400, response)
			return
		}

		if err := refreshTokenRequest.Validate(); err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "we did not receive the required fields"
			context.JSON(400, response)
			return
		}

		// call usecase
		refreshTokenResponse, err := h.authUseCase.RefreshToken(context, &refreshTokenRequest)
		if err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "Internal server error"
			context.JSON(500, response)
			return
		}

		response.Data = refreshTokenResponse
		response.Message = "success"
		response.Success = true
		context.JSON(200, response)
	}
}

func (h handler) GetUserMeta() gin.HandlerFunc {
	return func(context *gin.Context) {
		var response common.HttpReponse
		if context.IsAborted() {
			return
		}

		var getUserMetaRequest dto.AuthGetUserMetaRequest
		// get user id from  :user_id
		getUserMetaRequest.UserID = context.Param("user_id")

		if err := getUserMetaRequest.Validate(); err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "Bad request"
			context.JSON(400, response)
			return
		}

		// call usecase
		getUserMetaResponse, err := h.authUseCase.GetUserMeta(context, &getUserMetaRequest)
		if err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "Internal server error"
			context.JSON(500, response)
			return
		}

		response.Data = getUserMetaResponse
		context.JSON(200, response)

	}
}

func (h handler) UpdateUserMeta() gin.HandlerFunc {
	return func(context *gin.Context) {
		var response common.HttpReponse
		if context.IsAborted() {
			return
		}

		var updateUserMetaRequest dto.AuthUpdateUserMetaRequest

		if err := context.ShouldBindJSON(&updateUserMetaRequest); err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "Bad request"
			context.JSON(400, response)
			return
		}

		// validate request
		if err := updateUserMetaRequest.Validate(); err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "Errors in the data"
			context.JSON(400, response)
			return
		}

		// call usecase
		updateUserMetaResponse, err := h.authUseCase.UpdateUserMeta(context, &updateUserMetaRequest)
		if err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "Error while updating user meta"
			context.JSON(400, response)
			return
		}

		// return response
		response.Data = updateUserMetaResponse
		response.Message = "Successfully updated user meta"
		context.JSON(200, response)
	}
}

func (h handler) LoginCallback() gin.HandlerFunc {
	return func(context *gin.Context) {
		var response common.HttpReponse
		if context.IsAborted() {
			return
		}

		var loginCallbackRequest dto.AuthLoginCallbackRequest

		if err := context.ShouldBindJSON(&loginCallbackRequest); err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "Bad request"
			context.JSON(400, response)
			return
		}

		// validate request
		if err := loginCallbackRequest.Validate(); err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "Errors in the data"
			context.JSON(400, response)
			return
		}

		// call usecase
		loginCallbackResponse, err := h.authUseCase.LoginCallback(context, &loginCallbackRequest)
		if err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "Error while logging in"
			context.JSON(400, response)
			return
		}

		// return response
		response.Data = loginCallbackResponse
		response.Message = "Successfully logged in"
		context.JSON(200, response)
	}
}

func (h handler) RegisterCallback() gin.HandlerFunc {
	return func(context *gin.Context) {
		var response common.HttpReponse
		if context.IsAborted() {
			return
		}

		var registerCallbackRequest dto.AuthRegisterCallbackRequest

		if err := context.ShouldBindJSON(&registerCallbackRequest); err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "Bad request"
			context.JSON(400, response)
			return
		}

		// validate request
		if err := registerCallbackRequest.Validate(); err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "Errors in the data"
			context.JSON(400, response)
			return
		}

		// call usecase
		registerCallbackResponse, err := h.authUseCase.RegisterCallback(context, &registerCallbackRequest)
		if err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Message = "Error while registering"
			context.JSON(400, response)
			return
		}

		// return response
		response.Data = registerCallbackResponse
		response.Message = "Successfully registered"
		context.JSON(200, response)
	}
}

func (h handler) InitiateLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {

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
		ctx.JSON(200, response)

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
		ctx.JSON(200, response)

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
		ctx.JSON(200, response)

	}
}

func NewHandler(authUseCase ports.AuthUsecase) ports.AuthHttpHandler {
	return &handler{
		authUseCase: authUseCase,
	}
}
