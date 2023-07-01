package model

import (
	"time"
)

type SparePart struct {
	SparePartID         string    `json:"spare_part_id,omitempty"`
	Name                string    `json:"name,omitempty"`
	Description         string    `json:"description,omitempty"`
	Price               float64   `json:"price,omitempty"`
	Used                bool      `json:"used,omitempty"`
	CarModel            string    `json:"car_model,omitempty"`
	CarBrand            string    `json:"car_brand,omitempty"`
	OtherCompatibleCars []string  `json:"other_compatible_cars,omitempty"`
	CarYear             int32     `json:"car_year,omitempty"`
	IsUniversal         bool      `json:"is_universal,omitempty"`
	Category            string    `json:"category,omitempty"`
	PartNumber          string    `json:"part_number,omitempty"`
	DealershipID        string    `json:"dealership_id,omitempty"`
	DealerID            string    `json:"dealer_id,omitempty"`
	CreatedAt           time.Time `json:"created_at,omitempty"`
	UpdatedAt           time.Time `json:"updated_at,omitempty"`
	DeletedAt           time.Time `json:"deleted_at,omitempty"`
}

type SparePartImage struct {
	SparePartImageID string    `json:"spare_part_image_id,omitempty"`
	SparePartID      string    `json:"spare_part_id,omitempty"`
	ImageURL         string    `json:"image_url,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
	DeletedAt        time.Time `json:"deleted_at,omitempty"`
}
