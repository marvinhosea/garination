package spares

import (
	"context"
	"garination.com/db/internal/platform/postgres"
)

type SparesRepository interface {
	DeleteSparePart(ctx context.Context, sparePartID string) (string, error)
	FilterSparePartsByBrand(ctx context.Context, arg postgres.FilterSparePartsByBrandParams) ([]postgres.SparePart, error)
	FilterSparePartsByCategory(ctx context.Context, arg postgres.FilterSparePartsByCategoryParams) ([]postgres.SparePart, error)
	FilterSparePartsByModel(ctx context.Context, arg postgres.FilterSparePartsByModelParams) ([]postgres.SparePart, error)
	GetSparePartImageById(ctx context.Context, sparePartImageID string) (postgres.SparePartImage, error)
	InsertSparePart(ctx context.Context, arg postgres.InsertSparePartParams) (string, error)
	InsertSparePartImage(ctx context.Context, arg postgres.InsertSparePartImageParams) (string, error)
	ListSparePartImagesBySparePartPaged(ctx context.Context, arg postgres.ListSparePartImagesBySparePartPagedParams) ([]postgres.SparePartImage, error)
	ListSparePartImagesPaged(ctx context.Context, arg postgres.ListSparePartImagesPagedParams) ([]postgres.SparePartImage, error)
	ListSparePartsByDealerPaged(ctx context.Context, arg postgres.ListSparePartsByDealerPagedParams) ([]postgres.SparePart, error)
	ListSparePartsByDealershipPaged(ctx context.Context, arg postgres.ListSparePartsByDealershipPagedParams) ([]postgres.SparePart, error)
	ListSparePartsPaged(ctx context.Context, arg postgres.ListSparePartsPagedParams) ([]postgres.SparePart, error)
	SearchSparePartsPaged(ctx context.Context, lower postgres.SearchSparePartsPagedParams) ([]postgres.SparePart, error)
	UpdateSparePart(ctx context.Context, arg postgres.UpdateSparePartParams) (string, error)
	UpdateSparePartImage(ctx context.Context, arg postgres.UpdateSparePartImageParams) (string, error)
	GetSparePartByID(ctx context.Context, id string) (postgres.SparePart, error)
}
