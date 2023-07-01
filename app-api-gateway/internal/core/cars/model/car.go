package model

import (
	"garination.com/gateway/internal/platform/grpc/app-db/proto"
	"time"
)

type Car struct {
	CarID                     string    `json:"car_id,omitempty"`
	BrandID                   string    `json:"brand_id,omitempty"`
	Model                     string    `json:"model,omitempty"`
	Year                      int32     `json:"year,omitempty"`
	Price                     float64   `json:"price,omitempty"`
	Mileage                   int32     `json:"mileage,omitempty"`
	Color                     string    `json:"color,omitempty"`
	Transmission              string    `json:"transmission,omitempty"`
	FuelType                  string    `json:"fuel_type,omitempty"`
	EngineCapacity            int32     `json:"engine_capacity,omitempty"`
	Description               string    `json:"description,omitempty"`
	DealershipID              string    `json:"dealership_id,omitempty"`
	DealerID                  string    `json:"dealer_id,omitempty"`
	Title                     string    `json:"title,omitempty"`
	HorsePower                int32     `json:"horse_power,omitempty"`
	Torque                    int32     `json:"torque,omitempty"`
	TorqueRpm                 int32     `json:"torque_rpm,omitempty"`
	SafetySpecifications      []string  `json:"safety_specifications"`
	PerformanceSpecifications []string  `json:"performance_specifications"`
	ComfortSpecifications     []string  `json:"comfort_specifications"`
	Location                  string    `json:"location"`
	Ownership                 string    `json:"ownership"`
	CreatedAt                 time.Time `json:"created_at,omitempty"`
	UpdatedAt                 time.Time `json:"updated_at,omitempty"`
}

func CarFromProto(model *proto.Car) Car {
	return Car{
		CarID:                     model.CarId,
		BrandID:                   model.BrandId,
		Model:                     model.Model,
		Year:                      model.Year,
		Price:                     model.Price,
		Mileage:                   model.Mileage,
		Color:                     model.Color,
		Transmission:              model.Transmission,
		FuelType:                  model.FuelType,
		EngineCapacity:            model.EngineCapacity,
		Description:               model.Description,
		DealershipID:              model.DealershipId,
		DealerID:                  model.DealerId,
		Title:                     model.Title,
		HorsePower:                model.HorsePower,
		Torque:                    model.Torque,
		TorqueRpm:                 model.TorqueRpm,
		SafetySpecifications:      model.SafetySpecifications,
		PerformanceSpecifications: model.PerformanceSpecifications,
		ComfortSpecifications:     model.ComfortSpecifications,
		Location:                  model.Location,
		Ownership:                 model.Ownership,
		CreatedAt:                 model.CreatedAt.AsTime(),
		UpdatedAt:                 model.UpdatedAt.AsTime(),
	}
}
