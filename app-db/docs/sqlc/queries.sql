-- name: getUserMeta :one
SELECT * FROM user_meta WHERE user_id = $1 LIMIT 1;

-- name: insertUserMeta :one
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

-- name: updateUserMeta :one
UPDATE user_meta SET
    facebook_url = $2,
    twitter_url = $3,
    instagram_url = $4,
    linkedin_url = $5,
    website_url = $6
WHERE user_id = $1 RETURNING *;

-- name: getUserDealership :one
SELECT d.* FROM dealership d JOIN user_meta u ON d.dealership_id = u.dealership_id WHERE u.user_id = $1 LIMIT 1;

-- name: insertDealership :one
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

