package car

import (
	"context"
	"garination.com/db/internal/core/dto"
	"garination.com/db/internal/platform/postgres"
	"github.com/jackc/pgx/v5/pgtype"
)

type CarRepository interface {
	CarByBrandCount(ctx context.Context, brandID string) (int64, error)
	CarByDealerCount(ctx context.Context, dealerID pgtype.Text) (int64, error)
	CarByDealershipCount(ctx context.Context, dealershipID pgtype.Text) (int64, error)
	CarsByDealerPaged(ctx context.Context, arg postgres.CarsByDealerPagedParams) ([]postgres.Car, error)
	CarsByDealershipPaged(ctx context.Context, arg postgres.CarsByDealershipPagedParams) ([]postgres.Car, error)
	CreateCarBrand(ctx context.Context, arg postgres.CreateCarBrandParams) (postgres.CarBrand, error)
	CreateCarImage(ctx context.Context, arg postgres.CreateCarImageParams) (string, error)
	DeleteCar(ctx context.Context, carID string) (string, error)
	DeleteCarBrand(ctx context.Context, brandID string) (postgres.CarBrand, error)
	DeleteCarImage(ctx context.Context, carImageID string) (postgres.CarImage, error)
	DeleteDealership(ctx context.Context, dealershipID string) (postgres.Dealership, error)
	DeleteExtraFeature(ctx context.Context, carExtraFeatureID string) (postgres.CarExtraFeature, error)
	GetCarBrandById(ctx context.Context, brandID string) (postgres.CarBrand, error)
	GetCarById(ctx context.Context, carID string) (postgres.Car, error)
	InsertExtraFeature(ctx context.Context, arg postgres.InsertExtraFeatureParams) (postgres.CarExtraFeature, error)
	InsertCar(ctx context.Context, arg postgres.InsertCarParams) (postgres.Car, error)
	ListCarBrandsPaged(ctx context.Context, arg postgres.ListCarBrandsPagedParams) ([]postgres.CarBrand, error)
	ListCarImagesForCar(ctx context.Context, carID string) ([]postgres.CarImage, error)
	ListCarsByBrandPaged(ctx context.Context, arg postgres.ListCarsByBrandPagedParams) ([]postgres.Car, error)
	ListCarsByDealerPaged(ctx context.Context, arg postgres.ListCarsByDealerPagedParams) ([]postgres.Car, error)
	ListCarsByDealershipPaged(ctx context.Context, arg postgres.ListCarsByDealershipPagedParams) ([]postgres.Car, error)
	ListCarsPaged(ctx context.Context, arg postgres.ListCarsPagedParams) ([]postgres.Car, error)
	ListExtraFeaturesForCar(ctx context.Context, carID string) ([]postgres.CarExtraFeature, error)
	SearchCarsPaged(ctx context.Context, arg postgres.SearchCarsPagedParams) ([]postgres.Car, error)
	UpdateExtraFeature(ctx context.Context, arg postgres.UpdateExtraFeatureParams) (string, error)
	UpdateCarImage(ctx context.Context, arg postgres.UpdateCarImageParams) (string, error)
	UpdateCar(ctx context.Context, arg postgres.UpdateCarParams) (string, error)
	UpdateCarBrand(ctx context.Context, arg postgres.UpdateCarBrandParams) (postgres.CarBrand, error)
	GetCarByField(ctx context.Context, req map[string]dto.CarFilter) ([]postgres.Car, error)
}
