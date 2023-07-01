package http

import (
	"garination.com/gateway/internal/core/common/model"
	"garination.com/gateway/internal/core/spares/dto"
	"garination.com/gateway/internal/core/spares/ports"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
)

type handlerImpl struct {
	spareUseCase ports.SpareUseCase
}

func (h handlerImpl) GetSpareByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse
		var request dto.GetSparePartByIDRequest

		spareId := ctx.Param("spare_id")
		request.ID = spareId

		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		res, err := h.spareUseCase.GetSparePartByID(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
		}

		response.Data = res
		response.Message = "Spare part retrieved successfully"
		response.Success = true

		ctx.JSON(200, response)
	}
}

func (h handlerImpl) DeleteSparePart() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response model.HttpReponse
		var request dto.DeleteSparePartRequest

		spareId := c.Param("spare_id")
		request.ID = spareId

		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			c.JSON(code, res)
		}

		res, err := h.spareUseCase.DeleteSparePart(c, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			c.JSON(code, res)
		}

		response.Data = res
		response.Message = "Spare part deleted successfully"
		response.Success = true

		c.JSON(200, response)
	}
}

func (h handlerImpl) FilterSparePartByBrandPaginated() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response model.HttpReponse
		var request dto.FilterSparePartByBrandPaginatedRequest

		if err := c.ShouldBindQuery(&request); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			c.JSON(code, res)
		}

		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			c.JSON(code, res)
		}

		res, err := h.spareUseCase.FilterSparePartByBrandPaginated(c, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			c.JSON(code, res)
		}

		response.Data = res
		response.Message = "Spare parts retrieved successfully"
		response.Success = true

		c.JSON(200, response)
	}
}

func (h handlerImpl) FilterSparePartByCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse

		var request dto.FilterSparePartByCategoryPaginatedRequest

		if err := ctx.ShouldBindQuery(&request); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		res, err := h.spareUseCase.FilterSparePartByCategory(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
		}

		response.Data = res
		response.Message = "Spare parts retrieved successfully"
		response.Success = true

		ctx.JSON(200, response)
	}
}

func (h handlerImpl) FilterSparePartByCarModel() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse

		var request dto.FilterSparePartByCarModelPaginatedRequest

		if err := ctx.ShouldBindQuery(&request); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		res, err := h.spareUseCase.FilterSparePartByCarModel(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
		}

		response.Data = res
		response.Message = "Spare parts retrieved successfully"
		response.Success = true

		ctx.JSON(200, response)
	}
}

func (h handlerImpl) GetSparePartImageByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse

		var request dto.GetSparePartImageByIDRequest

		request.ID = ctx.Param("image_id")

		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		res, err := h.spareUseCase.GetSparePartImageByID(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
		}

		response.Data = res
		response.Message = "Spare part image retrieved successfully"
		response.Success = true

		ctx.JSON(200, response)
	}
}

func (h handlerImpl) InsertSparePart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse

		var request dto.InsertSparePartRequest

		if err := ctx.ShouldBindJSON(&request); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		res, err := h.spareUseCase.InsertSparePart(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
		}

		response.Data = res
		response.Message = "Spare part inserted successfully"
		response.Success = true

		ctx.JSON(200, response)
	}
}

func (h handlerImpl) InsertSparePartImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse

		var request dto.InsertSparePartImageRequest

		if err := ctx.ShouldBindJSON(&request); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		res, err := h.spareUseCase.InsertSparePartImage(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
		}

		response.Data = res
		response.Message = "Spare part image inserted successfully"
		response.Success = true

		ctx.JSON(200, response)
	}
}

func (h handlerImpl) ListSparePartImagesBySparePartPaged() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse

		var request dto.ListSparePartImagesBySparePartPagedRequest

		if err := ctx.ShouldBindQuery(&request); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		// get spare part id from path
		request.SparePartId = ctx.Param("spare_id")

		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		res, err := h.spareUseCase.ListSparePartImagesBySparePartPaged(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
		}

		response.Data = res
		response.Message = "Spare part images retrieved successfully"
		response.Success = true

		ctx.JSON(200, response)
	}
}

func (h handlerImpl) ListSparePartImagesPaged() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse

		var request dto.ListSparePartImagesPagedRequest

		if err := ctx.ShouldBindQuery(&request); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		res, err := h.spareUseCase.ListSparePartImagesPaged(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
		}

		response.Data = res
		response.Message = "Spare part images retrieved successfully"
		response.Success = true

		ctx.JSON(200, response)
	}
}

func (h handlerImpl) ListSparePartsByDealerPaged() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse

		var request dto.ListSparePartsByDealerPagedRequest

		if err := ctx.ShouldBindQuery(&request); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		// get dealer id from path
		request.DealerId = ctx.Param("dealer_id")

		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		res, err := h.spareUseCase.ListSparePartsByDealerPaged(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
		}

		response.Data = res
		response.Message = "Spare parts retrieved successfully"
		response.Success = true

		ctx.JSON(200, response)
	}
}

func (h handlerImpl) ListSparePartsByDealershipPaged() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse

		var request dto.ListSparePartsByDealershipPagedRequest

		if err := ctx.ShouldBindQuery(&request); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		// get dealership id from path
		request.DealershipId = ctx.Param("dealership_id")

		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		res, err := h.spareUseCase.ListSparePartsByDealershipPaged(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
		}

		response.Data = res
		response.Message = "Spare parts retrieved successfully"
		response.Success = true

		ctx.JSON(200, response)
	}
}

func (h handlerImpl) ListSparePartsPaged() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse

		var request dto.ListSparePartsPagedRequest

		if err := ctx.ShouldBindQuery(&request); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		res, err := h.spareUseCase.ListSparePartsPaged(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
		}

		response.Data = res
		response.Message = "Spare parts retrieved successfully"
		response.Success = true

		ctx.JSON(200, response)
	}
}

func (h handlerImpl) SearchSparePartsPaged() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse

		var request dto.SearchSparePartsPagedRequest

		if err := ctx.ShouldBindQuery(&request); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		res, err := h.spareUseCase.SearchSparePartsPaged(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
		}

		response.Data = res
		response.Message = "Spare parts retrieved successfully"
		response.Success = true

		ctx.JSON(200, response)
	}
}

func (h handlerImpl) UpdateSparePart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse

		var request dto.UpdateSparePartRequest

		if err := ctx.ShouldBindJSON(&request); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		// get spare part id from path
		request.SparePartId = ctx.Param("spare_id")

		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		res, err := h.spareUseCase.UpdateSparePart(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
		}

		response.Data = res
		response.Message = "Spare part updated successfully"
		response.Success = true

		ctx.JSON(200, response)
	}
}

func (h handlerImpl) UpdateSparePartImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response model.HttpReponse

		var request dto.UpdateSparePartImageRequest

		if err := ctx.ShouldBind(&request); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		// get spare part id from path
		request.SparePartId = ctx.Param("image_id")

		if err := request.Validate(); err != nil {
			code, res := model.ResponseFromErrorWithDetails(codes.InvalidArgument, err.Error())
			ctx.JSON(code, res)
		}

		res, err := h.spareUseCase.UpdateSparePartImage(ctx, &request)
		if err != nil {
			code, res := model.ResponseFromError(err)
			ctx.JSON(code, res)
		}

		response.Data = res
		response.Message = "Spare part image updated successfully"
		response.Success = true

		ctx.JSON(200, response)
	}
}

func NewSpareHandler(spareUseCase ports.SpareUseCase) ports.Handler {
	return &handlerImpl{spareUseCase}
}
