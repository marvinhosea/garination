package dto

import (
	"errors"
	"garination.com/gateway/internal/core/cars/model"
	"github.com/google/uuid"
	"time"
)

type InsertCarRequest struct {
	BrandID                   string   `json:"brand_id"`
	Model                     string   `json:"model"`
	Year                      int32    `json:"year"`
	Price                     float64  `json:"price"`
	Mileage                   int32    `json:"mileage"`
	Color                     string   `json:"color"`
	Transmission              string   `json:"transmission"`
	FuelType                  string   `json:"fuel_type"`
	EngineCapacity            int32    `json:"engine_capacity"`
	Description               string   `json:"description"`
	DealershipID              string   `json:"dealership_id"`
	DealerID                  string   `json:"dealer_id"`
	Title                     string   `json:"title"`
	HorsePower                int32    `json:"horse_power"`
	Torque                    int32    `json:"torque"`
	TorqueRpm                 int32    `json:"torque_rpm"`
	SafetySpecifications      []string `json:"safety_specifications"`
	PerformanceSpecifications []string `json:"performance_specifications"`
	ComfortSpecifications     []string `json:"comfort_specifications"`
	Location                  string   `json:"location"`
	Ownership                 string   `json:"ownership"`
}

func (i *InsertCarRequest) Validate() error {
	// check if brand id is uuid v4
	if _, err := uuid.Parse(i.BrandID); err != nil {
		return err
	}

	// check if dealership id is uuid v4
	if i.DealershipID != "" {
		if _, err := uuid.Parse(i.DealershipID); err != nil {
			return err
		}
	}

	if i.Title == "" {
		return errors.New("title cannot be empty")
	}

	// check if dealer id is uuid v4
	if i.DealerID != "" {
		if _, err := uuid.Parse(i.DealerID); err != nil {
			return err
		}
	}

	if i.DealershipID != "" && i.DealerID != "" {
		return errors.New("dealership_id and dealer_id cannot be both set")
	}

	// validate price is not negative
	if i.Price <= 0 {
		return errors.New("price cannot be negative")
	}

	// validate mileage is not negative
	if i.Mileage < 0 {
		return errors.New("mileage cannot be negative")
	}

	// validate year is not negative
	if i.Year < 1970 {
		return errors.New("year cannot be less than 1970")
	}

	if i.Year > int32(time.Now().Year()) {
		return errors.New("year cannot be greater than current year")
	}

	// validate color is not empty
	if i.Color == "" {
		return errors.New("color cannot be empty")
	}

	// validate fuel type is not empty and is petrol, diesel, or hybrid
	if i.FuelType == "" {
		return errors.New("fuel type cannot be empty")
	}

	if i.FuelType != "petrol" && i.FuelType != "diesel" && i.FuelType != "electric" {
		return errors.New("fuel type must be petrol, diesel, or electric")
	}

	// validate transmission is not empty and is automatic or manual or semi-automatic
	if i.Transmission == "" {
		return errors.New("transmission cannot be empty")
	}

	if i.Transmission != "automatic" && i.Transmission != "manual" && i.Transmission != "semi-automatic" {
		return errors.New("transmission must be automatic, manual, or semi-automatic")
	}

	if i.Ownership != "used" && i.Ownership != "new" && i.Ownership != "locally-used" {
		return errors.New("ownership must be used, new, or locally-used")
	}

	// validate engine capacity is not empty >= 500cc and <= 10000cc

	if i.EngineCapacity < 500 || i.EngineCapacity > 10000 {
		return errors.New("engine capacity must be >= 500cc and <= 10000cc")
	}

	// validate model is not empty
	if i.Model == "" {
		return errors.New("model cannot be empty")
	}

	// validate description is not empty
	if i.Description == "" {
		return errors.New("description cannot be empty")
	}

	return nil
}

type InsertCarResponse struct {
	CarID string `json:"car_id,omitempty"`
}

type UpdateCarRequest struct {
	CarID          string  `json:"car_id,omitempty"`
	Model          string  `json:"model,omitempty"`
	BrandID        string  `json:"brand_id,omitempty"`
	Year           int32   `json:"year,omitempty"`
	Price          float64 `json:"price,omitempty"`
	Mileage        int32   `json:"mileage,omitempty"`
	Color          string  `json:"color,omitempty"`
	Transmission   string  `json:"transmission,omitempty"`
	FuelType       string  `json:"fuel_type,omitempty"`
	EngineCapacity int32   `json:"engine_capacity,omitempty"`
	Description    string  `json:"description,omitempty"`
}

func (i *UpdateCarRequest) Validate() error {
	// check if brand id is uuid v4
	if _, err := uuid.Parse(i.BrandID); err != nil {
		return err
	}

	// validate price is not negative
	if i.Price < 0 {
		return errors.New("price cannot be negative")
	}

	// validate mileage is not negative
	if i.Mileage < 0 {
		return errors.New("mileage cannot be negative")
	}

	// validate year is not negative
	if i.Year < 1970 {
		return errors.New("year cannot be less than 1970")
	}

	if i.Year > int32(time.Now().Year()) {
		return errors.New("year cannot be greater than current year")
	}

	// validate color is not empty
	if i.Color == "" {
		return errors.New("color cannot be empty")
	}

	// validate fuel type is not empty and is petrol, diesel, or hybrid
	if i.FuelType == "" {
		return errors.New("fuel type cannot be empty")
	}

	if i.FuelType != "petrol" && i.FuelType != "diesel" && i.FuelType != "electric" {
		return errors.New("fuel type must be petrol, diesel, or electric")
	}

	// validate transmission is not empty and is automatic or manual or semi-automatic
	if i.Transmission == "" {
		return errors.New("transmission cannot be empty")
	}

	if i.Transmission != "automatic" && i.Transmission != "manual" && i.Transmission != "semi-automatic" {
		return errors.New("transmission must be automatic, manual, or semi-automatic")
	}

	// validate engine capacity is not empty >= 500cc and <= 10000cc

	if i.EngineCapacity < 500 || i.EngineCapacity > 10000 {
		return errors.New("engine capacity must be >= 500cc and <= 10000cc")
	}

	// validate model is not empty
	if i.Model == "" {
		return errors.New("model cannot be empty")
	}

	// validate description is not empty
	if i.Description == "" {
		return errors.New("description cannot be empty")
	}

	return nil
}

type UpdateCarResponse struct {
	CarID string `json:"car_id,omitempty"`
}

type GetOneCarRequest struct {
	CarID string `json:"car_id,omitempty"`
}

func (i *GetOneCarRequest) Validate() error {
	// check if car id is uuid v4
	if _, err := uuid.Parse(i.CarID); err != nil {
		return err
	}

	return nil
}

type GetOneCarResponse struct {
	Car           model.Car               `json:"car,omitempty"`
	Images        []model.CarImage        `json:"images,omitempty"`
	ExtraFeatures []model.CarExtraFeature `json:"extra_features,omitempty"`
}

type GetCarsPaginatedRequest struct {
	Offset int32 `json:"offset,omitempty"`
	Limit  int32 `json:"limit,omitempty"`
}

func (i *GetCarsPaginatedRequest) Validate() error {
	// validate offset is not negative
	if i.Offset < 0 {
		return errors.New("offset cannot be negative")
	}

	// validate limit is not negative
	if i.Limit < 10 {
		return errors.New("limit cannot be less than 10")
	}

	return nil
}

type GetCarsPaginatedResponse struct {
	Cars []model.Car `json:"cars,omitempty"`
}

type GetCarsByDealershipIDPaginatedRequest struct {
	DealershipID string `json:"dealership_id,omitempty"`
	Offset       int32  `json:"offset,omitempty"`
	Limit        int32  `json:"limit,omitempty"`
}

func (i *GetCarsByDealershipIDPaginatedRequest) Validate() error {
	// validate offset is not negative
	if i.Offset < 0 {
		return errors.New("offset cannot be negative")
	}

	// validate limit is not negative
	if i.Limit < 10 {
		return errors.New("limit cannot be less than 10")
	}

	// check if dealership id is uuid v4
	if _, err := uuid.Parse(i.DealershipID); err != nil {
		return err
	}

	return nil
}

type GetCarsByDealershipIDPaginatedResponse struct {
	Cars  []model.Car `json:"cars,omitempty"`
	Total int32       `json:"total,omitempty"`
}

type GetCarsByBrandIDPaginatedRequest struct {
	BrandID string `json:"brand_id,omitempty"`
	Offset  int32  `json:"offset,omitempty"`
	Limit   int32  `json:"limit,omitempty"`
}

func (i *GetCarsByBrandIDPaginatedRequest) Validate() error {
	// check if brand id is uuid v4
	if _, err := uuid.Parse(i.BrandID); err != nil {
		return err
	}

	if i.Offset < 0 {
		return errors.New("offset cannot be negative")
	}

	if i.Limit < 10 {
		return errors.New("limit cannot be less than 10")
	}

	return nil
}

type GetCarsByCarBrandIDPaginatedResponse struct {
	Cars  []model.Car `json:"cars,omitempty"`
	Total int32       `json:"total,omitempty"`
}

type GetCarsByDealerIDPaginatedRequest struct {
	DealerID string `json:"dealer_id,omitempty"`
	Offset   int32  `json:"offset,omitempty"`
	Limit    int32  `json:"limit,omitempty"`
}

func (i *GetCarsByDealerIDPaginatedRequest) Validate() error {
	// check if dealer id is uuid v4
	if _, err := uuid.Parse(i.DealerID); err != nil {
		return err
	}

	if i.Offset < 0 {
		return errors.New("offset cannot be negative")
	}

	if i.Limit < 10 {
		return errors.New("limit cannot be less than 10")
	}

	return nil
}

type GetCarsByDealerIDPaginatedResponse struct {
	Cars  []model.Car `json:"cars,omitempty"`
	Total int32       `json:"total,omitempty"`
}

type SearchCarsPaginatedRequest struct {
	Search string `json:"search,omitempty"`
	Offset int32  `json:"offset,omitempty"`
	Limit  int32  `json:"limit,omitempty"`
}

func (i *SearchCarsPaginatedRequest) Validate() error {
	if i.Offset < 0 {
		return errors.New("offset cannot be negative")
	}

	if i.Limit < 10 {
		return errors.New("limit cannot be less than 10")
	}

	if i.Search == "" {
		return errors.New("search cannot be empty")
	}

	return nil
}

type SearchCarsPaginatedResponse struct {
	Cars  []model.Car `json:"cars,omitempty"`
	Total int32       `json:"total,omitempty"`
}

type FieldFilter struct {
	Field     string `json:"field,omitempty"`
	Value     string `json:"value,omitempty"`
	Condition string `json:"condition,omitempty"`
}

type GetCarsByFieldRequest struct {
	Filters []FieldFilter `json:"filters,omitempty"`
	Offset  int32         `json:"offset,omitempty"`
	Limit   int32         `json:"limit,omitempty"`
}

func (i *GetCarsByFieldRequest) Validate() error {
	if i.Offset < 0 {
		return errors.New("offset cannot be negative")
	}

	if i.Limit < 10 {
		return errors.New("limit cannot be less than 10")
	}

	if len(i.Filters) == 0 {
		return errors.New("filters cannot be empty")
	}

	// validate filters
	for _, filter := range i.Filters {
		if filter.Field == "" {
			return errors.New("field cannot be empty")
		}

		if filter.Value == "" {
			return errors.New("value cannot be empty")
		}

		if filter.Condition == "" {
			return errors.New("condition cannot be empty")
		}
	}

	return nil
}

type GetCarsByFieldResponse struct {
	Cars  []model.Car `json:"cars,omitempty"`
	Total int32       `json:"total,omitempty"`
}

type GetCarByDealerCountRequest struct {
	DealerID string `json:"dealer_id,omitempty"`
}

func (i *GetCarByDealerCountRequest) Validate() error {
	// check if dealer id is uuid v4
	if _, err := uuid.Parse(i.DealerID); err != nil {
		return err
	}

	return nil
}

type GetCarByDealerCountResponse struct {
	Total int64 `json:"total,omitempty"`
}

type GetCarByDealershipCountRequest struct {
	DealershipID string `json:"dealership_id,omitempty"`
}

func (i *GetCarByDealershipCountRequest) Validate() error {
	// check if dealership id is uuid v4
	if _, err := uuid.Parse(i.DealershipID); err != nil {
		return err
	}

	return nil
}

type GetCarByDealershipCountResponse struct {
	Total int64 `json:"total,omitempty"`
}

type GetCarByBrandCountRequest struct {
	BrandID string `json:"brand_id,omitempty"`
}

func (i *GetCarByBrandCountRequest) Validate() error {
	// check if brand id is uuid v4
	if _, err := uuid.Parse(i.BrandID); err != nil {
		return err
	}

	return nil
}

type GetCarByCarBrandCountResponse struct {
	Total int64 `json:"total,omitempty"`
}

type DeleteCarRequest struct {
	CarID string `json:"car_id,omitempty"`
}

func (i *DeleteCarRequest) Validate() error {
	// check if car id is uuid v4
	if _, err := uuid.Parse(i.CarID); err != nil {
		return err
	}

	return nil
}

type DeleteCarResponse struct {
	CarID string `json:"car_id,omitempty"`
}
