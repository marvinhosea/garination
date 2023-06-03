// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package postgres

import (
	"github.com/jackc/pgx/v5/pgtype"
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
	CreatedAt    pgtype.Timestamp
	UpdatedAt    pgtype.Timestamp
}

type UserMetum struct {
	UserMetaID   string
	UserID       string
	FacebookUrl  pgtype.Text
	TwitterUrl   pgtype.Text
	InstagramUrl pgtype.Text
	LinkedinUrl  pgtype.Text
	WebsiteUrl   pgtype.Text
	DealershipID pgtype.Text
}
