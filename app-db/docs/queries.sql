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


-- name: GetUserDealership :one
SELECT d.* FROM dealership d JOIN user_meta u ON d.dealership_id = u.dealership_id WHERE u.user_id = $1 LIMIT 1;

