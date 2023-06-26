package usecase

import (
	"context"
	"garination.com/gateway/internal/core/cars/dto"
	"garination.com/gateway/internal/core/cars/ports"
)

type usecaseImpl struct {
	dbRepo ports.CarDatabaseRepo
}

func (u usecaseImpl) InsertCarBrand(ctx context.Context, in *dto.CreateCarBrandRequest) (*dto.CreateCarBrandResponse, error) {

}

func (u usecaseImpl) UpdateCarBrand(ctx context.Context, in *dto.UpdateCarBrandRequest) (*dto.UpdateCarBrandResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecaseImpl) GetCarBrandByID(ctx context.Context, in *dto.GetCarBrandByIDRequest) (*dto.GetCarBrandByIDResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecaseImpl) GetCarBrandsPaginated(ctx context.Context, in *dto.GetCarBrandsPaginatedRequest) (*dto.GetCarBrandsPaginatedResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecaseImpl) DeleteCarBrand(ctx context.Context, in *dto.DeleteCarBrandRequest) (*dto.DeleteCarBrandResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecaseImpl) InsertCarExtraFeature(ctx context.Context, in *dto.CreateCarExtraFeatureRequest) (*dto.CreateCarExtraFeatureResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecaseImpl) UpdateCarExtraFeature(ctx context.Context, in *dto.UpdateCarExtraFeatureRequest) (*dto.UpdateCarExtraFeatureResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecaseImpl) GetCarExtraFeaturePaginated(ctx context.Context, in *dto.GetCarExtraFeaturesPaginatedRequest) (*dto.GetCarExtraFeaturesPaginatedResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecaseImpl) DeleteCarExtraFeature(ctx context.Context, in *dto.DeleteCarExtraFeatureRequest) (*dto.DeleteCarExtraFeatureResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecaseImpl) CreateCarImage(ctx context.Context, in *dto.CreateCarImageRequest) (*dto.CreateCarImageResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecaseImpl) UpdateCarImage(ctx context.Context, in *dto.UpdateCarImageRequest) (*dto.UpdateCarImageResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecaseImpl) GetCarImagePaginated(ctx context.Context, in *dto.GetCarImagesPaginatedRequest) (*dto.GetCarImagesPaginatedResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecaseImpl) DeleteCarImage(ctx context.Context, in *dto.DeleteCarImageRequest) (*dto.DeleteCarImageResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecaseImpl) InsertCar(ctx context.Context, in *dto.InsertCarRequest) (*dto.InsertCarResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecaseImpl) UpdateCar(ctx context.Context, in *dto.UpdateCarRequest) (*dto.UpdateCarResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecaseImpl) GetOneCar(ctx context.Context, in *dto.GetOneCarRequest) (*dto.GetOneCarResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecaseImpl) GetCarsPaginated(ctx context.Context, in *dto.GetCarsPaginatedRequest) (*dto.GetCarsPaginatedResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecaseImpl) GetCarsByDealershipIDPaginated(ctx context.Context, in *dto.GetCarsByDealershipIDPaginatedRequest) (*dto.GetCarsByDealershipIDPaginatedResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecaseImpl) GetCarsByBrandIDPaginated(ctx context.Context, in *dto.GetCarsByBrandIDPaginatedRequest) (*dto.GetCarsByCarBrandIDPaginatedResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecaseImpl) GetCarsByDealerIDPaginated(ctx context.Context, in *dto.GetCarsByDealerIDPaginatedRequest) (*dto.GetCarsByDealerIDPaginatedResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecaseImpl) SearchCarsPaginated(ctx context.Context, in *dto.SearchCarsPaginatedRequest) (*dto.SearchCarsPaginatedResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecaseImpl) GetCarsByField(ctx context.Context, in *dto.GetCarsByFieldRequest) (*dto.GetCarsByFieldResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecaseImpl) GetCarByDealerCount(ctx context.Context, in *dto.GetCarByDealerCountRequest) (*dto.GetCarByDealerCountResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecaseImpl) GetCarByDealershipCount(ctx context.Context, in *dto.GetCarByDealershipCountRequest) (*dto.GetCarByDealershipCountResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecaseImpl) GetCarByBrandCount(ctx context.Context, in *dto.GetCarByBrandCountRequest) (*dto.GetCarByCarBrandCountResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecaseImpl) DeleteCar(ctx context.Context, in *dto.DeleteCarRequest) (*dto.DeleteCarResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewUsecase(dbRepo ports.CarDatabaseRepo) ports.CarUsercase {
	return &usecaseImpl{
		dbRepo: dbRepo,
	}
}
