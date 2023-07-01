package postgres

import (
	"context"
)

type IQuery interface {
	DeleteSparePart(ctx context.Context, sparePartID string) (string, error)
	FilterSparePartsByBrand(ctx context.Context, arg FilterSparePartsByBrandParams) ([]SparePart, error)
	FilterSparePartsByCategory(ctx context.Context, arg FilterSparePartsByCategoryParams) ([]SparePart, error)
	FilterSparePartsByModel(ctx context.Context, arg FilterSparePartsByModelParams) ([]SparePart, error)
	GetSparePartImageById(ctx context.Context, sparePartImageID string) (SparePartImage, error)
	InsertSparePart(ctx context.Context, arg InsertSparePartParams) (string, error)
	InsertSparePartImage(ctx context.Context, arg InsertSparePartImageParams) (string, error)
	ListSparePartImagesBySparePartPaged(ctx context.Context, arg ListSparePartImagesBySparePartPagedParams) ([]SparePartImage, error)
	ListSparePartImagesPaged(ctx context.Context, arg ListSparePartImagesPagedParams) ([]SparePartImage, error)
	ListSparePartsByDealerPaged(ctx context.Context, arg ListSparePartsByDealerPagedParams) ([]SparePart, error)
	ListSparePartsByDealershipPaged(ctx context.Context, arg ListSparePartsByDealershipPagedParams) ([]SparePart, error)
	ListSparePartsPaged(ctx context.Context, arg ListSparePartsPagedParams) ([]SparePart, error)
	SearchSparePartsPaged(ctx context.Context, lower SearchCarsPagedParams) ([]SparePart, error)
	UpdateSparePart(ctx context.Context, arg UpdateSparePartParams) (string, error)
	UpdateSparePartImage(ctx context.Context, arg UpdateSparePartImageParams) (string, error)
}
