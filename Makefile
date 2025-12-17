# Variables
DOCKER_IMAGE_19 = kill-app
DOCKER_IMAGE_124 = normal-app
DOCKERFILE = Dockerfile
GO_VERSION_19 = 1.9.7
GO_VERSION_124 = 1.24.1
PORT = 8080
MEMORY_LIMIT = 500Mb

.PHONY: build19 run19 build124 run124 clean clean19 clean124

build19: ## Build Go 1.9.7
	docker build -t $(DOCKER_IMAGE_19) \
		--build-arg GO_VERSION=$(GO_VERSION_19) \
		-f $(DOCKERFILE) .

run19: build19 ## Build and run with Go 1.9.7
	docker run --rm -p $(PORT):$(PORT) --memory=$(MEMORY_LIMIT) $(DOCKER_IMAGE_19)

build124: ## Build Go 1.24.1
	docker build -t $(DOCKER_IMAGE_124) \
		--build-arg GO_VERSION=$(GO_VERSION_124) \
		-f $(DOCKERFILE) .

run124: build124 ## Build and run with Go 1.24.1
	docker run --rm -p $(PORT):$(PORT) --memory=$(MEMORY_LIMIT) $(DOCKER_IMAGE_124)

clean19: ## Remove Docker image for Go 1.9.7
	docker rmi $(DOCKER_IMAGE_19) || true

clean124: ## Remove Docker image for Go 1.24.1
	docker rmi $(DOCKER_IMAGE_124) || true

clean: clean19 clean124 ## Remove all Docker images
