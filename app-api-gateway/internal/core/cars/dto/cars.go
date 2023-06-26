package dto

import (
	"garination.com/gateway/internal/core/cars/model"
	"time"
)

type InsertCarRequest struct {
	BrandID        string    `json:"brand_id,omitempty"`
	Model          string    `json:"model,omitempty"`
	Year           int32     `json:"year,omitempty"`
	Price          float64   `json:"price,omitempty"`
	Mileage        int32     `json:"mileage,omitempty"`
	Color          string    `json:"color,omitempty"`
	Transmission   string    `json:"transmission,omitempty"`
	FuelType       string    `json:"fuel_type,omitempty"`
	EngineCapacity string    `json:"engine_capacity,omitempty"`
	Description    string    `json:"description,omitempty"`
	DealershipID   string    `json:"dealership_id,omitempty"`
	DealerID       string    `json:"dealer_id,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}

type InsertCarResponse struct {
	CarID string `json:"car_id,omitempty"`
}

type UpdateCarRequest struct {
	Model          string  `json:"model,omitempty"`
	BrandID        string  `json:"brand_id,omitempty"`
	Year           int32   `json:"year,omitempty"`
	Price          float64 `json:"price,omitempty"`
	Mileage        int32   `json:"mileage,omitempty"`
	Color          string  `json:"color,omitempty"`
	Transmission   string  `json:"transmission,omitempty"`
	FuelType       string  `json:"fuel_type,omitempty"`
	EngineCapacity string  `json:"engine_capacity,omitempty"`
	Description    string  `json:"description,omitempty"`
}

type UpdateCarResponse struct {
	CarID string `json:"car_id,omitempty"`
}

type GetOneCarRequest struct {
	CarID string `json:"car_id,omitempty"`
}

type GetOneCarResponse struct {
	Car           *model.Car               `json:"car,omitempty"`
	Images        []*model.CarImage        `json:"images,omitempty"`
	ExtraFeatures []*model.CarExtraFeature `json:"extra_features,omitempty"`
}

type GetCarsPaginatedRequest struct {
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

type GetCarsPaginatedResponse struct {
	Cars []*model.Car `json:"cars,omitempty"`
}

type GetCarsByDealershipIDPaginatedRequest struct {
	DealershipID string `json:"dealership_id,omitempty"`
	Offset       int    `json:"offset,omitempty"`
	Limit        int    `json:"limit,omitempty"`
}

type GetCarsByDealershipIDPaginatedResponse struct {
	Cars  []*model.Car `json:"cars,omitempty"`
	Total int          `json:"total,omitempty"`
}

type GetCarsByBrandIDPaginatedRequest struct {
	BrandID string `json:"brand_id,omitempty"`
}

type GetCarsByCarBrandIDPaginatedResponse struct {
	Cars  []*model.Car `json:"cars,omitempty"`
	Total int          `json:"total,omitempty"`
}

type GetCarsByDealerIDPaginatedRequest struct {
	DealerID string `json:"dealer_id,omitempty"`
}

type GetCarsByDealerIDPaginatedResponse struct {
	Cars  []*model.Car `json:"cars,omitempty"`
	Total int          `json:"total,omitempty"`
}

type SearchCarsPaginatedRequest struct {
	Search string `json:"search,omitempty"`
	Offset int    `json:"offset,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

type SearchCarsPaginatedResponse struct {
	Cars  []*model.Car `json:"cars,omitempty"`
	Total int          `json:"total,omitempty"`
}

type GetCarsByFieldRequest struct {
	Field     string `json:"field,omitempty"`
	Value     string `json:"value,omitempty"`
	Condition string `json:"condition,omitempty"`
}

type GetCarsByFieldResponse struct {
	Cars  []*model.Car `json:"cars,omitempty"`
	Total int          `json:"total,omitempty"`
}

type GetCarByDealerCountRequest struct {
	DealerID string `json:"dealer_id,omitempty"`
}

type GetCarByDealerCountResponse struct {
	Total int `json:"total,omitempty"`
}

type GetCarByDealershipCountRequest struct {
	DealershipID string `json:"dealership_id,omitempty"`
}

type GetCarByDealershipCountResponse struct {
	Total int `json:"total,omitempty"`
}

type GetCarByBrandCountRequest struct {
	BrandID string `json:"brand_id,omitempty"`
}

type GetCarByCarBrandCountResponse struct {
	Total int `json:"total,omitempty"`
}

type DeleteCarRequest struct {
	CarID string `json:"car_id,omitempty"`
}

type DeleteCarResponse struct {
	CarID string `json:"car_id,omitempty"`
}
