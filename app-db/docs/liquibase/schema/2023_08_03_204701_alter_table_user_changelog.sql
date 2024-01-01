--liquibase formatted sql

--changeset marvin:1
    alter table "user" drop constraint user_pkey;
-- rollback
    ALTER TABLE "user" ADD CONSTRAINT user_pkey;
