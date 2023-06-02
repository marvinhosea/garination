PROJECT_NAME=car-marketplace-backend

build:
	docker-compose  -p $(PROJECT_NAME) build

up:
	docker-compose  -p $(PROJECT_NAME) up