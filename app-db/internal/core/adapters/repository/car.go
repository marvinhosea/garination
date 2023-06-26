package repository

import (
	"context"
	"errors"
	"garination.com/db/internal/core/dto"
	"garination.com/db/internal/core/ports/car"
	"garination.com/db/internal/platform/postgres"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
)

type carRepository struct {
	conn *postgres.Connection
}

func (c carRepository) InsertExtraFeature(ctx context.Context, arg postgres.InsertExtraFeatureParams) (postgres.CarExtraFeature, error) {
	return c.conn.Queries.InsertExtraFeature(ctx, arg)
}

func (c carRepository) UpdateExtraFeature(ctx context.Context, arg postgres.UpdateExtraFeatureParams) (string, error) {
	return c.conn.Queries.UpdateExtraFeature(ctx, arg)
}

func (c carRepository) UpdateCarImage(ctx context.Context, arg postgres.UpdateCarImageParams) (string, error) {
	return c.conn.Queries.UpdateCarImage(ctx, arg)
}

func (c carRepository) GetCarByField(ctx context.Context, req map[string]dto.CarFilter) ([]postgres.Car, error) {
	log.Printf("GetCarByField: %+v", req)
	return nil, errors.New("not implemented")
}

func (c carRepository) CarByBrandCount(ctx context.Context, brandID string) (int64, error) {
	return c.conn.Queries.CarByBrandCount(ctx, brandID)
}

func (c carRepository) CarByDealerCount(ctx context.Context, dealerID pgtype.Text) (int64, error) {
	return c.conn.Queries.CarByDealerCount(ctx, dealerID)
}

func (c carRepository) CarByDealershipCount(ctx context.Context, dealershipID pgtype.Text) (int64, error) {
	return c.conn.Queries.CarByDealershipCount(ctx, dealershipID)
}

func (c carRepository) CarsByDealerPaged(ctx context.Context, arg postgres.CarsByDealerPagedParams) ([]postgres.Car, error) {
	return c.conn.Queries.CarsByDealerPaged(ctx, arg)
}

func (c carRepository) CarsByDealershipPaged(ctx context.Context, arg postgres.CarsByDealershipPagedParams) ([]postgres.Car, error) {
	return c.conn.Queries.CarsByDealershipPaged(ctx, arg)
}

func (c carRepository) CreateCarBrand(ctx context.Context, arg postgres.CreateCarBrandParams) (postgres.CarBrand, error) {
	return c.conn.Queries.CreateCarBrand(ctx, arg)
}

func (c carRepository) CreateCarImage(ctx context.Context, arg postgres.CreateCarImageParams) (string, error) {
	return c.conn.Queries.CreateCarImage(ctx, arg)
}

func (c carRepository) DeleteCar(ctx context.Context, carID string) (string, error) {
	return c.conn.Queries.DeleteCar(ctx, carID)
}

func (c carRepository) DeleteCarBrand(ctx context.Context, brandID string) (postgres.CarBrand, error) {
	return c.conn.Queries.DeleteCarBrand(ctx, brandID)
}

func (c carRepository) DeleteCarImage(ctx context.Context, carImageID string) (postgres.CarImage, error) {
	return c.conn.Queries.DeleteCarImage(ctx, carImageID)
}

func (c carRepository) DeleteDealership(ctx context.Context, dealershipID string) (postgres.Dealership, error) {
	return c.conn.Queries.DeleteDealership(ctx, dealershipID)
}

func (c carRepository) DeleteExtraFeature(ctx context.Context, carExtraFeatureID string) (postgres.CarExtraFeature, error) {
	return c.conn.Queries.DeleteExtraFeature(ctx, carExtraFeatureID)
}

func (c carRepository) GetCarBrandById(ctx context.Context, brandID string) (postgres.CarBrand, error) {
	return c.conn.Queries.GetCarBrandById(ctx, brandID)
}

func (c carRepository) GetCarById(ctx context.Context, carID string) (postgres.Car, error) {
	return c.conn.Queries.GetCarById(ctx, carID)
}

func (c carRepository) InsertCar(ctx context.Context, arg postgres.InsertCarParams) (postgres.Car, error) {
	return c.conn.Queries.InsertCar(ctx, arg)
}

func (c carRepository) ListCarBrandsPaged(ctx context.Context, arg postgres.ListCarBrandsPagedParams) ([]postgres.CarBrand, error) {
	return c.conn.Queries.ListCarBrandsPaged(ctx, arg)
}

func (c carRepository) ListCarImagesForCar(ctx context.Context, carID string) ([]postgres.CarImage, error) {
	return c.conn.Queries.ListCarImagesForCar(ctx, carID)
}

func (c carRepository) ListCarsByBrandPaged(ctx context.Context, arg postgres.ListCarsByBrandPagedParams) ([]postgres.Car, error) {
	return c.conn.Queries.ListCarsByBrandPaged(ctx, arg)
}

func (c carRepository) ListCarsByDealerPaged(ctx context.Context, arg postgres.ListCarsByDealerPagedParams) ([]postgres.Car, error) {
	return c.conn.Queries.ListCarsByDealerPaged(ctx, arg)
}

func (c carRepository) ListCarsByDealershipPaged(ctx context.Context, arg postgres.ListCarsByDealershipPagedParams) ([]postgres.Car, error) {
	return c.conn.Queries.ListCarsByDealershipPaged(ctx, arg)
}

func (c carRepository) ListCarsPaged(ctx context.Context, arg postgres.ListCarsPagedParams) ([]postgres.Car, error) {
	return c.conn.Queries.ListCarsPaged(ctx, arg)
}

func (c carRepository) ListExtraFeaturesForCar(ctx context.Context, carID string) ([]postgres.CarExtraFeature, error) {
	return c.conn.Queries.ListExtraFeaturesForCar(ctx, carID)
}

func (c carRepository) SearchCarsPaged(ctx context.Context, arg postgres.SearchCarsPagedParams) ([]postgres.Car, error) {
	return c.conn.Queries.SearchCarsPaged(ctx, arg)
}

func (c carRepository) UpdateCar(ctx context.Context, arg postgres.UpdateCarParams) (string, error) {
	return c.conn.Queries.UpdateCar(ctx, arg)
}

func (c carRepository) UpdateCarBrand(ctx context.Context, arg postgres.UpdateCarBrandParams) (postgres.CarBrand, error) {
	return c.conn.Queries.UpdateCarBrand(ctx, arg)
}

func NewCarRepository(conn *postgres.Connection) car.CarRepository {
	return &carRepository{
		conn: conn,
	}
}
