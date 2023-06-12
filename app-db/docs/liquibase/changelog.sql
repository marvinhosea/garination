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





