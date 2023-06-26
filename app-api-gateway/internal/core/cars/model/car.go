package model

import "time"

type Car struct {
	CarID          string    `json:"car_id,omitempty"`
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
