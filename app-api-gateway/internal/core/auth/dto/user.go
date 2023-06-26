package dto

import (
	"errors"
	"github.com/google/uuid"
	"regexp"
)

type AuthUpdateUserMetaRequest struct {
	UserID       string `json:"user_id"`
	FacebookUrl  string `json:"facebook_url"`
	TwitterUrl   string `json:"twitter_url"`
	InstagramUrl string `json:"instagram_url"`
	LinkedinUrl  string `json:"linkedin_url"`
	WebsiteUrl   string `json:"website_url"`
	DealershipID string `json:"dealership_id"`
}

func (r AuthUpdateUserMetaRequest) Validate() error {
	if r.UserID == "" {
		return errors.New("user_id is required")
	}

	// check if it is a valid uuid
	_, err := uuid.Parse(r.UserID)
	if err != nil {
		return errors.New("user_id is not a valid uuid")
	}

	// validate dealership id
	if r.DealershipID != "" {
		_, err := uuid.Parse(r.DealershipID)
		if err != nil {
			return errors.New("dealership_id is not a valid uuid")
		}
	}

	re := regexp.MustCompile(urlRegex)

	// validate facebook url
	if r.FacebookUrl != "" {
		if !re.MatchString(r.FacebookUrl) {
			return errors.New("facebook_url is not a valid url")
		}
	}

	// validate twitter url
	if r.TwitterUrl != "" {
		if !re.MatchString(r.TwitterUrl) {
			return errors.New("twitter_url is not a valid url")
		}
	}

	// validate instagram url
	if r.InstagramUrl != "" {
		if !re.MatchString(r.InstagramUrl) {
			return errors.New("instagram_url is not a valid url")
		}
	}

	// validate linkedin url
	if r.LinkedinUrl != "" {
		if !re.MatchString(r.LinkedinUrl) {
			return errors.New("linkedin_url is not a valid url")
		}
	}

	// validate website url
	if r.WebsiteUrl != "" {
		if !re.MatchString(r.WebsiteUrl) {
			return errors.New("website_url is not a valid url")
		}
	}

	return nil
}

type AuthUpdateUserMetaResponse struct {
	UserMetaID   string `json:"user_meta_id"`
	UserID       string `json:"user_id"`
	FacebookUrl  string `json:"facebook_url"`
	TwitterUrl   string `json:"twitter_url"`
	InstagramUrl string `json:"instagram_url"`
	LinkedinUrl  string `json:"linkedin_url"`
	WebsiteUrl   string `json:"website_url"`
	DealershipID string `json:"dealership_id"`
}

type AuthGetUserMetaRequest struct {
	UserID string `json:"user_id"`
}

func (r AuthGetUserMetaRequest) Validate() error {
	if r.UserID == "" {
		return errors.New("user_id is required")
	}

	// check if it is a valid uuid
	_, err := uuid.Parse(r.UserID)
	if err != nil {
		return errors.New("user_id is not a valid uuid")
	}

	return nil
}

type AuthGetUserMetaResponse struct {
	UserMetaID   string `json:"user_meta_id"`
	UserID       string `json:"user_id"`
	FacebookUrl  string `json:"facebook_url"`
	TwitterUrl   string `json:"twitter_url"`
	InstagramUrl string `json:"instagram_url"`
	LinkedinUrl  string `json:"linkedin_url"`
	WebsiteUrl   string `json:"website_url"`
	DealershipID string `json:"dealership_id"`
}

type AuthChangeUserDealershipRequest struct {
	UserID       string `json:"user_id"`
	DealershipID string `json:"dealership_id"`
}

func (r AuthChangeUserDealershipRequest) Validate() error {

	var err error
	if r.UserID == "" {
		err = errors.Join(err, errors.New("user_id is required"))
	}

	if r.DealershipID == "" {
		err = errors.Join(err, errors.New("dealership_id is required"))
	}

	// check if it is a valid uuid
	_, err = uuid.Parse(r.UserID)
	if err != nil {
		err = errors.Join(err, errors.New("user_id is not a valid uuid"))
	}

	_, err = uuid.Parse(r.DealershipID)
	if err != nil {
		err = errors.Join(err, errors.New("dealership_id is not a valid uuid"))
	}

	return err
}

type AuthChangeUserDealershipResponse struct {
	DealershipID string `json:"dealership_id"`
}
