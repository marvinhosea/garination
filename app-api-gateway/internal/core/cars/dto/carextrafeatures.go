package dto

import (
	"errors"
	"garination.com/gateway/internal/core/cars/model"
	"github.com/google/uuid"
)

type CreateCarExtraFeatureRequest struct {
	Name  string `json:"name,omitempty"`
	CarID string `json:"car_id,omitempty"`
	Value string `json:"value,omitempty"`
}

func (r *CreateCarExtraFeatureRequest) Validate() error {
	if r.Name == "" {
		return errors.New("name is required")
	}

	if r.CarID == "" {
		return errors.New("car_id is required")
	}

	if _, err := uuid.Parse(r.CarID); err != nil {
		return errors.New("car_id is invalid")
	}

	if r.Value == "" {
		return errors.New("value is required")
	}

	return nil
}

type CreateCarExtraFeatureResponse struct {
	ExtraFeatureID string `json:"extra_feature_id,omitempty"`
}

type UpdateCarExtraFeatureRequest struct {
	ExtraFeatureID string `json:"extra_feature_id,omitempty"`
	Name           string `json:"name,omitempty"`
	Value          string `json:"value,omitempty"`
}

func (r *UpdateCarExtraFeatureRequest) Validate() error {
	if r.ExtraFeatureID == "" {
		return errors.New("extra_feature_id is required")
	}

	if r.Name == "" {
		return errors.New("name is required")
	}

	if r.Value == "" {
		return errors.New("value is required")
	}

	if _, err := uuid.Parse(r.ExtraFeatureID); err != nil {
		return errors.New("extra_feature_id is invalid")
	}

	return nil
}

type UpdateCarExtraFeatureResponse struct {
	ExtraFeatureID string `json:"extra_feature_id,omitempty"`
}

type GetCarExtraFeaturesPaginatedRequest struct {
	Offset int32 `json:"offset,omitempty"`
	Limit  int32 `json:"limit,omitempty"`
}

func (r *GetCarExtraFeaturesPaginatedRequest) Validate() error {
	if r.Offset < 0 {
		return errors.New("offset is invalid")
	}

	if r.Limit < 10 {
		return errors.New("limit should be greater than 10")
	}

	return nil
}

type GetCarExtraFeaturesPaginatedResponse struct {
	ExtraFeatures []model.CarExtraFeature `json:"extra_features,omitempty"`
	Total         int32                   `json:"total,omitempty"`
}

type DeleteCarExtraFeatureRequest struct {
	ExtraFeatureID string `json:"extra_feature_id,omitempty"`
}

func (r *DeleteCarExtraFeatureRequest) Validate() error {
	if r.ExtraFeatureID == "" {
		return errors.New("extra_feature_id is required")
	}

	if _, err := uuid.Parse(r.ExtraFeatureID); err != nil {
		return errors.New("extra_feature_id is invalid")
	}

	return nil
}

type DeleteCarExtraFeatureResponse struct {
	ExtraFeatureID string `json:"extra_feature_id,omitempty"`
}
