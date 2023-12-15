DOCKER_COMPOSE_FILE := docker-compose.yml
CONTAINER_NAME := mongo-go-todos


run:
	@go run cmd/main.go

setup: 
	sudo docker-compose up -d

stop-setup:
	@sudo docker stop $(CONTAINER_NAME) || true
	@sudo docker rm $(CONTAINER_NAME) || true

restart-setup: stop-setup setup

