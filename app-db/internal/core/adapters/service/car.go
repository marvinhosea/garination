package service

import (
	"context"
	"errors"
	"garination.com/db/internal/core/dto"
	"garination.com/db/internal/core/ports/car"
	"garination.com/db/internal/core/ports/dealership"
	"garination.com/db/internal/core/ports/user"
	"garination.com/db/internal/platform/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type carService struct {
	carRepo        car.CarRepository
	userMetaRepo   user.UserRepository
	dealershipRepo dealership.DealershipRepository
}

func (c carService) UpdateExtraFeature(ctx context.Context, arg postgres.UpdateExtraFeatureParams) (string, error) {
	return c.carRepo.UpdateExtraFeature(ctx, arg)
}

func (c carService) UpdateCarImage(ctx context.Context, arg postgres.UpdateCarImageParams) (string, error) {
	return c.carRepo.UpdateCarImage(ctx, arg)
}

func (c carService) InsertExtraFeature(ctx context.Context, arg postgres.InsertExtraFeatureParams) (postgres.CarExtraFeature, error) {
	return c.carRepo.InsertExtraFeature(ctx, arg)
}

func (c carService) GetCarByField(ctx context.Context, req map[string]dto.CarFilter) ([]postgres.Car, error) {
	return c.carRepo.GetCarByField(ctx, req)
}

func (c carService) CarByBrandCount(ctx context.Context, brandID string) (int64, error) {
	return c.carRepo.CarByBrandCount(ctx, brandID)
}

func (c carService) CarByDealerCount(ctx context.Context, dealerID pgtype.Text) (int64, error) {
	return c.carRepo.CarByDealerCount(ctx, dealerID)
}

func (c carService) CarByDealershipCount(ctx context.Context, dealershipID pgtype.Text) (int64, error) {
	return c.carRepo.CarByDealershipCount(ctx, dealershipID)
}

func (c carService) CarsByDealerPaged(ctx context.Context, arg postgres.CarsByDealerPagedParams) ([]postgres.Car, error) {
	return c.carRepo.CarsByDealerPaged(ctx, arg)
}

func (c carService) CarsByDealershipPaged(ctx context.Context, arg postgres.CarsByDealershipPagedParams) ([]postgres.Car, error) {
	return c.carRepo.CarsByDealershipPaged(ctx, arg)
}

func (c carService) CreateCarBrand(ctx context.Context, arg postgres.CreateCarBrandParams) (postgres.CarBrand, error) {
	arg.BrandID = uuid.New().String()
	return c.carRepo.CreateCarBrand(ctx, arg)
}

func (c carService) CreateCarImage(ctx context.Context, arg dto.CarImagesParam) ([]string, error) {

	var carImages []string

	for _, image := range arg.CarImages {
		image.CarImageID = uuid.New().String()
		_, err := c.carRepo.GetCarById(ctx, image.CarID)
		if err != nil {
			return nil, err
		}

		req := postgres.CreateCarImageParams{
			CarImageID: image.CarImageID,
			CarID:      image.CarID,
			ImageUrl:   pgtype.Text{String: image.ImageUrl, Valid: true},
		}

		res, err := c.carRepo.CreateCarImage(ctx, req)
		if err != nil {
			return nil, err
		}

		carImages = append(carImages, res)
	}

	// create car image
	return carImages, nil
}

func (c carService) DeleteCar(ctx context.Context, carID string) (string, error) {
	return c.carRepo.DeleteCar(ctx, carID)
}

func (c carService) DeleteCarBrand(ctx context.Context, brandID string) (postgres.CarBrand, error) {
	return c.carRepo.DeleteCarBrand(ctx, brandID)
}

func (c carService) DeleteCarImage(ctx context.Context, carImageID string) (postgres.CarImage, error) {
	return c.carRepo.DeleteCarImage(ctx, carImageID)
}

func (c carService) DeleteExtraFeature(ctx context.Context, carExtraFeatureID string) (postgres.CarExtraFeature, error) {
	return c.carRepo.DeleteExtraFeature(ctx, carExtraFeatureID)
}

func (c carService) GetCarBrandById(ctx context.Context, brandID string) (postgres.CarBrand, error) {
	return c.carRepo.GetCarBrandById(ctx, brandID)
}

func (c carService) GetCarById(ctx context.Context, carID string) (postgres.Car, error) {
	return c.carRepo.GetCarById(ctx, carID)
}

func (c carService) InsertCar(ctx context.Context, arg postgres.InsertCarParams) (postgres.Car, error) {
	arg.CarID = uuid.New().String()
	// check if dealer exists
	if arg.DealershipID.String == "" && arg.DealerID.String == "" {
		return postgres.Car{}, errors.New("invalid_request: dealer or dealership id is required")
	}

	if arg.DealershipID.String != "" {
		_, err := c.dealershipRepo.GetDealershipByID(ctx, arg.DealershipID.String)
		if err != nil {
			return postgres.Car{}, err
		}
	}

	if arg.DealerID.String != "" {
		_, err := c.userMetaRepo.GetUserMeta(ctx, arg.DealerID.String)
		if err != nil {
			return postgres.Car{}, err
		}
	}

	return c.carRepo.InsertCar(ctx, arg)
}

func (c carService) ListCarBrandsPaged(ctx context.Context, arg postgres.ListCarBrandsPagedParams) ([]postgres.CarBrand, error) {
	return c.carRepo.ListCarBrandsPaged(ctx, arg)
}

func (c carService) ListCarImagesForCar(ctx context.Context, carID string) ([]postgres.CarImage, error) {
	return c.carRepo.ListCarImagesForCar(ctx, carID)
}

func (c carService) ListCarsByBrandPaged(ctx context.Context, arg postgres.ListCarsByBrandPagedParams) ([]postgres.Car, error) {
	return c.carRepo.ListCarsByBrandPaged(ctx, arg)
}

func (c carService) ListCarsByDealerPaged(ctx context.Context, arg postgres.ListCarsByDealerPagedParams) ([]postgres.Car, error) {
	return c.carRepo.ListCarsByDealerPaged(ctx, arg)
}

func (c carService) ListCarsByDealershipPaged(ctx context.Context, arg postgres.ListCarsByDealershipPagedParams) ([]postgres.Car, error) {
	return c.carRepo.ListCarsByDealershipPaged(ctx, arg)
}

func (c carService) ListCarsPaged(ctx context.Context, arg postgres.ListCarsPagedParams) ([]postgres.Car, error) {
	return c.carRepo.ListCarsPaged(ctx, arg)
}

func (c carService) ListExtraFeaturesForCar(ctx context.Context, carID string) ([]postgres.CarExtraFeature, error) {
	return c.carRepo.ListExtraFeaturesForCar(ctx, carID)
}

func (c carService) SearchCarsPaged(ctx context.Context, arg postgres.SearchCarsPagedParams) ([]postgres.Car, error) {
	return c.carRepo.SearchCarsPaged(ctx, arg)
}

func (c carService) UpdateCar(ctx context.Context, arg postgres.UpdateCarParams) (string, error) {
	return c.carRepo.UpdateCar(ctx, arg)
}

func (c carService) UpdateCarBrand(ctx context.Context, arg postgres.UpdateCarBrandParams) (postgres.CarBrand, error) {
	return c.carRepo.UpdateCarBrand(ctx, arg)
}

func NewCarService(carRepo car.CarRepository, userMetaRepo user.UserRepository, dealershipRepo dealership.DealershipRepository) car.CarService {
	return &carService{
		carRepo:        carRepo,
		userMetaRepo:   userMetaRepo,
		dealershipRepo: dealershipRepo,
	}
}
