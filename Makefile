all: prep build

LISTEN_PORT = "8081"

prep:
	@export HTTP_PORT=${HTTP_PORT}

build:
	@echo "Creating a Docker image"
	@docker rmi -f hello-world
	@export HTTP_PORT=${HTTP_PORT}
	docker build -t hello-world --build-arg LISTEN_PORT=${LISTEN_PORT} .
