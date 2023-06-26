package model

type CarExtraFeature struct {
	ExtraFeatureID string `json:"extra_feature_id,omitempty"`
	Name           string `json:"name,omitempty"`
	Value          string `json:"value,omitempty"`
	CarID          string `json:"car_id,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
}
