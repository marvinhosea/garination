package dto

import (
	"errors"
	"garination.com/gateway/pkg/validationutil"
	"github.com/google/uuid"
	"time"
)

type UpdateDealershipRequest struct {
	DealershipID string `json:"dealership_id"`
	Name         string `json:"name"`
	DisplayName  string `json:"display_name"`
	Address      string `json:"address"`
	City         string `json:"city"`
	State        string `json:"state"`
	Zip          string `json:"zip"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Website      string `json:"website"`
	FacebookUrl  string `json:"facebook_url"`
	TwitterUrl   string `json:"twitter_url"`
	InstagramUrl string `json:"instagram_url"`
	LinkedinUrl  string `json:"linkedin_url"`
	LogoUrl      string `json:"logo_url"`
	CoverUrl     string `json:"cover_url"`
	Description  string `json:"description"`
}

func (r UpdateDealershipRequest) Validate() error {
	// check if it is a valid uuid
	_, err := uuid.Parse(r.DealershipID)
	if err != nil {
		return errors.New("dealership_id is not a valid uuid")
	}

	// validate name is not empty
	if r.Name == "" {
		return errors.New("name is required")
	}

	// validate display name is not empty
	if r.DisplayName == "" {
		return errors.New("display_name is required")
	}

	// validate address is not empty
	if r.Address == "" {
		return errors.New("address is required")
	}

	// validate city is not empty
	if r.City == "" {
		return errors.New("city is required")
	}

	// validate state is not empty
	if r.State == "" {
		return errors.New("state is required")
	}

	// validate zip is not empty
	if r.Zip == "" {
		return errors.New("zip is required")
	}

	// validate phone is not empty
	if !validationutil.ValidatePhone(r.Phone) {
		return errors.New("valid phone is required")
	}

	// validate email is not empty
	if !validationutil.ValidateEmail(r.Email) {
		return errors.New("valid email is required")
	}

	// validate website is not empty
	if r.Website != "" {
		if !validationutil.ValidateURL(r.Website) {
			return errors.New("website is not a valid url")
		}
	}

	if r.FacebookUrl != "" {
		if !validationutil.ValidateURL(r.FacebookUrl) {
			return errors.New("facebook_url is not a valid url")
		}
	}

	if r.TwitterUrl != "" {
		if !validationutil.ValidateURL(r.TwitterUrl) {
			return errors.New("twitter_url is not a valid url")
		}
	}

	if r.InstagramUrl != "" {
		if !validationutil.ValidateURL(r.InstagramUrl) {
			return errors.New("instagram_url is not a valid url")
		}
	}

	if r.LinkedinUrl != "" {
		if !validationutil.ValidateURL(r.LinkedinUrl) {
			return errors.New("linkedin_url is not a valid url")
		}
	}

	if r.LogoUrl != "" {
		if !validationutil.ValidateURL(r.LogoUrl) {
			return errors.New("logo_url is not a valid url")
		}
	}

	if r.CoverUrl != "" {
		if !validationutil.ValidateURL(r.CoverUrl) {
			return errors.New("cover_url is not a valid url")
		}
	}

	if r.Description != "" {
		return errors.New("it'd be nice to have a description")
	}

	return nil

}

type UpdateDealershipResponse struct {
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
}
