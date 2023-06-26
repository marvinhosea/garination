package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type IQuery interface {
	WithTx(tx pgx.Tx) *Queries
	CarByBrandCount(ctx context.Context, brandID string) (int64, error)
	CarByDealerCount(ctx context.Context, dealerID pgtype.Text) (int64, error)
	CarByDealershipCount(ctx context.Context, dealershipID pgtype.Text) (int64, error)
	CarsByDealerPaged(ctx context.Context, arg CarsByDealerPagedParams) ([]Car, error)
	CarsByDealershipPaged(ctx context.Context, arg CarsByDealershipPagedParams) ([]Car, error)
	CreateCarBrand(ctx context.Context, arg CreateCarBrandParams) (CarBrand, error)
	CreateCarImage(ctx context.Context, arg CreateCarImageParams) ([]string, error)
	DeleteCar(ctx context.Context, carID string) (string, error)
	DeleteCarBrand(ctx context.Context, brandID string) (CarBrand, error)
	DeleteCarImage(ctx context.Context, carImageID string) (CarImage, error)
	DeleteDealership(ctx context.Context, dealershipID string) (Dealership, error)
	DeleteExtraFeature(ctx context.Context, carExtraFeatureID string) (CarExtraFeature, error)
	GetCarBrandById(ctx context.Context, brandID string) (CarBrand, error)
	GetCarById(ctx context.Context, carID string) (Car, error)
	GetDealershipById(ctx context.Context, dealershipID string) (Dealership, error)
	GetUserDealership(ctx context.Context, userID string) (Dealership, error)
	GetUserMeta(ctx context.Context, userID string) (UserMetum, error)
	InsertDealership(ctx context.Context, arg InsertDealershipParams) (Dealership, error)
	InsertExtraFeature(ctx context.Context, arg InsertExtraFeatureParams) (CarExtraFeature, error)
	InsertUserMeta(ctx context.Context, arg InsertUserMetaParams) (UserMetum, error)
	InsertUserMetaWithoutDealership(ctx context.Context, arg InsertUserMetaWithoutDealershipParams) (UserMetum, error)
	UpdateDealership(ctx context.Context, arg UpdateDealershipParams) (Dealership, error)
	UpdateExtraFeature(ctx context.Context, arg UpdateExtraFeatureParams) ([]string, error)
	UpdateUserDealership(ctx context.Context, arg UpdateUserDealershipParams) (UserMetum, error)
	UpdateUserMeta(ctx context.Context, arg UpdateUserMetaParams) (UserMetum, error)
}
