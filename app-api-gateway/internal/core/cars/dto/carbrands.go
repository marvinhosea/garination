package dto

import "github.com/google/uuid"

type CreateCarBrandRequest struct {
	Name    string `json:"name,omitempty"`
	LogoURL string `json:"logo_url,omitempty"`
}

type CreateCarBrandResponse struct {
	BrandID string `json:"brand_id,omitempty"`
}

type UpdateCarBrandRequest struct {
	BrandID string `json:"brand_id,omitempty"`
	Name    string `json:"name,omitempty"`
	LogoURL string `json:"logo_url,omitempty"`
}

type UpdateCarBrandResponse struct {
	BrandID string `json:"brand_id,omitempty"`
}

type GetCarBrandByIDRequest struct {
	BrandID string `json:"brand_id,omitempty"`
}

func (r *GetCarBrandByIDRequest) Validate() error {
	_, err := uuid.Parse(r.BrandID)
	if err != nil {
		return err
	}
	return nil
}

type GetCarBrandByIDResponse struct {
	BrandID   string `json:"brand_id,omitempty"`
	Name      string `json:"name,omitempty"`
	LogoURL   string `json:"logo_url,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type GetCarBrandsPaginatedRequest struct {
	Offset int32 `json:"offset,omitempty"`
	Limit  int32 `json:"limit,omitempty"`
}

type GetCarBrandsPaginatedResponse struct {
	Brands []GetCarBrandByIDResponse `json:"brands,omitempty"`
	Total  int32                     `json:"total,omitempty"`
}

type DeleteCarBrandRequest struct {
	BrandID string `json:"brand_id,omitempty"`
}

type DeleteCarBrandResponse struct {
	BrandID string `json:"brand_id,omitempty"`
}
