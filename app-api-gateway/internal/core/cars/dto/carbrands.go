package dto

import (
	"errors"
	"github.com/google/uuid"
	"mime/multipart"
	"time"
)

type CreateCarBrandRequest struct {
	Name    string                `form:"name"`
	Logo    *multipart.FileHeader `form:"logo"`
	Country string                `form:"country"`
}

func (r *CreateCarBrandRequest) Validate() error {
	if r.Name == "" {
		return errors.New("name is required")
	}

	if r.Logo == nil {
		return errors.New("logo is required")
	}

	if r.Country == "" {
		return errors.New("country is required")
	}
	return nil
}

type CreateCarBrandResponse struct {
	BrandID string `json:"brand_id"`
	Logo    string `json:"logo"`
}

type UpdateCarBrandRequest struct {
	BrandID string `json:"brand_id,omitempty"`
	Name    string `json:"name,omitempty"`
	LogoURL string `json:"logo_url,omitempty"`
}

func (r *UpdateCarBrandRequest) Validate() error {
	if r.BrandID == "" {
		return errors.New("brand_id is required")
	}

	if r.Name == "" {
		return errors.New("name is required")
	}

	if r.LogoURL == "" {
		return errors.New("logo_url is required")
	}

	if _, err := uuid.Parse(r.BrandID); err != nil {
		return errors.New("brand_id is invalid")
	}

	return nil
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
	BrandID   string    `json:"brand_id,omitempty"`
	Name      string    `json:"name,omitempty"`
	LogoURL   string    `json:"logo_url,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type GetCarBrandsPaginatedRequest struct {
	Offset int32 `form:"offset,omitempty"`
	Limit  int32 `form:"limit,omitempty"`
}

func (r *GetCarBrandsPaginatedRequest) Validate() error {
	if r.Offset < 0 {
		return errors.New("offset must be greater than 0")
	}

	if r.Limit < 10 {
		return errors.New("limit must be greater than 10 or more")
	}

	return nil
}

type GetCarBrandsPaginatedResponse struct {
	Brands []GetCarBrandByIDResponse `json:"brands,omitempty"`
	Total  int32                     `json:"total,omitempty"`
}

type DeleteCarBrandRequest struct {
	BrandID string `json:"brand_id,omitempty"`
}

func (r *DeleteCarBrandRequest) Validate() error {
	_, err := uuid.Parse(r.BrandID)
	if err != nil {
		return err
	}
	return nil
}

type DeleteCarBrandResponse struct {
	BrandID string `json:"brand_id,omitempty"`
}
