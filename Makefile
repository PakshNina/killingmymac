# Variables
DOCKER_IMAGE_19 = kill-app
DOCKER_IMAGE_124 = normal-app
DOCKERFILE_19 = Dockerfile19
DOCKERFILE_124 = Dockerfile124
PORT = 8080

.PHONY: help build19 run19 build124 run124 clean clean19 clean124

# Default target
.DEFAULT_GOAL := help

help: ## Show this help message
	@echo "Available targets:"
	@echo "  build19   - Build Docker image using Go 1.9.7 (Dockerfile19)"
	@echo "  run19     - Run Docker container with Go 1.9.7 version"
	@echo "  build124  - Build Docker image using Go 1.24.1 (Dockerfile124)"
	@echo "  run124    - Run Docker container with Go 1.24.1 version"
	@echo "  clean19   - Remove Docker image for Go 1.9.7"
	@echo "  clean124  - Remove Docker image for Go 1.24.1"
	@echo "  clean     - Remove all Docker images"

build19: ## Build Docker image with Go 1.9.7
	docker build -t $(DOCKER_IMAGE_19) -f $(DOCKERFILE_19) .

run19: build19 ## Build and run Docker container with Go 1.9.7
	docker run --rm -p $(PORT):$(PORT) $(DOCKER_IMAGE_19)

build124: ## Build Docker image with Go 1.24.1
	docker build -t $(DOCKER_IMAGE_124) -f $(DOCKERFILE_124) .

run124: build124 ## Build and run Docker container with Go 1.24.1
	docker run --rm -p $(PORT):$(PORT) $(DOCKER_IMAGE_124)

clean19: ## Remove Docker image for Go 1.9.7
	docker rmi $(DOCKER_IMAGE_19) || true

clean124: ## Remove Docker image for Go 1.24.1
	docker rmi $(DOCKER_IMAGE_124) || true

clean: clean19 clean124 ## Remove all Docker images

