all: prep build

HTTP_PORT = 8080

prep:
	@export HTTP_PORT=${HTTP_PORT}

build:
	@echo "Creating a Docker image"
	@docker rmi -f hello-world
	docker build -t hello-world .
