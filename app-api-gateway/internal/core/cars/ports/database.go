package ports

import (
	"context"
	"garination.com/gateway/internal/core/cars/dto"
	"garination.com/gateway/internal/platform/grpc/app-db/proto"
	"google.golang.org/grpc"
)

type CarDatabaseRepo interface {
	InsertCarBrand(ctx context.Context, in *proto.InsertCarBrandRequest, opts ...grpc.CallOption) (*proto.InsertCarBrandResponse, error)
	UpdateCarBrand(ctx context.Context, in *proto.UpdateCarBrandRequest, opts ...grpc.CallOption) (*proto.UpdateCarBrandResponse, error)
	GetCarBrandByID(ctx context.Context, in *proto.GetCarBrandByIDRequest, opts ...grpc.CallOption) (*proto.GetCarBrandByIDResponse, error)
	GetCarBrandsPaginated(ctx context.Context, in *proto.GetCarBrandsPaginatedRequest, opts ...grpc.CallOption) (*proto.GetCarBrandsPaginatedResponse, error)
	DeleteCarBrand(ctx context.Context, in *proto.DeleteCarBrandRequest, opts ...grpc.CallOption) (*proto.DeleteCarBrandResponse, error)
	InsertCarExtraFeature(ctx context.Context, in *proto.InsertCarExtraFeatureRequest, opts ...grpc.CallOption) (*proto.InsertCarExtraFeatureResponse, error)
	UpdateCarExtraFeature(ctx context.Context, in *proto.UpdateCarExtraFeatureRequest, opts ...grpc.CallOption) (*proto.UpdateCarExtraFeatureResponse, error)
	GetCarExtraFeaturePaginated(ctx context.Context, in *proto.GetCarExtraFeaturesPaginatedRequest, opts ...grpc.CallOption) (*proto.GetCarExtraFeaturesPaginatedResponse, error)
	DeleteCarExtraFeature(ctx context.Context, in *proto.DeleteCarExtraFeatureRequest, opts ...grpc.CallOption) (*proto.DeleteCarExtraFeatureResponse, error)
	CreateCarImage(ctx context.Context, in *proto.CreateCarImageRequest, opts ...grpc.CallOption) (*proto.CreateCarImageResponse, error)
	UpdateCarImage(ctx context.Context, in *proto.UpdateCarImageRequest, opts ...grpc.CallOption) (*proto.UpdateCarImageResponse, error)
	GetCarImagePaginated(ctx context.Context, in *proto.GetCarImagesPaginatedRequest, opts ...grpc.CallOption) (*proto.GetCarImagesPaginatedResponse, error)
	DeleteCarImage(ctx context.Context, in *proto.DeleteCarImageRequest, opts ...grpc.CallOption) (*proto.DeleteCarImageResponse, error)
	InsertCar(ctx context.Context, in *proto.InsertCarRequest, opts ...grpc.CallOption) (*proto.InsertCarResponse, error)
	UpdateCar(ctx context.Context, in *proto.UpdateCarRequest, opts ...grpc.CallOption) (*proto.UpdateCarResponse, error)
	GetOneCar(ctx context.Context, in *proto.GetOneCarRequest, opts ...grpc.CallOption) (*proto.GetOneCarResponse, error)
	GetCarsPaginated(ctx context.Context, in *proto.GetCarsPaginatedRequest, opts ...grpc.CallOption) (*proto.GetCarsPaginatedResponse, error)
	GetCarsByDealershipIDPaginated(ctx context.Context, in *proto.GetCarsByDealershipIDPaginatedRequest, opts ...grpc.CallOption) (*proto.GetCarsByDealershipIDPaginatedResponse, error)
	GetCarsByBrandIDPaginated(ctx context.Context, in *proto.GetCarsByBrandIDPaginatedRequest, opts ...grpc.CallOption) (*proto.GetCarsByCarBrandIDPaginatedResponse, error)
	GetCarsByDealerIDPaginated(ctx context.Context, in *proto.GetCarsByDealerIDPaginatedRequest, opts ...grpc.CallOption) (*proto.GetCarsByDealerIDPaginatedResponse, error)
	SearchCarsPaginated(ctx context.Context, in *proto.SearchCarsPaginatedRequest, opts ...grpc.CallOption) (*proto.SearchCarsPaginatedResponse, error)
	GetCarByField(ctx context.Context, in *proto.GetCarByFieldRequest, opts ...grpc.CallOption) (*proto.GetCarByFieldResponse, error)
	GetCarByDealerCount(ctx context.Context, in *proto.GetCarByDealerCountRequest, opts ...grpc.CallOption) (*proto.GetCarByDealerCountResponse, error)
	GetCarByDealershipCount(ctx context.Context, in *proto.GetCarByDealershipCountRequest, opts ...grpc.CallOption) (*proto.GetCarByDealershipCountResponse, error)
	GetCarByBrandCount(ctx context.Context, in *proto.GetCarByBrandCountRequest, opts ...grpc.CallOption) (*proto.GetCarByCarBrandCountResponse, error)
	DeleteCar(ctx context.Context, in *proto.DeleteCarRequest, opts ...grpc.CallOption) (*proto.DeleteCarResponse, error)
}

type CarUsercase interface {
	InsertCarBrand(ctx context.Context, in *dto.CreateCarBrandRequest) (*dto.CreateCarBrandResponse, error)
	UpdateCarBrand(ctx context.Context, in *dto.UpdateCarBrandRequest) (*dto.UpdateCarBrandResponse, error)
	GetCarBrandByID(ctx context.Context, in *dto.GetCarBrandByIDRequest) (*dto.GetCarBrandByIDResponse, error)
	GetCarBrandsPaginated(ctx context.Context, in *dto.GetCarBrandsPaginatedRequest) (*dto.GetCarBrandsPaginatedResponse, error)
	DeleteCarBrand(ctx context.Context, in *dto.DeleteCarBrandRequest) (*dto.DeleteCarBrandResponse, error)
	InsertCarExtraFeature(ctx context.Context, in *dto.CreateCarExtraFeatureRequest) (*dto.CreateCarExtraFeatureResponse, error)
	UpdateCarExtraFeature(ctx context.Context, in *dto.UpdateCarExtraFeatureRequest) (*dto.UpdateCarExtraFeatureResponse, error)
	GetCarExtraFeaturePaginated(ctx context.Context, in *dto.GetCarExtraFeaturesPaginatedRequest) (*dto.GetCarExtraFeaturesPaginatedResponse, error)
	DeleteCarExtraFeature(ctx context.Context, in *dto.DeleteCarExtraFeatureRequest) (*dto.DeleteCarExtraFeatureResponse, error)
	CreateCarImage(ctx context.Context, in *dto.CreateCarImageRequest) (*dto.CreateCarImageResponse, error)
	UpdateCarImage(ctx context.Context, in *dto.UpdateCarImageRequest) (*dto.UpdateCarImageResponse, error)
	GetCarImagePaginated(ctx context.Context, in *dto.GetCarImagesPaginatedRequest) (*dto.GetCarImagesPaginatedResponse, error)
	DeleteCarImage(ctx context.Context, in *dto.DeleteCarImageRequest) (*dto.DeleteCarImageResponse, error)
	InsertCar(ctx context.Context, in *dto.InsertCarRequest) (*dto.InsertCarResponse, error)
	UpdateCar(ctx context.Context, in *dto.UpdateCarRequest) (*dto.UpdateCarResponse, error)
	GetOneCar(ctx context.Context, in *dto.GetOneCarRequest) (*dto.GetOneCarResponse, error)
	GetCarsPaginated(ctx context.Context, in *dto.GetCarsPaginatedRequest) (*dto.GetCarsPaginatedResponse, error)
	GetCarsByDealershipIDPaginated(ctx context.Context, in *dto.GetCarsByDealershipIDPaginatedRequest) (*dto.GetCarsByDealershipIDPaginatedResponse, error)
	GetCarsByBrandIDPaginated(ctx context.Context, in *dto.GetCarsByBrandIDPaginatedRequest) (*dto.GetCarsByCarBrandIDPaginatedResponse, error)
	GetCarsByDealerIDPaginated(ctx context.Context, in *dto.GetCarsByDealerIDPaginatedRequest) (*dto.GetCarsByDealerIDPaginatedResponse, error)
	SearchCarsPaginated(ctx context.Context, in *dto.SearchCarsPaginatedRequest) (*dto.SearchCarsPaginatedResponse, error)
	GetCarsByField(ctx context.Context, in *dto.GetCarsByFieldRequest) (*dto.GetCarsByFieldResponse, error)
	GetCarByDealerCount(ctx context.Context, in *dto.GetCarByDealerCountRequest) (*dto.GetCarByDealerCountResponse, error)
	GetCarByDealershipCount(ctx context.Context, in *dto.GetCarByDealershipCountRequest) (*dto.GetCarByDealershipCountResponse, error)
	GetCarByBrandCount(ctx context.Context, in *dto.GetCarByBrandCountRequest) (*dto.GetCarByCarBrandCountResponse, error)
	DeleteCar(ctx context.Context, in *dto.DeleteCarRequest) (*dto.DeleteCarResponse, error)
}
