#!/bin/sh

while getopts c:u: flag
do
  case "${flag}" in
  c) changelogname=${OPTARG};;
  u) username=${OPTARG};;
  esac
done

if [ -z $changelogname ]; then
    echo "changelog name not provided"
    exit 1
fi

if [ -z $username ]; then
    username=$USER
fi

current_time=$(date "+%Y_%m_%d_%H%M%S")

new_fileName="_"$changelogname"_changelog.sql"

converted="$(echo $new_fileName | sed 's/\([^A-Z]\)\([A-Z0-9]\)/\1_\2/g' \
                           | sed 's/\([A-Z0-9]\)\([A-Z0-9]\)\([^A-Z]\)/\1_\2\3/g' \
                           | tr '[:upper:]' '[:lower:]')"

echo "--liquibase formatted sql

--changeset $username:1

-- rollback " > ./docs/liquibase/schema/$current_time$converted

echo "Changelog created successfully : $current_time$converted"
