-- name: GetUserMeta :one
SELECT
    *
FROM
    user_meta
WHERE
        user_id = $1
AND deleted_at IS NULL
LIMIT
    1;
-- name: InsertUserMeta :one
INSERT INTO user_meta (
    user_meta_id, user_id, facebook_url,
    twitter_url, instagram_url, linkedin_url,
    website_url, dealership_id, created_at, updated_at
)
VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8, now(), now()) RETURNING *;
-- name: InsertUserMetaWithoutDealership :one
INSERT INTO user_meta (
    user_meta_id, user_id, facebook_url,
    twitter_url, instagram_url, linkedin_url,
    website_url
)
VALUES
    ($1, $2, $3, $4, $5, $6, $7) RETURNING *;
-- name: UpdateUserMeta :one
UPDATE
    user_meta
SET
    facebook_url = $2,
    twitter_url = $3,
    instagram_url = $4,
    linkedin_url = $5,
    website_url = $6,
    user_meta_id = $7,
    updated_at = now()
WHERE
        user_id = $1
RETURNING *;
-- name: GetUserDealership :one
SELECT
    d.*
FROM
    dealership d
        JOIN user_meta u ON d.dealership_id = u.dealership_id
WHERE
        u.user_id = $1
AND d.deleted_at IS NULL
LIMIT
    1;
-- name: UpdateUserDealership :one
UPDATE
    user_meta
SET
    dealership_id = $2,
    updated_at = now()
WHERE
        user_id = $1 RETURNING *;
-- name: InsertDealership :one
INSERT INTO dealership (
    dealership_id, owner_id, name, display_name,
    address, city, state, zip, phone, email,
    website, facebook_url, twitter_url,
    instagram_url, linkedin_url, logo_url,
    cover_url, description, created_at,
    updated_at
)
VALUES
    (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
        $11, $12, $13, $14, $15, $16, $17, $18,
        now(), now()
    ) RETURNING *;
-- name: UpdateDealership :one
UPDATE
    dealership
SET
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
    updated_at = now()
WHERE
        dealership_id = $1 RETURNING *;
-- name: GetDealershipById :one
SELECT
    *
FROM
    dealership
WHERE
        dealership_id = $1
AND deleted_at IS NULL
LIMIT
    1;
-- name: DeleteDealership :one
UPDATE dealership SET deleted_at = now(), updated_at = now() WHERE dealership_id = $1 RETURNING *;

-- name: CreateCarBrand :one
INSERT INTO car_brands (
    brand_id, name, country, logo_url, created_at,
    updated_at
)
VALUES
    ($1, $2, $3, $4, now(), now()) RETURNING *;
-- name: UpdateCarBrand :one
UPDATE
    car_brands
SET
    name = $2,
    logo_url = $3,
    country = $4,
    updated_at = now()
WHERE
        brand_id = $1 RETURNING *;
-- name: ListCarBrandsPaged :many
SELECT
    *
FROM
    car_brands
WHERE
        deleted_at IS NULL
ORDER BY
    name ASC
LIMIT
    $1 OFFSET $2;

-- name: GetCarBrandById :one
SELECT
    *
FROM
    car_brands
WHERE
        brand_id = $1
AND deleted_at IS NULL
LIMIT
    1;

-- name: DeleteCarBrand :one
UPDATE car_brands SET deleted_at = now(), updated_at = now() WHERE brand_id = $1 RETURNING *;

-- name: InsertExtraFeature :one
INSERT INTO car_extra_features (
    car_extra_feature_id, car_id, name, value , created_at,
    updated_at
) VALUES ($1, $2, $3, $4, now(), now()) RETURNING *;


-- name: UpdateExtraFeature :one
UPDATE
    car_extra_features
SET
    name = $2,
    value = $3,
    updated_at = now()
WHERE
        car_extra_feature_id = $1 RETURNING car_extra_feature_id;

-- name: ListExtraFeaturesForCar :many
SELECT
    *
FROM
    car_extra_features
WHERE
        car_id = $1
AND deleted_at IS NULL
ORDER BY
    name ASC;

-- name: DeleteExtraFeature :one
UPDATE car_extra_features SET deleted_at = now(), updated_at = now() WHERE car_extra_feature_id = $1 RETURNING *;


-- name: CreateCarImage :one
INSERT INTO car_images (
    car_image_id, car_id, image_url, created_at,
    updated_at
) VALUES ($1, $2, $3, now(), now()) RETURNING car_image_id;

-- name: ListCarImagesForCar :many
SELECT
    *
FROM
    car_images
WHERE
        car_id = $1
AND deleted_at IS NULL
ORDER BY
    created_at ASC;

-- name: UpdateCarImage :one
UPDATE
    car_images
SET
    image_url = $2,
    updated_at = now()
WHERE
        car_image_id = $1 RETURNING car_image_id;

-- name: DeleteCarImage :one
UPDATE car_images SET deleted_at = now() WHERE car_image_id = $1 RETURNING *;

-- name: InsertCar :one
INSERT INTO cars (
    car_id, brand_id, model,
    year, price, mileage, color, transmission,
    fuel_type, engine_capacity, description,
    dealership_id, dealer_id, comfort_specifications, safety_specifications, performance_specifications,
    horse_power, torque, torque_rpm ,ownership, title,
    created_at,updated_at
)
VALUES
    (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
        $11, $12, $13, $14, $15, $16, $17, $18, $19, $20,$21, now(), now()
    ) RETURNING *;

-- name: UpdateCar :one
UPDATE
    cars
SET
    brand_id = $2,
    model = $3,
    year = $4,
    price = $5,
    mileage = $6,
    color = $7,
    transmission = $8,
    fuel_type = $9,
    engine_capacity = $10,
    description = $11,
    dealership_id = $12,
    dealer_id = $13,
    status = $14,
    horse_power = $15,
    torque = $16,
    safety_specifications = $17,
    comfort_specifications = $18,
    ownership = $19,
    title = $20,
    updated_at = now()
WHERE
        car_id = $1 RETURNING car_id;

-- name: GetCarById :one
SELECT * FROM cars WHERE car_id = $1 AND deleted_at IS NULL LIMIT 1;

-- name: ListCarsPaged :many
SELECT * FROM cars WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2;

-- name: ListCarsByDealerPaged :many
SELECT * FROM cars WHERE dealer_id = $1 AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: ListCarsByDealershipPaged :many
SELECT * FROM cars WHERE dealership_id = $1 AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: ListCarsByBrandPaged :many
SELECT * FROM cars WHERE brand_id = $1 AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $2 OFFSET $3;


-- name: SearchCarsPaged :many
SELECT
    c.*
FROM
    cars c
    LEFT JOIN car_extra_features f ON c.car_id = f.car_id
    LEFT JOIN car_brands b ON c.brand_id = b.brand_id
WHERE
    (LOWER(c.model) LIKE LOWER($1)
    OR LOWER(c.description) LIKE LOWER($1)
    OR LOWER(c.color) LIKE LOWER($1)
    OR LOWER(c.transmission) LIKE LOWER($1)
    OR LOWER(c.fuel_type) LIKE LOWER($1)
    OR LOWER(f.name) LIKE LOWER($1)
    OR LOWER(f.value) LIKE LOWER($1)
    OR LOWER(b.name) LIKE LOWER($1)
    OR LOWER(b.country) LIKE LOWER($1))
    AND c.deleted_at IS NULL
ORDER BY f.created_at DESC LIMIT $2 OFFSET $3;

-- name: CarByDealerCount :one
SELECT COUNT(*) FROM cars WHERE dealer_id = $1 AND deleted_at IS NULL;

-- name: CarByDealershipCount :one
SELECT COUNT(*) FROM cars WHERE dealership_id = $1 AND deleted_at IS NULL;

-- name: CarByBrandCount :one
SELECT COUNT(*) FROM cars WHERE brand_id = $1 AND deleted_at IS NULL;

-- name: CarsByDealerPaged :many
SELECT * FROM cars WHERE dealer_id = $1  AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: CarsByDealershipPaged :many
SELECT * FROM cars WHERE dealership_id = $1 AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: DeleteCar :one
UPDATE cars SET deleted_at = now() , updated_at = now() WHERE car_id = $1 RETURNING car_id;

-- name: InsertSparePart :one
INSERT INTO spare_parts (
    spare_part_id,name,description,price,used,
    car_model,car_brand,other_compatible_cars,
    car_year,is_universal,category,
    part_number,dealership_id,dealer_id,created_at,
    updated_at,deleted_at
) VALUES (
    $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    NULL
) RETURNING spare_part_id;


-- name: UpdateSparePart :one
UPDATE
    spare_parts
SET
    name = $2,description = $3,price = $4,used = $5,
    car_model = $6,car_brand = $7,other_compatible_cars = $8,
    car_year = $9,is_universal = $10,category = $11,part_number = $12,
    updated_at = now()
WHERE
        spare_part_id = $1 RETURNING spare_part_id;

-- name: GetSparePartById :one
SELECT * FROM spare_parts WHERE spare_part_id = $1 AND deleted_at IS NULL LIMIT 1;


-- name: ListSparePartsPaged :many
SELECT * FROM spare_parts WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2;


-- name: ListSparePartsByDealerPaged :many
SELECT * FROM spare_parts WHERE dealer_id = $1 AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: ListSparePartsByDealershipPaged :many
SELECT * FROM spare_parts WHERE dealership_id = $1 AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: SearchSparePartsPaged :many
SELECT
    s.*
FROM
    spare_parts s
WHERE
    (LOWER(s.name) LIKE LOWER($1)
    OR LOWER(s.description) LIKE LOWER($1)
    OR LOWER(s.car_model) LIKE LOWER($1)
    OR LOWER(s.car_brand) LIKE LOWER($1)
    OR LOWER(s.other_compatible_cars) LIKE LOWER($1)
    OR LOWER(s.category) LIKE LOWER($1)
    OR LOWER(s.part_number) LIKE LOWER($1))
    AND s.deleted_at IS NULL ORDER BY s.created_at DESC LIMIT $2 OFFSET $3;

-- name: SparePartsByDealerCount :one
SELECT COUNT(*) FROM spare_parts WHERE dealer_id = $1 AND deleted_at IS NULL;

-- name: SparePartsByDealershipCount :one
SELECT COUNT(*) FROM spare_parts WHERE dealership_id = $1 AND deleted_at IS NULL;

-- name: FilterSparePartsByCategory :many
SELECT * FROM spare_parts WHERE category = $1 AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: FilterSparePartsByBrand :many
SELECT * FROM spare_parts WHERE car_brand = $1 AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: FilterSparePartsByModel :many
SELECT * FROM spare_parts WHERE car_model = $1 AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: DeleteSparePart :one
UPDATE spare_parts SET deleted_at = now() , updated_at = now() WHERE spare_part_id = $1 RETURNING spare_part_id;

-- name: InsertSparePartImage :one
INSERT INTO spare_part_images (
    spare_part_image_id,spare_part_id,image_url,
    created_at,updated_at,deleted_at
) VALUES (
    $1,$2,$3,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP,NULL
) RETURNING spare_part_image_id;

-- name: UpdateSparePartImage :one
UPDATE
    spare_part_images
SET
    spare_part_id = $2,image_url = $3,updated_at = now()
WHERE
        spare_part_image_id = $1 RETURNING spare_part_image_id;

-- name: GetSparePartImageById :one
SELECT * FROM spare_part_images WHERE spare_part_image_id = $1 AND deleted_at IS NULL LIMIT 1;

-- name: ListSparePartImagesPaged :many
SELECT * FROM spare_part_images WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2;

-- name: ListSparePartImagesBySparePartPaged :many
SELECT * FROM spare_part_images WHERE spare_part_id = $1 AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $2 OFFSET $3;


-- name: DeleteSparePartImage :one
UPDATE spare_part_images SET deleted_at = now() , updated_at = now() WHERE spare_part_image_id = $1 RETURNING spare_part_image_id;