// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: queries.sql

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getUserDealership = `-- name: GetUserDealership :one
SELECT d.dealership_id, d.owner_id, d.name, d.display_name, d.address, d.city, d.state, d.zip, d.phone, d.email, d.website, d.facebook_url, d.twitter_url, d.instagram_url, d.linkedin_url, d.logo_url, d.cover_url, d.description, d.created_at, d.updated_at FROM dealership d JOIN user_meta u ON d.dealership_id = u.dealership_id WHERE u.user_id = $1 LIMIT 1
`

func (q *Queries) GetUserDealership(ctx context.Context, userID string) (Dealership, error) {
	row := q.db.QueryRow(ctx, getUserDealership, userID)
	var i Dealership
	err := row.Scan(
		&i.DealershipID,
		&i.OwnerID,
		&i.Name,
		&i.DisplayName,
		&i.Address,
		&i.City,
		&i.State,
		&i.Zip,
		&i.Phone,
		&i.Email,
		&i.Website,
		&i.FacebookUrl,
		&i.TwitterUrl,
		&i.InstagramUrl,
		&i.LinkedinUrl,
		&i.LogoUrl,
		&i.CoverUrl,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserMeta = `-- name: GetUserMeta :one
SELECT user_meta_id, user_id, facebook_url, twitter_url, instagram_url, linkedin_url, website_url, dealership_id FROM user_meta WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetUserMeta(ctx context.Context, userID string) (UserMetum, error) {
	row := q.db.QueryRow(ctx, getUserMeta, userID)
	var i UserMetum
	err := row.Scan(
		&i.UserMetaID,
		&i.UserID,
		&i.FacebookUrl,
		&i.TwitterUrl,
		&i.InstagramUrl,
		&i.LinkedinUrl,
		&i.WebsiteUrl,
		&i.DealershipID,
	)
	return i, err
}

const insertDealership = `-- name: InsertDealership :one
INSERT INTO dealership (
    dealership_id,
    owner_id,
    name,
    display_name,
    address,
    city,
    state,
    zip,
    phone,
    email,
    website,
    facebook_url,
    twitter_url,
    instagram_url,
    linkedin_url,
    logo_url,
    cover_url,
    description,
    created_at,
    updated_at
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20
         ) RETURNING dealership_id, owner_id, name, display_name, address, city, state, zip, phone, email, website, facebook_url, twitter_url, instagram_url, linkedin_url, logo_url, cover_url, description, created_at, updated_at
`

type InsertDealershipParams struct {
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

func (q *Queries) InsertDealership(ctx context.Context, arg InsertDealershipParams) (Dealership, error) {
	row := q.db.QueryRow(ctx, insertDealership,
		arg.DealershipID,
		arg.OwnerID,
		arg.Name,
		arg.DisplayName,
		arg.Address,
		arg.City,
		arg.State,
		arg.Zip,
		arg.Phone,
		arg.Email,
		arg.Website,
		arg.FacebookUrl,
		arg.TwitterUrl,
		arg.InstagramUrl,
		arg.LinkedinUrl,
		arg.LogoUrl,
		arg.CoverUrl,
		arg.Description,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Dealership
	err := row.Scan(
		&i.DealershipID,
		&i.OwnerID,
		&i.Name,
		&i.DisplayName,
		&i.Address,
		&i.City,
		&i.State,
		&i.Zip,
		&i.Phone,
		&i.Email,
		&i.Website,
		&i.FacebookUrl,
		&i.TwitterUrl,
		&i.InstagramUrl,
		&i.LinkedinUrl,
		&i.LogoUrl,
		&i.CoverUrl,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertUserMeta = `-- name: InsertUserMeta :one
INSERT INTO user_meta (
    user_meta_id,
    user_id,
    facebook_url,
    twitter_url,
    instagram_url,
    linkedin_url,
    website_url,
    dealership_id
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8
         ) RETURNING user_meta_id, user_id, facebook_url, twitter_url, instagram_url, linkedin_url, website_url, dealership_id
`

type InsertUserMetaParams struct {
	UserMetaID   string
	UserID       string
	FacebookUrl  pgtype.Text
	TwitterUrl   pgtype.Text
	InstagramUrl pgtype.Text
	LinkedinUrl  pgtype.Text
	WebsiteUrl   pgtype.Text
	DealershipID pgtype.Text
}

func (q *Queries) InsertUserMeta(ctx context.Context, arg InsertUserMetaParams) (UserMetum, error) {
	row := q.db.QueryRow(ctx, insertUserMeta,
		arg.UserMetaID,
		arg.UserID,
		arg.FacebookUrl,
		arg.TwitterUrl,
		arg.InstagramUrl,
		arg.LinkedinUrl,
		arg.WebsiteUrl,
		arg.DealershipID,
	)
	var i UserMetum
	err := row.Scan(
		&i.UserMetaID,
		&i.UserID,
		&i.FacebookUrl,
		&i.TwitterUrl,
		&i.InstagramUrl,
		&i.LinkedinUrl,
		&i.WebsiteUrl,
		&i.DealershipID,
	)
	return i, err
}

const updateUserMeta = `-- name: UpdateUserMeta :one
UPDATE user_meta SET
    facebook_url = $2,
    twitter_url = $3,
    instagram_url = $4,
    linkedin_url = $5,
    website_url = $6,
    user_meta_id = $7
WHERE user_id = $1 RETURNING user_meta_id, user_id, facebook_url, twitter_url, instagram_url, linkedin_url, website_url, dealership_id
`

type UpdateUserMetaParams struct {
	UserID       string
	FacebookUrl  pgtype.Text
	TwitterUrl   pgtype.Text
	InstagramUrl pgtype.Text
	LinkedinUrl  pgtype.Text
	WebsiteUrl   pgtype.Text
	UserMetaID   string
}

func (q *Queries) UpdateUserMeta(ctx context.Context, arg UpdateUserMetaParams) (UserMetum, error) {
	row := q.db.QueryRow(ctx, updateUserMeta,
		arg.UserID,
		arg.FacebookUrl,
		arg.TwitterUrl,
		arg.InstagramUrl,
		arg.LinkedinUrl,
		arg.WebsiteUrl,
		arg.UserMetaID,
	)
	var i UserMetum
	err := row.Scan(
		&i.UserMetaID,
		&i.UserID,
		&i.FacebookUrl,
		&i.TwitterUrl,
		&i.InstagramUrl,
		&i.LinkedinUrl,
		&i.WebsiteUrl,
		&i.DealershipID,
	)
	return i, err
}
