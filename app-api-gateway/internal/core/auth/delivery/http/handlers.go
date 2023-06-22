package http

import (
	"errors"
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
			code, res := common.ResponseFromError(errors.Join(err, common.MissingData))
			context.JSON(code, res)
			return
		}

		if err := refreshTokenRequest.Validate(); err != nil {
			code, res := common.ResponseFromError(errors.Join(err, common.InvalidData))
			context.JSON(code, res)
			return
		}

		// call usecase
		refreshTokenResponse, err := h.authUseCase.RefreshToken(context, &refreshTokenRequest)
		if err != nil {
			code, res := common.ResponseFromError(errors.Join(err))
			context.JSON(code, res)
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
			code, res := common.ResponseFromError(errors.Join(err, common.InvalidData))
			context.JSON(code, res)
			return
		}

		// call usecase
		getUserMetaResponse, err := h.authUseCase.GetUserMeta(context, &getUserMetaRequest)
		if err != nil {
			code, res := common.ResponseFromError(errors.Join(err))
			context.JSON(code, res)
			return
		}

		response.Data = getUserMetaResponse
		response.Message = "success"
		response.Success = true
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
			code, res := common.ResponseFromError(errors.Join(err, common.MissingData))
			context.JSON(code, res)
			return
		}

		// validate request
		if err := updateUserMetaRequest.Validate(); err != nil {
			code, res := common.ResponseFromError(errors.Join(err, common.InvalidData))
			context.JSON(code, res)
			return
		}

		// call usecase
		updateUserMetaResponse, err := h.authUseCase.UpdateUserMeta(context, &updateUserMetaRequest)
		if err != nil {
			code, res := common.ResponseFromError(errors.Join(err))
			context.JSON(code, res)
			return
		}

		// return response
		response.Data = updateUserMetaResponse
		response.Success = true
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
			code, res := common.ResponseFromError(errors.Join(err, common.MissingData))
			context.JSON(code, res)
			return
		}

		// validate request
		if err := loginCallbackRequest.Validate(); err != nil {
			code, res := common.ResponseFromError(errors.Join(err, common.InvalidData))
			context.JSON(code, res)
			return
		}

		// call usecase
		loginCallbackResponse, err := h.authUseCase.LoginCallback(context, &loginCallbackRequest)
		if err != nil {
			code, res := common.ResponseFromError(errors.Join(err))
			context.JSON(code, res)
			return
		}

		// return response
		response.Data = loginCallbackResponse
		response.Message = "Successfully logged in"
		response.Success = true
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
			code, res := common.ResponseFromError(errors.Join(err, common.MissingData))
			context.JSON(code, res)
			return
		}

		// validate request
		if err := registerCallbackRequest.Validate(); err != nil {
			code, res := common.ResponseFromError(errors.Join(err, common.InvalidData))
			context.JSON(code, res)
			return
		}

		// call usecase
		registerCallbackResponse, err := h.authUseCase.RegisterCallback(context, &registerCallbackRequest)
		if err != nil {
			code, res := common.ResponseFromError(err)
			context.JSON(code, res)
			return
		}

		// return response
		response.Data = registerCallbackResponse
		response.Message = "Successfully registered"
		response.Success = true
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
			code, res := common.ResponseFromError(errors.Join(err, common.MissingData))
			ctx.JSON(code, res)
			return
		}

		// validate request
		if err := loginRequest.Validate(); err != nil {
			code, res := common.ResponseFromError(errors.Join(err, common.InvalidData))
			ctx.JSON(code, res)
			return
		}

		// call usecase
		loginResponse, err := h.authUseCase.InitiateLogin(ctx, &loginRequest)
		if err != nil {
			code, res := common.ResponseFromError(errors.Join(err))
			ctx.JSON(code, res)
			return
		}

		// update response
		response.Data = loginResponse
		response.Message = "Login request initiated"
		response.Success = true

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
			code, res := common.ResponseFromError(errors.Join(err, common.MissingData))
			ctx.JSON(code, res)
			return
		}

		// validate request
		if err := registerRequest.Validate(); err != nil {
			code, res := common.ResponseFromError(errors.Join(err, common.InvalidData))
			ctx.JSON(code, res)
			return
		}

		// call usecase
		registerResponse, err := h.authUseCase.InitiateRegister(ctx, &registerRequest)
		if err != nil {
			code, res := common.ResponseFromError(errors.Join(err))
			ctx.JSON(code, res)
			return
		}

		// update response
		response.Data = registerResponse
		response.Message = "Register request initiated"
		response.Success = true

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
			code, res := common.ResponseFromError(errors.Join(err, common.MissingData))
			ctx.JSON(code, res)
			return
		}

		// validate request
		if err := logoutRequest.Validate(); err != nil {
			code, res := common.ResponseFromError(errors.Join(err, common.InvalidData))
			ctx.JSON(code, res)
			return
		}

		// call usecase
		logoutResponse, err := h.authUseCase.Logout(ctx, &logoutRequest)
		if err != nil {
			code, res := common.ResponseFromError(errors.Join(err))
			ctx.JSON(code, res)
			return
		}

		// update response
		response.Data = logoutResponse
		response.Message = "Logout request initiated"
		response.Success = true
		// return response
		ctx.JSON(200, response)

	}
}

func NewHandler(authUseCase ports.AuthUsecase) ports.AuthHttpHandler {
	return &handler{
		authUseCase: authUseCase,
	}
}
