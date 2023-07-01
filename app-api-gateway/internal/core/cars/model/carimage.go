package model

import (
	"garination.com/gateway/internal/platform/grpc/app-db/proto"
	"time"
)

type CarImage struct {
	ImageID   string    `json:"image_id,omitempty"`
	CarID     string    `json:"car_id,omitempty"`
	URL       string    `json:"url,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func CarImageFromProto(model *proto.CarImage) CarImage {
	return CarImage{
		ImageID:   model.CarImageId,
		CarID:     model.CarId,
		URL:       model.Url,
		CreatedAt: model.CreatedAt.AsTime(),
		UpdatedAt: model.UpdatedAt.AsTime(),
	}
}
