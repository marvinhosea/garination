package model

import "time"

type CarImage struct {
	ImageID   string    `json:"image_id,omitempty"`
	CarID     string    `json:"car_id,omitempty"`
	URL       string    `json:"url,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
