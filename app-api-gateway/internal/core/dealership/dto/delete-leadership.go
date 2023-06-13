package dto

import (
	"errors"
	"github.com/google/uuid"
)

type DeleteDealershipRequest struct {
	DealershipID string `json:"dealership_id"`
}

func (r *DeleteDealershipRequest) Validate() error {
	if r.DealershipID == "" {
		return errors.New("dealership_id is required")
	}

	if _, err := uuid.Parse(r.DealershipID); err != nil {
		return errors.New("dealership_id is not a valid uuid")
	}

	return nil
}

type DeleteLeadershipResponse struct {
	DealershipID string `json:"dealership_id"`
}
