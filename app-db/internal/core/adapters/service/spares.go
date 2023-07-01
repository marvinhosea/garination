package service

import (
	"context"
	"errors"
	"garination.com/db/internal/core/ports/spares"
	"garination.com/db/internal/platform/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"strings"
)

type sparesServiceImpl struct {
	sparesRepo spares.SparesRepository
}

func (s sparesServiceImpl) GetSparePartByID(ctx context.Context, id string) (postgres.SparePart, error) {
	return s.sparesRepo.GetSparePartByID(ctx, id)
}

func (s sparesServiceImpl) DeleteSparePart(ctx context.Context, sparePartID string) (string, error) {
	return s.sparesRepo.DeleteSparePart(ctx, sparePartID)
}

func (s sparesServiceImpl) FilterSparePartsByBrand(ctx context.Context, arg postgres.FilterSparePartsByBrandParams) ([]postgres.SparePart, error) {
	// escape % character
	arg.CarBrand.String = strings.ReplaceAll(arg.CarBrand.String, "%", "\\%")
	// add like % to the beginning and end of the string
	arg.CarBrand.String = "%" + arg.CarBrand.String + "%"
	return s.sparesRepo.FilterSparePartsByBrand(ctx, arg)
}

func (s sparesServiceImpl) FilterSparePartsByCategory(ctx context.Context, arg postgres.FilterSparePartsByCategoryParams) ([]postgres.SparePart, error) {
	// escape % character
	arg.Category.String = strings.ReplaceAll(arg.Category.String, "%", "\\%")
	// add like % to the beginning and end of the string
	arg.Category.String = "%" + arg.Category.String + "%"
	return s.sparesRepo.FilterSparePartsByCategory(ctx, arg)
}

func (s sparesServiceImpl) FilterSparePartsByModel(ctx context.Context, arg postgres.FilterSparePartsByModelParams) ([]postgres.SparePart, error) {
	// escape % character
	arg.CarModel.String = strings.ReplaceAll(arg.CarModel.String, "%", "\\%")
	// add like % to the beginning and end of the string
	arg.CarModel.String = "%" + arg.CarModel.String + "%"
	return s.sparesRepo.FilterSparePartsByModel(ctx, arg)
}

func (s sparesServiceImpl) GetSparePartImageById(ctx context.Context, sparePartImageID string) (postgres.SparePartImage, error) {
	return s.sparesRepo.GetSparePartImageById(ctx, sparePartImageID)
}

func (s sparesServiceImpl) InsertSparePart(ctx context.Context, arg postgres.InsertSparePartParams) (string, error) {
	arg.SparePartID = uuid.New().String()
	// check if dealer exists
	if arg.DealershipID.String == "" && arg.DealerID.String == "" {
		return "", errors.New("invalid_request: dealer or dealership id is required")
	}

	if arg.DealershipID.String != "" && arg.DealerID.String == "" {
		// make daler id null
		arg.DealerID = pgtype.Text{String: "", Valid: false}
	}

	if arg.DealerID.String != "" && arg.DealershipID.String == "" {
		// make dealership id null
		arg.DealershipID = pgtype.Text{String: "", Valid: false}
	}

	return s.sparesRepo.InsertSparePart(ctx, arg)
}

func (s sparesServiceImpl) InsertSparePartImage(ctx context.Context, arg postgres.InsertSparePartImageParams) (string, error) {
	arg.SparePartImageID = uuid.New().String()
	return s.sparesRepo.InsertSparePartImage(ctx, arg)
}

func (s sparesServiceImpl) ListSparePartImagesBySparePartPaged(ctx context.Context, arg postgres.ListSparePartImagesBySparePartPagedParams) ([]postgres.SparePartImage, error) {
	return s.sparesRepo.ListSparePartImagesBySparePartPaged(ctx, arg)
}

func (s sparesServiceImpl) ListSparePartImagesPaged(ctx context.Context, arg postgres.ListSparePartImagesPagedParams) ([]postgres.SparePartImage, error) {
	return s.sparesRepo.ListSparePartImagesPaged(ctx, arg)
}

func (s sparesServiceImpl) ListSparePartsByDealerPaged(ctx context.Context, arg postgres.ListSparePartsByDealerPagedParams) ([]postgres.SparePart, error) {
	return s.sparesRepo.ListSparePartsByDealerPaged(ctx, arg)
}

func (s sparesServiceImpl) ListSparePartsByDealershipPaged(ctx context.Context, arg postgres.ListSparePartsByDealershipPagedParams) ([]postgres.SparePart, error) {
	return s.sparesRepo.ListSparePartsByDealershipPaged(ctx, arg)
}

func (s sparesServiceImpl) ListSparePartsPaged(ctx context.Context, arg postgres.ListSparePartsPagedParams) ([]postgres.SparePart, error) {
	return s.sparesRepo.ListSparePartsPaged(ctx, arg)
}

func (s sparesServiceImpl) SearchSparePartsPaged(ctx context.Context, lower postgres.SearchSparePartsPagedParams) ([]postgres.SparePart, error) {
	// escape % character
	lower.Lower = strings.ReplaceAll(lower.Lower, "%", "\\%")
	lower.Lower = "%" + lower.Lower + "%"
	return s.sparesRepo.SearchSparePartsPaged(ctx, lower)
}

func (s sparesServiceImpl) UpdateSparePart(ctx context.Context, arg postgres.UpdateSparePartParams) (string, error) {
	return s.sparesRepo.UpdateSparePart(ctx, arg)
}

func (s sparesServiceImpl) UpdateSparePartImage(ctx context.Context, arg postgres.UpdateSparePartImageParams) (string, error) {
	return s.sparesRepo.UpdateSparePartImage(ctx, arg)
}

func NewSparesService(sparesRepo spares.SparesRepository) spares.SparesService {
	return &sparesServiceImpl{sparesRepo: sparesRepo}
}
