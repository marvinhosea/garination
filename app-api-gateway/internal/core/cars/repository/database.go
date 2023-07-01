package repository

import (
	"context"
	"garination.com/gateway/internal/core/cars/ports"
	"garination.com/gateway/internal/platform/grpc/app-db/proto"
	"google.golang.org/grpc"
)

type databaseRepoImpl struct {
	dbService proto.DatabaseServiceClient
}

func (d databaseRepoImpl) InsertCarBrand(ctx context.Context, in *proto.InsertCarBrandRequest, opts ...grpc.CallOption) (*proto.InsertCarBrandResponse, error) {
	return d.dbService.InsertCarBrand(ctx, in, opts...)
}

func (d databaseRepoImpl) UpdateCarBrand(ctx context.Context, in *proto.UpdateCarBrandRequest, opts ...grpc.CallOption) (*proto.UpdateCarBrandResponse, error) {
	return d.dbService.UpdateCarBrand(ctx, in, opts...)
}

func (d databaseRepoImpl) GetCarBrandByID(ctx context.Context, in *proto.GetCarBrandByIDRequest, opts ...grpc.CallOption) (*proto.GetCarBrandByIDResponse, error) {
	return d.dbService.GetCarBrandByID(ctx, in, opts...)
}

func (d databaseRepoImpl) GetCarBrandsPaginated(ctx context.Context, in *proto.GetCarBrandsPaginatedRequest, opts ...grpc.CallOption) (*proto.GetCarBrandsPaginatedResponse, error) {
	return d.dbService.GetCarBrandsPaginated(ctx, in, opts...)
}

func (d databaseRepoImpl) DeleteCarBrand(ctx context.Context, in *proto.DeleteCarBrandRequest, opts ...grpc.CallOption) (*proto.DeleteCarBrandResponse, error) {
	return d.dbService.DeleteCarBrand(ctx, in, opts...)
}

func (d databaseRepoImpl) InsertCarExtraFeature(ctx context.Context, in *proto.InsertCarExtraFeatureRequest, opts ...grpc.CallOption) (*proto.InsertCarExtraFeatureResponse, error) {
	return d.dbService.InsertCarExtraFeature(ctx, in, opts...)
}

func (d databaseRepoImpl) UpdateCarExtraFeature(ctx context.Context, in *proto.UpdateCarExtraFeatureRequest, opts ...grpc.CallOption) (*proto.UpdateCarExtraFeatureResponse, error) {
	return d.dbService.UpdateCarExtraFeature(ctx, in, opts...)
}

func (d databaseRepoImpl) GetCarExtraFeaturePaginated(ctx context.Context, in *proto.GetCarExtraFeaturesPaginatedRequest, opts ...grpc.CallOption) (*proto.GetCarExtraFeaturesPaginatedResponse, error) {
	return d.dbService.GetCarExtraFeaturePaginated(ctx, in, opts...)
}

func (d databaseRepoImpl) DeleteCarExtraFeature(ctx context.Context, in *proto.DeleteCarExtraFeatureRequest, opts ...grpc.CallOption) (*proto.DeleteCarExtraFeatureResponse, error) {
	return d.dbService.DeleteCarExtraFeature(ctx, in, opts...)
}

func (d databaseRepoImpl) CreateCarImage(ctx context.Context, in *proto.CreateCarImageRequest, opts ...grpc.CallOption) (*proto.CreateCarImageResponse, error) {
	return d.dbService.CreateCarImage(ctx, in, opts...)
}

func (d databaseRepoImpl) UpdateCarImage(ctx context.Context, in *proto.UpdateCarImageRequest, opts ...grpc.CallOption) (*proto.UpdateCarImageResponse, error) {
	return d.dbService.UpdateCarImage(ctx, in, opts...)
}

func (d databaseRepoImpl) GetCarImagePaginated(ctx context.Context, in *proto.GetCarImagesPaginatedRequest, opts ...grpc.CallOption) (*proto.GetCarImagesPaginatedResponse, error) {
	return d.dbService.GetCarImagePaginated(ctx, in, opts...)
}

func (d databaseRepoImpl) DeleteCarImage(ctx context.Context, in *proto.DeleteCarImageRequest, opts ...grpc.CallOption) (*proto.DeleteCarImageResponse, error) {
	return d.dbService.DeleteCarImage(ctx, in, opts...)
}

func (d databaseRepoImpl) InsertCar(ctx context.Context, in *proto.InsertCarRequest, opts ...grpc.CallOption) (*proto.InsertCarResponse, error) {
	return d.dbService.InsertCar(ctx, in, opts...)
}

func (d databaseRepoImpl) UpdateCar(ctx context.Context, in *proto.UpdateCarRequest, opts ...grpc.CallOption) (*proto.UpdateCarResponse, error) {
	return d.dbService.UpdateCar(ctx, in, opts...)
}

func (d databaseRepoImpl) GetOneCar(ctx context.Context, in *proto.GetOneCarRequest, opts ...grpc.CallOption) (*proto.GetOneCarResponse, error) {
	return d.dbService.GetOneCar(ctx, in, opts...)
}

func (d databaseRepoImpl) GetCarsPaginated(ctx context.Context, in *proto.GetCarsPaginatedRequest, opts ...grpc.CallOption) (*proto.GetCarsPaginatedResponse, error) {
	return d.dbService.GetCarsPaginated(ctx, in, opts...)
}

func (d databaseRepoImpl) GetCarsByDealershipIDPaginated(ctx context.Context, in *proto.GetCarsByDealershipIDPaginatedRequest, opts ...grpc.CallOption) (*proto.GetCarsByDealershipIDPaginatedResponse, error) {
	return d.dbService.GetCarsByDealershipIDPaginated(ctx, in, opts...)
}

func (d databaseRepoImpl) GetCarsByBrandIDPaginated(ctx context.Context, in *proto.GetCarsByBrandIDPaginatedRequest, opts ...grpc.CallOption) (*proto.GetCarsByCarBrandIDPaginatedResponse, error) {
	return d.dbService.GetCarsByBrandIDPaginated(ctx, in, opts...)
}

func (d databaseRepoImpl) GetCarsByDealerIDPaginated(ctx context.Context, in *proto.GetCarsByDealerIDPaginatedRequest, opts ...grpc.CallOption) (*proto.GetCarsByDealerIDPaginatedResponse, error) {
	return d.dbService.GetCarsByDealerIDPaginated(ctx, in, opts...)
}

func (d databaseRepoImpl) SearchCarsPaginated(ctx context.Context, in *proto.SearchCarsPaginatedRequest, opts ...grpc.CallOption) (*proto.SearchCarsPaginatedResponse, error) {
	return d.dbService.SearchCarsPaginated(ctx, in, opts...)
}

func (d databaseRepoImpl) GetCarByField(ctx context.Context, in *proto.GetCarByFieldRequest, opts ...grpc.CallOption) (*proto.GetCarByFieldResponse, error) {
	return d.dbService.GetCarByField(ctx, in, opts...)
}

func (d databaseRepoImpl) GetCarByDealerCount(ctx context.Context, in *proto.GetCarByDealerCountRequest, opts ...grpc.CallOption) (*proto.GetCarByDealerCountResponse, error) {
	return d.dbService.GetCarByDealerCount(ctx, in, opts...)
}

func (d databaseRepoImpl) GetCarByDealershipCount(ctx context.Context, in *proto.GetCarByDealershipCountRequest, opts ...grpc.CallOption) (*proto.GetCarByDealershipCountResponse, error) {
	return d.dbService.GetCarByDealershipCount(ctx, in, opts...)
}

func (d databaseRepoImpl) GetCarByBrandCount(ctx context.Context, in *proto.GetCarByBrandCountRequest, opts ...grpc.CallOption) (*proto.GetCarByCarBrandCountResponse, error) {
	return d.dbService.GetCarByBrandCount(ctx, in, opts...)
}

func (d databaseRepoImpl) DeleteCar(ctx context.Context, in *proto.DeleteCarRequest, opts ...grpc.CallOption) (*proto.DeleteCarResponse, error) {
	return d.dbService.DeleteCar(ctx, in, opts...)
}

func NewCarRepo(dbService proto.DatabaseServiceClient) ports.CarDatabaseRepo {
	return &databaseRepoImpl{dbService: dbService}
}
