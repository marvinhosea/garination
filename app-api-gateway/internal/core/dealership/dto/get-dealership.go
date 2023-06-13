package dto

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type GetDealershipRequest struct {
	ID string `uri:"id" binding:"required"`
}

func (r GetDealershipRequest) Validate() error {
	if r.ID == "" {
		return errors.New("id is required")
	}

	if _, err := uuid.Parse(r.ID); err != nil {
		return errors.New("id is not a valid uuid")
	}

	return nil
}

type GetDealershipResponse struct {
	DealershipID string    `json:"dealership_id"`
	OwnerID      string    `json:"owner_id"`
	Name         string    `json:"name"`
	DisplayName  string    `json:"display_name"`
	Address      string    `json:"address"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	Zip          string    `json:"zip"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	Website      string    `json:"website"`
	FacebookUrl  string    `json:"facebook_url"`
	TwitterUrl   string    `json:"twitter_url"`
	InstagramUrl string    `json:"instagram_url"`
	LinkedinUrl  string    `json:"linkedin_url"`
	LogoUrl      string    `json:"logo_url"`
	CoverUrl     string    `json:"cover_url"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
