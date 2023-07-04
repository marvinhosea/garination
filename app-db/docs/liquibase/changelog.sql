--liquibase formatted sql

--changeSet oyamo:1
alter table "user"
    drop constraint user_pkey;

alter table "user"
    add primary key (owner, name, id);

create table dealership
(
    dealership_id varchar(100) not null primary key,
    owner_id varchar(100) not null,
    name varchar(100) not null,
    display_name varchar(100) not null,
    address varchar(100) not null,
    city varchar(100) not null,
    state varchar(100) not null,
    zip varchar(100) not null,
    phone varchar(100) not null,
    email varchar(100) not null,
    website varchar(100) not null,
    facebook_url varchar(100) not null,
    twitter_url varchar(100) not null,
    instagram_url varchar(100) not null,
    linkedin_url varchar(100) not null,
    logo_url varchar(100) not null,
    cover_url varchar(100) not null,
    description varchar(100) not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

CREATE TABLE "user_meta"
(
    user_meta_id VARCHAR(100) not null primary key,
    user_id VARCHAR(100) not null,
    facebook_url varchar(100) ,
    twitter_url varchar(100) ,
    instagram_url varchar(100) ,
    linkedin_url varchar(100) ,
    website_url varchar(100) ,
    dealership_id varchar(100) ,
    FOREIGN KEY (dealership_id) REFERENCES "dealership" (dealership_id)
);

--changeSet oyamo:2
ALTER TABLE "user_meta"
    ADD CONSTRAINT unique_id UNIQUE (user_id);

--changeSet oyamo:3
CREATE TABLE user_followers
(
    user_follower_id varchar(100) not null primary key,
    user_id varchar(100) not null,
    follower_id varchar(100) not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    FOREIGN KEY (user_id) REFERENCES "user_meta" (user_meta_id),
    FOREIGN KEY (follower_id) REFERENCES "user_meta" (user_meta_id)
);

ALTER TABLE "user_followers"
    ADD CONSTRAINT unique_id_followers UNIQUE (user_follower_id);

CREATE TABLE user_reviews
(
    user_review_id varchar(100) not null primary key,
    user_id varchar(100) not null,
    reviewer_id varchar(100) not null,
    rating int not null,
    review varchar(1024) not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    FOREIGN KEY (user_id) REFERENCES "user_meta" (user_meta_id),
    FOREIGN KEY (reviewer_id) REFERENCES "user_meta" (user_meta_id)
);

ALTER TABLE "user_reviews"
    ADD CONSTRAINT unique_id_reviews UNIQUE (user_review_id);

create table dealership_followers
(
    dealership_follower_id varchar(100) not null primary key,
    dealership_id varchar(100) not null,
    user_id varchar(100) not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    FOREIGN KEY (dealership_id) REFERENCES "dealership" (dealership_id),
    FOREIGN KEY (user_id) REFERENCES "user_meta" (user_meta_id)
);

ALTER TABLE "dealership_followers"
    ADD CONSTRAINT unique_id_dealership_followers UNIQUE (dealership_follower_id);


create table dealership_reviews
(
    dealership_review_id varchar(100) not null primary key,
    dealership_id varchar(100) not null,
    user_id varchar(100) not null,
    rating int not null,
    review varchar(1024) not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    FOREIGN KEY (dealership_id) REFERENCES "dealership" (dealership_id),
    FOREIGN KEY (user_id) REFERENCES "user_meta" (user_meta_id)
);

ALTER TABLE "dealership_reviews"
    ADD CONSTRAINT unique_id_dealership_reviews UNIQUE (dealership_review_id);

CREATE TABLE car_brands (
                            brand_id varchar(100) not null PRIMARY KEY,
                            name VARCHAR(50),
                            country VARCHAR(50),
                            logo_url VARCHAR(100),
                            created_at TIMESTAMP,
                            updated_at TIMESTAMP
);

ALTER TABLE "car_brands"
    ADD CONSTRAINT unique_id_car_brands UNIQUE (brand_id);


CREATE TABLE cars (
                      car_id varchar(100)  not null PRIMARY KEY,
                      brand_id varchar(100) not null,
                      category_id varchar(100) not null,
                      model VARCHAR(50),
                      year INT,
                      price DECIMAL(10, 2),
                      mileage INT,
                      color VARCHAR(20),
                      transmission VARCHAR(20),
                      fuel_type VARCHAR(20),
                      engine_capacity VARCHAR(10),
                      description TEXT,
                      dealership_id varchar(100),
                      dealer_id varchar(100),
                      created_at TIMESTAMP,
                      updated_at TIMESTAMP,
                      FOREIGN KEY (brand_id) REFERENCES car_brands(brand_id),
                      FOREIGN KEY (dealership_id) REFERENCES dealership(dealership_id),
                      FOREIGN KEY (dealer_id) REFERENCES "user_meta" (user_meta_id)
);

ALTER TABLE "cars"
    ADD CONSTRAINT unique_id_cars UNIQUE (car_id);

CREATE TABLE car_extra_features (
                                    car_extra_feature_id varchar(100) not null PRIMARY KEY,
                                    car_id varchar(100) not null,
                                    name VARCHAR(100),
                                    value VARCHAR(100),
                                    created_at TIMESTAMP,
                                    updated_at TIMESTAMP,
                                    FOREIGN KEY (car_id) REFERENCES cars(car_id)
);

ALTER TABLE "car_extra_features"
    ADD CONSTRAINT unique_id_car_extra_fx UNIQUE (car_extra_feature_id);

CREATE TABLE car_images (
                            car_image_id varchar(100) not null PRIMARY KEY,
                            car_id varchar(100) not null,
                            image_url VARCHAR(100),
                            created_at TIMESTAMP,
                            updated_at TIMESTAMP,
                            FOREIGN KEY (car_id) REFERENCES cars(car_id)
);

ALTER TABLE "car_images"
    ADD CONSTRAINT unique_id_car_images UNIQUE (car_image_id);

CREATE TABLE car_videos (
                            car_video_id varchar(100) not null PRIMARY KEY,
                            car_id varchar(100) not null,
                            video_url VARCHAR(100),
                            created_at TIMESTAMP,
                            updated_at TIMESTAMP,
                            FOREIGN KEY (car_id) REFERENCES cars(car_id)
);
ALTER TABLE "car_videos"
    ADD CONSTRAINT unique_id_car_videos UNIQUE (car_video_id);

CREATE TABLE car_reviews (
                             car_review_id varchar(100) not null PRIMARY KEY,
                             car_id varchar(100) not null,
                             user_id varchar(100) not null,
                             rating INT,
                             review TEXT,
                             created_at TIMESTAMP,
                             updated_at TIMESTAMP,
                             FOREIGN KEY (car_id) REFERENCES cars(car_id),
                             FOREIGN KEY (user_id) REFERENCES "user_meta" (user_meta_id)
);

ALTER TABLE "car_reviews"
    ADD CONSTRAINT unique_id_car_reviews UNIQUE (car_review_id);

CREATE TABLE car_ratings (
                             car_rating_id varchar(100) not null PRIMARY KEY,
                             car_id varchar(100) not null,
                             user_id varchar(100) not null,
                             rating INT,
                             created_at TIMESTAMP,
                             updated_at TIMESTAMP,
                             FOREIGN KEY (car_id) REFERENCES cars(car_id),
                             FOREIGN KEY (user_id) REFERENCES "user_meta" (user_meta_id)
);

ALTER TABLE "car_ratings"
    ADD CONSTRAINT unique_id_car_ratings UNIQUE (car_rating_id);


CREATE TABLE car_likes (
                           car_like_id varchar(100) not null PRIMARY KEY,
                           car_id varchar(100) not null,
                           user_id varchar(100) not null,
                           created_at TIMESTAMP,
                           updated_at TIMESTAMP,
                           FOREIGN KEY (car_id) REFERENCES cars(car_id),
                           FOREIGN KEY (user_id) REFERENCES "user_meta" (user_meta_id)
);

ALTER TABLE "car_likes"
    ADD CONSTRAINT unique_id_car_likes UNIQUE (car_like_id);


-- changeSet oyamo:4
ALTER TABLE "user_meta"
    ALTER COLUMN dealership_id DROP NOT NULL;

-- changeSet oyamo:5
ALTER TABLE "user_meta"
    ALTER COLUMN facebook_url DROP NOT NULL;

ALTER TABLE "user_meta"
    ALTER COLUMN twitter_url DROP NOT NULL;

ALTER TABLE "user_meta"
    ALTER COLUMN instagram_url DROP NOT NULL;

ALTER TABLE "user_meta"
    ALTER COLUMN linkedin_url DROP NOT NULL;


-- changeSet oyamo:6
ALTER TABLE "cars"
    DROP COLUMN category_id;

-- changeSet oyamo:7
ALTER TABLE "cars"
    ADD COLUMN title VARCHAR(100);


-- changeSet oyamo:8
ALTER TABLE "user_meta"
    ADD COLUMN deleted_at TIMESTAMP;

ALTER TABLE "dealership"
    ADD COLUMN deleted_at TIMESTAMP;

ALTER TABLE "dealership_followers"
    ADD COLUMN deleted_at TIMESTAMP;

ALTER TABLE "dealership_reviews"
    ADD COLUMN deleted_at TIMESTAMP;

ALTER TABLE "car_brands"
    ADD COLUMN deleted_at TIMESTAMP;

ALTER TABLE "cars"
    ADD COLUMN deleted_at TIMESTAMP;

ALTER TABLE "car_extra_features"
    ADD COLUMN deleted_at TIMESTAMP;

ALTER TABLE "car_images"
    ADD COLUMN deleted_at TIMESTAMP;

ALTER TABLE "car_videos"
    ADD COLUMN deleted_at TIMESTAMP;

ALTER TABLE "car_reviews"
    ADD COLUMN deleted_at TIMESTAMP;

ALTER TABLE "car_ratings"
    ADD COLUMN deleted_at TIMESTAMP;

ALTER TABLE "car_likes"
    ADD COLUMN deleted_at TIMESTAMP;

-- changeSet oyamo:9
ALTER TABLE "user_meta"
    ADD COLUMN created_at TIMESTAMP;

ALTER TABLE "user_meta"
    ADD COLUMN updated_at TIMESTAMP;

-- changeSet oyamo:10
ALTER TABLE "cars"
   ALTER COLUMN engine_capacity TYPE integer USING (engine_capacity::integer);

-- changeSet oyamo:11
ALTER TABLE "car_brands"
   ADD constraint unique_name_car_brands UNIQUE (name);


-- changeSet oyamo:12
ALTER TABLE "cars"
   ADD COLUMN status VARCHAR(20) DEFAULT 'pending';

ALTER TABLE "cars"
   ADD COLUMN is_featured BOOLEAN;

ALTER TABLE "cars"
    ADD COLUMN is_sold BOOLEAN;

ALTER TABLE "cars"
    ADD COLUMN horse_power INT;

ALTER TABLE "cars"
    ADD COLUMN torque INT;

ALTER TABLE "cars"
    ADD COLUMN torque_rpm INT;

ALTER TABLE "cars"
    ADD COLUMN safety_specifications varchar(100)[];

ALTER TABLE "cars"
    ADD COLUMN performance_specifications varchar(100)[];

ALTER TABLE "cars"
    ADD COLUMN comfort_specifications varchar(100)[];

-- changeSet oyamo:13
ALTER TABLE "cars"
    ADD COLUMN LOCATION VARCHAR(100);

-- changeSet oyamo:14
ALTER TABLE "cars"
    ADD COLUMN ownership VARCHAR(32);

-- changeSet oyamo:15
CREATE TABLE IF NOT EXISTS "spare_parts" (
    spare_part_id varchar(100) not null PRIMARY KEY,
    name VARCHAR(100),
    description TEXT,
    price INT,
    used BOOLEAN,
    car_model VARCHAR(100),
    car_brand VARCHAR(100),
    other_compatible_cars varchar(100)[],
    car_year int,
    is_universal BOOLEAN,
    category VARCHAR(100),
    part_number VARCHAR(100),
    dealership_id varchar(100) not null,
    dealer_id varchar(100) not null,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (dealership_id) REFERENCES dealership(dealership_id),
    FOREIGN KEY (dealer_id) REFERENCES "user_meta" (user_meta_id)
);

ALTER TABLE "spare_parts"
    ADD CONSTRAINT unique_id_spare_parts UNIQUE (spare_part_id);

CREATE TABLE IF NOT EXISTS "spare_part_images" (
    spare_part_image_id varchar(100) not null PRIMARY KEY,
    spare_part_id varchar(100) not null,
    image_url VARCHAR(100),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (spare_part_id) REFERENCES spare_parts(spare_part_id)
);

ALTER TABLE "spare_part_images"
    ADD CONSTRAINT unique_id_spare_part_images UNIQUE (spare_part_image_id);

CREATE TABLE IF NOT EXISTS "spare_part_reviews" (
    spare_part_review_id varchar(100) not null PRIMARY KEY,
    spare_part_id varchar(100) not null,
    user_id varchar(100) not null,
    rating INT,
    review TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (spare_part_id) REFERENCES spare_parts(spare_part_id),
    FOREIGN KEY (user_id) REFERENCES "user_meta" (user_meta_id)
);

ALTER TABLE "spare_part_reviews"
    ADD CONSTRAINT unique_id_spare_part_reviews UNIQUE (spare_part_review_id);

CREATE TABLE IF NOT EXISTS "spare_part_ratings" (
spare_part_rating_id varchar(100) not null PRIMARY KEY,
    spare_part_id varchar(100) not null,
    user_id varchar(100) not null,
    rating INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (spare_part_id) REFERENCES spare_parts(spare_part_id),
    FOREIGN KEY (user_id) REFERENCES "user_meta" (user_meta_id)
);


-- changeSet oyamo:16
ALTER TABLE "spare_parts"
    ALTER COLUMN price TYPE decimal(10,2) USING (price::decimal);

-- changeSet oyamo:17
ALTER TABLE "spare_parts"
    ALTER COLUMN dealer_id DROP NOT NULL;

ALTER TABLE "spare_parts"
    ALTER COLUMN dealership_id DROP NOT NULL;

-- changeSet oyamo:18
CREATE TABLE IF NOT EXISTS "dealership_kyc" (
    kyc_ref varchar(64) not null PRIMARY KEY,
    dealership_id varchar(100) not null,
    latitude DECIMAL(10, 8),
    longitude DECIMAL(10, 8),
    kra_cert_url varchar(126) DEFAULT NULL,
    business_permit_url varchar(126) DEFAULT NULL,
    owner_passport_url varchar(126) DEFAULT NULL,                --- front photo of ID / Passport
    owner_passport_back_url varchar(126) DEFAULT NULL,           --- back photo of ID
    verified BOOLEAN,
    verified_at TIMESTAMP default NULL,
    verified_by  varchar(126),                                   --- organisation name of the user that verified
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now(),
    FOREIGN KEY (dealership_id) REFERENCES dealership(dealership_id),
    CONSTRAINT unique_dealership_kyc_ref UNIQUE (kyc_ref)
);

-- changeSet oyamo:19
CREATE TABLE IF NOT EXISTS "dealer_kyc" (
    kyc_ref varchar(64) not null PRIMARY KEY,
    user_meta_id varchar(100) not null,
    kra_cert_url varchar(126) DEFAULT NULL,
    passport_url varchar(126),                                      --- front photo of pasport/ just ID
    passport_back_url varchar(126),                                --- back photo of passport
    verified BOOLEAN,
    verified_at TIMESTAMP default NULL,
    verified_by  varchar(126),                                   --- organisation name of the user that verified
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now(),
    FOREIGN KEY (user_meta_id) REFERENCES user_meta(user_meta_id),
    CONSTRAINT unique_dealer_kyc_ref UNIQUE (kyc_ref)
);

-- changeSet oyamo: 20
CREATE TABLE IF NOT EXISTS "car_kyc" (
    kyc_ref varchar(64) not null PRIMARY KEY,
    car_id varchar(100) not null,
    registration_number varchar(10) not null,
    log_book_url varchar(126) not null null,
    verified BOOLEAN,
    verified_at TIMESTAMP default NULL,
    verified_by  varchar(126),                                   --- organisation name of the user that verified
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now(),
    FOREIGN KEY (car_id) REFERENCES cars(car_id),
    CONSTRAINT unique_dealer_kyc_ref UNIQUE (kyc_ref)
)