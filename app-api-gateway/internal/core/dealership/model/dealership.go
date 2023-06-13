package model

import (
	"time"
)

type Dealership struct {
	DealershipID string
	OwnerID      string
	Name         string
	DisplayName  string
	Address      string
	City         string
	State        string
	Zip          string
	Phone        string
	Email        string
	Website      string
	FacebookUrl  string
	TwitterUrl   string
	InstagramUrl string
	LinkedinUrl  string
	LogoUrl      string
	CoverUrl     string
	Description  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
