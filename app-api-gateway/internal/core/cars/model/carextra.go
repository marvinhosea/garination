package model

import (
	"garination.com/gateway/internal/platform/grpc/app-db/proto"
	"time"
)

type CarExtraFeature struct {
	ExtraFeatureID string    `json:"extra_feature_id,omitempty"`
	Name           string    `json:"name,omitempty"`
	Value          string    `json:"value,omitempty"`
	CarID          string    `json:"car_id,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}

func CarExtraFromProto(mode *proto.CarExtraFeature) CarExtraFeature {
	return CarExtraFeature{
		ExtraFeatureID: mode.CarExtraId,
		Name:           mode.Name,
		Value:          mode.Value,
		CarID:          mode.CarId,
		CreatedAt:      mode.CreatedAt.AsTime(),
		UpdatedAt:      mode.UpdatedAt.AsTime(),
	}
}
