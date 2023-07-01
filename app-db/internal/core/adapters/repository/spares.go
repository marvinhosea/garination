package repository

import (
	"context"
	"garination.com/db/internal/core/ports/spares"
	"garination.com/db/internal/platform/postgres"
)

type sparesRepoImpl struct {
	conn *postgres.Connection
}

func (s sparesRepoImpl) GetSparePartByID(ctx context.Context, id string) (postgres.SparePart, error) {
	return s.conn.Queries.GetSparePartById(ctx, id)
}

func (s sparesRepoImpl) DeleteSparePart(ctx context.Context, sparePartID string) (string, error) {
	return s.conn.Queries.DeleteSparePart(ctx, sparePartID)
}

func (s sparesRepoImpl) FilterSparePartsByBrand(ctx context.Context, arg postgres.FilterSparePartsByBrandParams) ([]postgres.SparePart, error) {
	return s.conn.Queries.FilterSparePartsByBrand(ctx, arg)
}

func (s sparesRepoImpl) FilterSparePartsByCategory(ctx context.Context, arg postgres.FilterSparePartsByCategoryParams) ([]postgres.SparePart, error) {
	return s.conn.Queries.FilterSparePartsByCategory(ctx, arg)
}

func (s sparesRepoImpl) FilterSparePartsByModel(ctx context.Context, arg postgres.FilterSparePartsByModelParams) ([]postgres.SparePart, error) {
	return s.conn.Queries.FilterSparePartsByModel(ctx, arg)
}

func (s sparesRepoImpl) GetSparePartImageById(ctx context.Context, sparePartImageID string) (postgres.SparePartImage, error) {
	return s.conn.Queries.GetSparePartImageById(ctx, sparePartImageID)
}

func (s sparesRepoImpl) InsertSparePart(ctx context.Context, arg postgres.InsertSparePartParams) (string, error) {
	return s.conn.Queries.InsertSparePart(ctx, arg)
}

func (s sparesRepoImpl) InsertSparePartImage(ctx context.Context, arg postgres.InsertSparePartImageParams) (string, error) {
	return s.conn.Queries.InsertSparePartImage(ctx, arg)
}

func (s sparesRepoImpl) ListSparePartImagesBySparePartPaged(ctx context.Context, arg postgres.ListSparePartImagesBySparePartPagedParams) ([]postgres.SparePartImage, error) {
	return s.conn.Queries.ListSparePartImagesBySparePartPaged(ctx, arg)
}

func (s sparesRepoImpl) ListSparePartImagesPaged(ctx context.Context, arg postgres.ListSparePartImagesPagedParams) ([]postgres.SparePartImage, error) {
	return s.conn.Queries.ListSparePartImagesPaged(ctx, arg)
}

func (s sparesRepoImpl) ListSparePartsByDealerPaged(ctx context.Context, arg postgres.ListSparePartsByDealerPagedParams) ([]postgres.SparePart, error) {
	return s.conn.Queries.ListSparePartsByDealerPaged(ctx, arg)
}

func (s sparesRepoImpl) ListSparePartsByDealershipPaged(ctx context.Context, arg postgres.ListSparePartsByDealershipPagedParams) ([]postgres.SparePart, error) {
	return s.conn.Queries.ListSparePartsByDealershipPaged(ctx, arg)
}

func (s sparesRepoImpl) ListSparePartsPaged(ctx context.Context, arg postgres.ListSparePartsPagedParams) ([]postgres.SparePart, error) {
	return s.conn.Queries.ListSparePartsPaged(ctx, arg)
}

func (s sparesRepoImpl) SearchSparePartsPaged(ctx context.Context, lower postgres.SearchSparePartsPagedParams) ([]postgres.SparePart, error) {
	return s.conn.Queries.SearchSparePartsPaged(ctx, lower)
}

func (s sparesRepoImpl) UpdateSparePart(ctx context.Context, arg postgres.UpdateSparePartParams) (string, error) {
	return s.conn.Queries.UpdateSparePart(ctx, arg)
}

func (s sparesRepoImpl) UpdateSparePartImage(ctx context.Context, arg postgres.UpdateSparePartImageParams) (string, error) {
	return s.conn.Queries.UpdateSparePartImage(ctx, arg)
}

func NewSparesRepository(conn *postgres.Connection) spares.SparesRepository {
	return &sparesRepoImpl{
		conn: conn,
	}
}
