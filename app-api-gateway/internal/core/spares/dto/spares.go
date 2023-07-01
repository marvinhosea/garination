package dto

import (
	"errors"
	"garination.com/gateway/internal/core/spares/model"
	"github.com/google/uuid"
	"mime/multipart"
)

type DeleteSparePartRequest struct {
	ID string `form:"spare_part_id" json:"spare_part_id" binding:"required"`
}

func (r DeleteSparePartRequest) Validate() error {
	// validate uuid
	// validate uuid
	if _, err := uuid.Parse(r.ID); err != nil {
		return errors.New("invalid spare part ID")
	}

	return nil
}

type FilterSparePartByBrandPaginatedRequest struct {
	Query  string `form:"query" json:"query" binding:"required"`
	Offset int32  `form:"offset" json:"offset"`
	Limit  int32  `form:"limit" json:"limit"`
}

func (r *FilterSparePartByBrandPaginatedRequest) Validate() error {
	if r.Limit <= 0 {
		r.Limit = 25
	}

	if r.Offset <= 0 {
		r.Offset = 0
	}

	return nil
}

type FilterSparePartByCategoryPaginatedRequest struct {
	Query  string `form:"query" json:"query" binding:"required"`
	Offset int32  `form:"offset" json:"offset"`
	Limit  int32  `form:"limit" json:"limit"`
}

func (r *FilterSparePartByCategoryPaginatedRequest) Validate() error {
	if r.Limit <= 0 {
		r.Limit = 25
	}

	if r.Offset <= 0 {
		r.Offset = 0
	}

	return nil
}

type FilterSparePartByCarModelPaginatedRequest struct {
	Query  string `form:"query" json:"query" binding:"required"`
	Offset int32  `form:"offset" json:"offset"`
	Limit  int32  `form:"limit" json:"limit"`
}

func (r *FilterSparePartByCarModelPaginatedRequest) Validate() error {
	if r.Limit <= 0 {
		r.Limit = 25
	}

	if r.Offset <= 0 {
		r.Offset = 0
	}

	return nil
}

type GetSparePartImageByIDRequest struct {
	ID string `form:"spare_part_image_id" json:"spare_part_image_id" binding:"required"`
}

func (r GetSparePartImageByIDRequest) Validate() error {
	if _, err := uuid.Parse(r.ID); err != nil {
		return errors.New("invalid spare part image ID")
	}
	return nil
}

type InsertSparePartRequest struct {
	Name                string   `json:"name"`
	Description         string   `json:"description"`
	Price               float64  `json:"price"`
	Used                bool     `json:"used"`
	CarModel            string   `json:"car_model"`
	CarBrand            string   `json:"car_brand"`
	OtherCompatibleCars []string `json:"other_compatible_cars"`
	CarYear             int32    `json:"car_year"`
	IsUniversal         bool     `json:"is_universal"`
	Category            string   `json:"category"`
	PartNumber          string   `json:"part_number"`
	DealershipID        string   `json:"dealership_id"`
	DealerID            string   `json:"dealer_id"`
}

func (r *InsertSparePartRequest) Validate() error {
	if r.Name == "" {
		return errors.New("missing name")
	}

	if r.Description == "" {
		return errors.New("missing description")
	}

	if r.Price <= 0 {
		return errors.New("missing price")
	}

	if r.CarModel == "" {
		return errors.New("missing car model")
	}

	if r.CarBrand == "" {
		return errors.New("missing car brand")
	}

	if r.CarYear <= 0 {
		return errors.New("missing car year")
	}

	if r.Category == "" {
		return errors.New("missing category")
	}

	if r.DealershipID == "" && r.DealerID == "" {
		return errors.New("missing dealership ID or dealer ID")
	}

	if r.DealershipID != "" && r.DealerID != "" {
		return errors.New("cannot have both dealership ID and dealer ID")
	}

	if r.DealershipID != "" {
		if _, err := uuid.Parse(r.DealershipID); err != nil {
			return errors.New("invalid dealership ID")
		}
	}

	if r.DealerID != "" {
		if _, err := uuid.Parse(r.DealerID); err != nil {
			return errors.New("invalid dealer ID")
		}
	}

	return nil
}

type InsertSparePartImageRequest struct {
	SparePartImageID string                  `form:"spare_part_image_id"`
	SparePartID      string                  `form:"spare_part_id"`
	Images           []*multipart.FileHeader `form:"image[]"`
}

func (r *InsertSparePartImageRequest) Validate() error {
	if r.SparePartID == "" {
		return errors.New("missing spare part ID")
	}

	if _, err := uuid.Parse(r.SparePartID); err != nil {
		return errors.New("invalid spare part ID")
	}

	if len(r.Images) == 0 {
		return errors.New("missing image(s)")
	}

	return nil
}

type ListSparePartImagesBySparePartPagedRequest struct {
	SparePartId string `form:"spare_part_id,omitempty"`
	Offset      int32  `form:"offset,omitempty"`
	Limit       int32  `form:"limit,omitempty"`
}

func (r *ListSparePartImagesBySparePartPagedRequest) Validate() error {
	if r.Limit <= 0 {
		r.Limit = 25
	}

	if r.Offset <= 0 {
		r.Offset = 0
	}

	if r.SparePartId == "" {
		return errors.New("missing spare part ID")
	}

	if _, err := uuid.Parse(r.SparePartId); err != nil {
		return errors.New("invalid spare part ID")
	}

	return nil
}

type ListSparePartImagesPagedRequest struct {
	Offset int32 `form:"offset,omitempty"`
	Limit  int32 `form:"limit,omitempty"`
}

func (r *ListSparePartImagesPagedRequest) Validate() error {
	if r.Limit <= 0 {
		r.Limit = 25
	}

	if r.Offset <= 0 {
		r.Offset = 0
	}

	return nil
}

type ListSparePartsByDealerPagedRequest struct {
	DealerId string `form:"dealer_id,omitempty"`
	Offset   int32  `form:"offset,omitempty"`
	Limit    int32  `form:"limit,omitempty"`
}

func (r *ListSparePartsByDealerPagedRequest) Validate() error {
	if r.Limit <= 0 {
		r.Limit = 25
	}

	if r.Offset <= 0 {
		r.Offset = 0
	}

	if r.DealerId == "" {
		return errors.New("missing dealer ID")
	}

	if _, err := uuid.Parse(r.DealerId); err != nil {
		return errors.New("invalid dealer ID")
	}

	return nil

}

type ListSparePartsByDealershipPagedRequest struct {
	DealershipId string `form:"dealership_id,omitempty"`
	Offset       int32  `form:"offset,omitempty"`
	Limit        int32  `form:"limit,omitempty"`
}

func (r *ListSparePartsByDealershipPagedRequest) Validate() error {
	if r.Limit <= 0 {
		r.Limit = 25
	}

	if r.Offset <= 0 {
		r.Offset = 0
	}

	if r.DealershipId == "" {
		return errors.New("missing dealership ID")
	}

	if _, err := uuid.Parse(r.DealershipId); err != nil {
		return errors.New("invalid dealership ID")
	}

	return nil
}

type ListSparePartsPagedRequest struct {
	Offset int32 `form:"offset,omitempty"`
	Limit  int32 `form:"limit,omitempty"`
}

func (r *ListSparePartsPagedRequest) Validate() error {
	if r.Limit <= 0 {
		r.Limit = 25
	}

	if r.Offset <= 0 {
		r.Offset = 0
	}

	return nil
}

type SearchSparePartsPagedRequest struct {
	Query  string `form:"query,omitempty"`
	Offset int32  `form:"offset,omitempty"`
	Limit  int32  `form:"limit,omitempty"`
}

func (r *SearchSparePartsPagedRequest) Validate() error {
	if r.Limit <= 0 {
		r.Limit = 25
	}

	if r.Offset <= 0 {
		r.Offset = 0
	}

	if r.Query == "" {
		return errors.New("missing query")
	}

	return nil
}

type UpdateSparePartRequest struct {
	SparePartId         string   `json:"spare_part_id,omitempty"`
	Name                string   `json:"name,omitempty"`
	Description         string   `json:"description,omitempty"`
	Price               int32    `json:"price,omitempty"`
	Used                bool     `json:"used,omitempty"`
	CarModel            string   `json:"car_model,omitempty"`
	CarBrand            string   `json:"car_brand,omitempty"`
	OtherCompatibleCars []string `json:"other_compatible_cars,omitempty"`
	CarYear             int32    `json:"car_year,omitempty"`
	IsUniversal         bool     `json:"is_universal,omitempty"`
	Category            string   `json:"category,omitempty"`
	PartNumber          string   `json:"part_number,omitempty"`
}

func (r *UpdateSparePartRequest) Validate() error {
	if r.SparePartId == "" {
		return errors.New("missing spare part ID")
	}

	if _, err := uuid.Parse(r.SparePartId); err != nil {
		return errors.New("invalid spare part ID")
	}

	if r.Name == "" {
		return errors.New("missing name")
	}

	if r.Price <= 0 {
		return errors.New("missing price")
	}

	if r.CarModel == "" {
		return errors.New("missing car model")
	}

	if r.CarBrand == "" {
		return errors.New("missing car brand")
	}

	if r.CarYear <= 0 {
		return errors.New("missing car year")
	}

	if r.Category == "" {
		return errors.New("missing category")
	}

	return nil

}

type UpdateSparePartImageRequest struct {
	SparePartId string                `form:"spare_part_id,omitempty"`
	Image       *multipart.FileHeader `form:"image[],omitempty"`
}

func (r *UpdateSparePartImageRequest) Validate() error {
	if r.SparePartId == "" {
		return errors.New("missing spare part ID")
	}

	if _, err := uuid.Parse(r.SparePartId); err != nil {
		return errors.New("invalid spare part ID")
	}

	return nil
}

// 	GetSparePartByID(ctx context.Context, request *dto.GetSparePartByIDRequest) (*dto.GetSparePartByIDResponse, error)

type GetSparePartByIDRequest struct {
	ID string `form:"spare_part_id" json:"spare_part_id" binding:"required"`
}

func (r GetSparePartByIDRequest) Validate() error {
	if r.ID == "" {
		return errors.New("missing spare part ID")
	}

	if _, err := uuid.Parse(r.ID); err != nil {
		return errors.New("invalid spare part ID")
	}

	return nil
}

type GetSparePartByIDResponse struct {
	SparePart model.SparePart        `json:"spare_part,omitempty"`
	Images    []model.SparePartImage `json:"images,omitempty"`
}

type DeleteSparePartResponse struct {
	SparePartId string `json:"spare_part_id,omitempty"`
}

type FilterSparePartByBrandPaginatedResponse struct {
	SpareParts      []model.SparePart `json:"spare_parts,omitempty"`
	Size            int32             `json:"size,omitempty"`
	TotalSpareParts int32             `json:"total_spare_parts,omitempty"`
}

type FilterSparePartByCategoryPaginatedResponse struct {
	SpareParts      []model.SparePart `json:"spare_parts,omitempty"`
	Size            int32             `json:"size,omitempty"`
	TotalSpareParts int32             `json:"total_spare_parts,omitempty"`
}

type FilterSparePartByCarModelPaginatedResponse struct {
	SpareParts      []model.SparePart `json:"spare_parts,omitempty"`
	Size            int32             `json:"size,omitempty"`
	TotalSpareParts int32             `json:"total_spare_parts,omitempty"`
}

type GetSparePartImageByIDResponse struct {
	SparePartImage model.SparePartImage `json:"spare_part_image,omitempty"`
}

type InsertSparePartResponse struct {
	SparePartId string `json:"spare_part_id,omitempty"`
}

type InsertSparePartImageResponse struct {
	Images []model.SparePartImage `json:"images,omitempty"`
}

type ListSparePartImagesBySparePartPagedResponse struct {
	SparePartImages []model.SparePartImage `json:"spare_part_images,omitempty"`
	Size            int32                  `json:"size,omitempty"`
	TotalSpareParts int32                  `json:"total_spare_parts,omitempty"`
}

type ListSparePartImagesPagedResponse struct {
	SparePartImages []model.SparePartImage `json:"spare_part_images,omitempty"`
	Size            int32                  `json:"size,omitempty"`
	TotalSpareParts int32                  `json:"total_spare_parts,omitempty"`
}

type ListSparePartsByDealerPagedResponse struct {
	SpareParts      []model.SparePart `json:"spare_parts,omitempty"`
	Size            int32             `json:"size,omitempty"`
	TotalSpareParts int32             `json:"total_spare_parts,omitempty"`
}

type ListSparePartsByDealershipPagedResponse struct {
	SpareParts      []model.SparePart `json:"spare_parts,omitempty"`
	Size            int32             `json:"size,omitempty"`
	TotalSpareParts int32             `json:"total_spare_parts,omitempty"`
}

type ListSparePartsPagedResponse struct {
	SpareParts      []model.SparePart `json:"spare_parts,omitempty"`
	Size            int32             `json:"size,omitempty"`
	TotalSpareParts int32             `json:"total_spare_parts,omitempty"`
}

type SearchSparePartsPagedResponse struct {
	SpareParts      []model.SparePart `json:"spare_parts,omitempty"`
	Size            int32             `json:"size,omitempty"`
	TotalSpareParts int32             `json:"total_spare_parts,omitempty"`
}

type UpdateSparePartResponse struct {
	SparePartId string `json:"spare_part_id,omitempty"`
}

type UpdateSparePartImageResponse struct {
	SparePartImageId string `json:"spare_part_image_id,omitempty"`
}
