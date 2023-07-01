package ports

import (
	"context"
	"garination.com/gateway/internal/core/spares/dto"
	"garination.com/gateway/internal/platform/grpc/app-db/proto"
)

type SpareRepository interface {
	GetSparePartByID(ctx context.Context, request *proto.GetSparePartByIDRequest) (*proto.GetSparePartByIDResponse, error)
	DeleteSparePart(context.Context, *proto.DeleteSparePartRequest) (*proto.DeleteSparePartResponse, error)
	FilterSparePartByBrandPaginated(context.Context, *proto.FilterSparePartByBrandPaginatedRequest) (*proto.FilterSparePartByBrandPaginatedResponse, error)
	FilterSparePartByCategory(context.Context, *proto.FilterSparePartByCategoryPaginatedRequest) (*proto.FilterSparePartByCategoryPaginatedResponse, error)
	FilterSparePartByCarModel(context.Context, *proto.FilterSparePartByCarModelPaginatedRequest) (*proto.FilterSparePartByCarModelPaginatedResponse, error)
	GetSparePartImageByID(context.Context, *proto.GetSparePartImageByIDRequest) (*proto.GetSparePartImageByIDResponse, error)
	InsertSparePart(context.Context, *proto.InsertSparePartRequest) (*proto.InsertSparePartResponse, error)
	InsertSparePartImage(context.Context, *proto.InsertSparePartImageRequest) (*proto.InsertSparePartImageResponse, error)
	ListSparePartImagesBySparePartPaged(context.Context, *proto.ListSparePartImagesBySparePartPagedRequest) (*proto.ListSparePartImagesBySparePartPagedResponse, error)
	ListSparePartImagesPaged(context.Context, *proto.ListSparePartImagesPagedRequest) (*proto.ListSparePartImagesPagedResponse, error)
	ListSparePartsByDealerPaged(context.Context, *proto.ListSparePartsByDealerPagedRequest) (*proto.ListSparePartsByDealerPagedResponse, error)
	ListSparePartsByDealershipPaged(context.Context, *proto.ListSparePartsByDealershipPagedRequest) (*proto.ListSparePartsByDealershipPagedResponse, error)
	ListSparePartsPaged(context.Context, *proto.ListSparePartsPagedRequest) (*proto.ListSparePartsPagedResponse, error)
	SearchSparePartsPaged(context.Context, *proto.SearchSparePartsPagedRequest) (*proto.SearchSparePartsPagedResponse, error)
	UpdateSparePart(context.Context, *proto.UpdateSparePartRequest) (*proto.UpdateSparePartResponse, error)
	UpdateSparePartImage(context.Context, *proto.UpdateSparePartImageRequest) (*proto.UpdateSparePartImageResponse, error)
}

type SpareUseCase interface {
	GetSparePartByID(ctx context.Context, request *dto.GetSparePartByIDRequest) (*dto.GetSparePartByIDResponse, error)
	DeleteSparePart(context.Context, *dto.DeleteSparePartRequest) (*dto.DeleteSparePartResponse, error)
	FilterSparePartByBrandPaginated(context.Context, *dto.FilterSparePartByBrandPaginatedRequest) (*dto.FilterSparePartByBrandPaginatedResponse, error)
	FilterSparePartByCategory(context.Context, *dto.FilterSparePartByCategoryPaginatedRequest) (*dto.FilterSparePartByCategoryPaginatedResponse, error)
	FilterSparePartByCarModel(context.Context, *dto.FilterSparePartByCarModelPaginatedRequest) (*dto.FilterSparePartByCarModelPaginatedResponse, error)
	GetSparePartImageByID(context.Context, *dto.GetSparePartImageByIDRequest) (*dto.GetSparePartImageByIDResponse, error)
	InsertSparePart(context.Context, *dto.InsertSparePartRequest) (*dto.InsertSparePartResponse, error)
	InsertSparePartImage(context.Context, *dto.InsertSparePartImageRequest) (*dto.InsertSparePartImageResponse, error)
	ListSparePartImagesBySparePartPaged(context.Context, *dto.ListSparePartImagesBySparePartPagedRequest) (*dto.ListSparePartImagesBySparePartPagedResponse, error)
	ListSparePartImagesPaged(context.Context, *dto.ListSparePartImagesPagedRequest) (*dto.ListSparePartImagesPagedResponse, error)
	ListSparePartsByDealerPaged(context.Context, *dto.ListSparePartsByDealerPagedRequest) (*dto.ListSparePartsByDealerPagedResponse, error)
	ListSparePartsByDealershipPaged(context.Context, *dto.ListSparePartsByDealershipPagedRequest) (*dto.ListSparePartsByDealershipPagedResponse, error)
	ListSparePartsPaged(context.Context, *dto.ListSparePartsPagedRequest) (*dto.ListSparePartsPagedResponse, error)
	SearchSparePartsPaged(context.Context, *dto.SearchSparePartsPagedRequest) (*dto.SearchSparePartsPagedResponse, error)
	UpdateSparePart(context.Context, *dto.UpdateSparePartRequest) (*dto.UpdateSparePartResponse, error)
	UpdateSparePartImage(context.Context, *dto.UpdateSparePartImageRequest) (*dto.UpdateSparePartImageResponse, error)
}
