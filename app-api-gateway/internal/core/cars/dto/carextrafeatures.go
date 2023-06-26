package dto

import "garination.com/gateway/internal/core/cars/model"

type CreateCarExtraFeatureRequest struct {
	Name  string `json:"name,omitempty"`
	CarID string `json:"car_id,omitempty"`
	Value string `json:"value,omitempty"`
}

type CreateCarExtraFeatureResponse struct {
	ExtraFeatureID string `json:"extra_feature_id,omitempty"`
}

type UpdateCarExtraFeatureRequest struct {
	ExtraFeatureID string `json:"extra_feature_id,omitempty"`
	Name           string `json:"name,omitempty"`
	Value          string `json:"value,omitempty"`
}

type UpdateCarExtraFeatureResponse struct {
	ExtraFeatureID string `json:"extra_feature_id,omitempty"`
}

type GetCarExtraFeaturesPaginatedRequest struct {
	Offset int32 `json:"offset,omitempty"`
	Limit  int32 `json:"limit,omitempty"`
}

type GetCarExtraFeaturesPaginatedResponse struct {
	ExtraFeatures []model.CarExtraFeature `json:"extra_features,omitempty"`
	Total         int32                   `json:"total,omitempty"`
}

type DeleteCarExtraFeatureRequest struct {
	ExtraFeatureID string `json:"extra_feature_id,omitempty"`
}

type DeleteCarExtraFeatureResponse struct {
	ExtraFeatureID string `json:"extra_feature_id,omitempty"`
}
