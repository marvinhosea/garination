package dto

import "garination.com/gateway/internal/core/cars/model"

type CreateCarImageRequest struct {
	CarID string `json:"car_id,omitempty"`
	URL   string `json:"url,omitempty"`
}

type CreateCarImageResponse struct {
	ImageID string `json:"image_id,omitempty"`
	URL     string `json:"url,omitempty"`
}

type UpdateCarImageRequest struct {
	ImageID string `json:"image_id,omitempty"`
	URL     string `json:"url,omitempty"`
}

type UpdateCarImageResponse struct {
	ImageID string `json:"image_id,omitempty"`
	URL     string `json:"url,omitempty"`
}

type GetCarImagesPaginatedRequest struct {
	CarID  string `json:"car_id,omitempty"`
	Offset int    `json:"offset,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

type GetCarImagesPaginatedResponse struct {
	Images []*model.CarImage `json:"images,omitempty"`
}

type DeleteCarImageRequest struct {
	ImageID string `json:"image_id,omitempty"`
}

type DeleteCarImageResponse struct {
	ImageID string `json:"image_id,omitempty"`
}
