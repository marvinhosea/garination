-- Database schema for the app-db
-- Tables

-- user-meta-meta
CREATE TABLE "user_meta"
(
    user_meta_id VARCHAR(100) not null primary key,
    user_id VARCHAR(100) not null,
    facebook_url varchar(100) ,
    twitter_url varchar(100) ,
    instagram_url varchar(100) ,
    linkedin_url varchar(100) ,
    website_url varchar(100) ,
    dealership_id VARCHAR(100) DEFAULT NULL,
    FOREIGN KEY (dealership_id) REFERENCES "dealership" (dealership_id)
);

-- dealership
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


-- user-followers
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

-- user-reviews
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

-- dealership followers
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


-- dealership reviews
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

-- Create the CarBrands table
CREATE TABLE car_brands (
   brand_id varchar(100) not null PRIMARY KEY,
   name VARCHAR(50),
   country VARCHAR(50),
   logo_url VARCHAR(100),
   created_at TIMESTAMP,
   updated_at TIMESTAMP
);



-- Create the Cars table
CREATE TABLE cars (
      car_id varchar(100)  not null PRIMARY KEY,
      brand_id varchar(100) not null,
      model VARCHAR(50),
      year INT,
      price DECIMAL(10, 2),
      mileage INT,
      color VARCHAR(20),
      transmission VARCHAR(20),
      fuel_type VARCHAR(20),
      engine_capacity int,
      description TEXT,
      dealership_id varchar(100),
      dealer_id varchar(100),
      created_at TIMESTAMP,
      updated_at TIMESTAMP,
      status                     varchar(20),
      title                      varchar(100),
      is_featured                boolean,
      is_sold                    boolean,
      horse_power                integer,
      torque                     integer,
      torque_rpm                 integer,
      safety_specifications      varchar[],
      performance_specifications varchar[],
      comfort_specifications     varchar[],
      location                   varchar(100),
      ownership                  varchar(32),
      FOREIGN KEY (brand_id) REFERENCES car_brands(brand_id),
      FOREIGN KEY (dealership_id) REFERENCES dealership(dealership_id),
      FOREIGN KEY (dealer_id) REFERENCES "user_meta" (user_meta_id)
);

-- create car extra features table
CREATE TABLE car_extra_features (
    car_extra_feature_id varchar(100) not null PRIMARY KEY,
    car_id varchar(100) not null,
    name VARCHAR(100),
    value VARCHAR(100),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (car_id) REFERENCES cars(car_id)
);

-- create car images table
CREATE TABLE car_images (
    car_image_id varchar(100) not null PRIMARY KEY,
    car_id varchar(100) not null,
    image_url VARCHAR(100),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (car_id) REFERENCES cars(car_id)
);

-- create car videos table
CREATE TABLE car_videos (
    car_video_id varchar(100) not null PRIMARY KEY,
    car_id varchar(100) not null,
    video_url VARCHAR(100),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (car_id) REFERENCES cars(car_id)
);


-- create car reviews table
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

-- create car ratings table
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


-- create car likes table
CREATE TABLE car_likes (
    car_like_id varchar(100) not null PRIMARY KEY,
    car_id varchar(100) not null,
    user_id varchar(100) not null,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (car_id) REFERENCES cars(car_id),
    FOREIGN KEY (user_id) REFERENCES "user_meta" (user_meta_id)
);

-- spareparts table
CREATE TABLE IF NOT EXISTS "spare_parts" (
    spare_part_id varchar(100) not null PRIMARY KEY,
    name VARCHAR(100),
    description TEXT,
    price DECIMAL(10, 2),
    used BOOLEAN,
    car_model VARCHAR(100),
    car_brand VARCHAR(100),
    other_compatible_cars varchar(100)[],
    car_year int,
    is_universal BOOLEAN,
    category VARCHAR(100),
    part_number VARCHAR(100),
    dealership_id varchar(100) default null,
    dealer_id varchar(100) default null,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (dealership_id) REFERENCES dealership(dealership_id),
    FOREIGN KEY (dealer_id) REFERENCES "user_meta" (user_meta_id)
    );


-- spare parts images
CREATE TABLE IF NOT EXISTS "spare_part_images" (
    spare_part_image_id varchar(100) not null PRIMARY KEY,
    spare_part_id varchar(100) not null,
    image_url VARCHAR(100),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (spare_part_id) REFERENCES spare_parts(spare_part_id)
    );


-- spare parts reviews
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


-- spare parts ratings
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
