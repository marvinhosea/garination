-- Database schema for the app-db
-- Tables

-- user-meta
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
    FOREIGN KEY (user_id) REFERENCES "user" (user_id),
    FOREIGN KEY (dealership_id) REFERENCES "dealership" (dealership_id)
);

-- dealership
create table dealership
(
    dealership_id varchar(100) not null primary key,
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