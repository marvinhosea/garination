-- name: GetUserMeta :one
SELECT * FROM user_meta WHERE user_id = $1 LIMIT 1;


-- name: InsertUserMeta :one
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
         ) RETURNING *;

-- name: UpdateUserMeta :one
UPDATE user_meta SET
    facebook_url = $2,
    twitter_url = $3,
    instagram_url = $4,
    linkedin_url = $5,
    website_url = $6,
    user_meta_id = $7
WHERE user_id = $1 RETURNING *;

-- name: GetUserDealership :one
SELECT d.* FROM dealership d JOIN user_meta u ON d.dealership_id = u.dealership_id WHERE u.user_id = $1 LIMIT 1;

-- name: InsertDealership :one
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
         ) RETURNING *;


-- name: UpdateDealership :one
UPDATE dealership SET
    name = $2,
    display_name = $3,
    address = $4,
    city = $5,
    state = $6,
    zip = $7,
    phone = $8,
    email = $9,
    website = $10,
    facebook_url = $11,
    twitter_url = $12,
    instagram_url = $13,
    linkedin_url = $14,
    logo_url = $15,
    cover_url = $16,
    description = $17,
    updated_at = $18
WHERE dealership_id = $1 RETURNING *;


-- name: GetDealershipById :one
SELECT * FROM dealership WHERE dealership_id = $1 LIMIT 1;

-- name: DeleteDealership :one
DELETE FROM dealership WHERE dealership_id = $1 RETURNING *;