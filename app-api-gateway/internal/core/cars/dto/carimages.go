package dto

import (
	"errors"
	"garination.com/gateway/internal/core/cars/model"
	"github.com/google/uuid"
	"mime/multipart"
)

type CreateCarImageRequest struct {
	CarID  string                  `form:"car_id,omitempty" binding:"required"`
	Images []*multipart.FileHeader `form:"image[],omitempty" binding:"required"`
}

func (r *CreateCarImageRequest) Validate() error {
	if r.CarID == "" {
		return errors.New("car_id is required")
	}

	if _, err := uuid.Parse(r.CarID); err != nil {
		return errors.New("car_id is invalid")
	}

	if len(r.Images) == 0 {
		return errors.New("images is required")
	}

	return nil
}

type CreateCarImageResponse struct {
	ImageID []string `json:"image_ids,omitempty"`
}

type UpdateCarImageRequest struct {
	ImageID string                `json:"image_id,omitempty"`
	URL     string                `json:"-,omitempty"`
	Image   *multipart.FileHeader `json:"image,omitempty" binding:"required"`
}

func (r *UpdateCarImageRequest) Validate() error {
	if r.ImageID == "" {
		return errors.New("image_id is required")
	}

	if _, err := uuid.Parse(r.ImageID); err != nil {
		return errors.New("image_id is invalid")
	}

	if r.Image == nil {
		return errors.New("image is required")
	}

	return nil
}

type UpdateCarImageResponse struct {
	ImageID string `json:"image_id,omitempty"`
	URL     string `json:"url,omitempty"`
}

type GetCarImagesPaginatedRequest struct {
	CarID  string `json:"car_id,omitempty"`
	Offset int32  `json:"offset,omitempty"`
	Limit  int32  `json:"limit,omitempty"`
}

func (r *GetCarImagesPaginatedRequest) Validate() error {
	if r.CarID == "" {
		return errors.New("car_id is required")
	}

	if _, err := uuid.Parse(r.CarID); err != nil {
		return errors.New("car_id is invalid")
	}

	if r.Offset < 0 {
		return errors.New("offset is invalid")
	}

	if r.Limit < 10 {
		return errors.New("limit is invalid")
	}

	return nil
}

type GetCarImagesPaginatedResponse struct {
	Images    []model.CarImage `json:"images,omitempty"`
	TotalCars int32            `json:"total_cars,omitempty"`
}

type DeleteCarImageRequest struct {
	ImageID string `json:"image_id,omitempty"`
}

func (r *DeleteCarImageRequest) Validate() error {
	if r.ImageID == "" {
		return errors.New("image_id is required")
	}

	if _, err := uuid.Parse(r.ImageID); err != nil {
		return errors.New("image_id is invalid")
	}

	return nil
}

type DeleteCarImageResponse struct {
	ImageID string `json:"image_id,omitempty"`
}
