PROJECT_NAME=car-marketplace-backend

build:
	docker-compose  -p $(PROJECT_NAME) build
	# remove dangling images
	docker rmi -f $$(docker images -f "dangling=true" -q)

up:
	docker-compose  -p $(PROJECT_NAME) up --remove-orphans --build
