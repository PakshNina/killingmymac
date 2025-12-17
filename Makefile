# Variables
DOCKER_IMAGE_19 = kill-app
DOCKER_IMAGE_112 = normal-app
DOCKERFILE = Dockerfile
GO_VERSION_19 = 1.9.7
GO_VERSION_112 = 1.12.1
PORT = 8080
MEMORY_LIMIT = 500m

.PHONY: build19 run19 build112 run112 clean clean19 clean112

build19: ## Build Go 1.9.7
	docker build -t $(DOCKER_IMAGE_19) \
		--build-arg GO_VERSION=$(GO_VERSION_19) \
		-f $(DOCKERFILE) .

run19: build19 ## Build and run with Go 1.9.7
	docker run --rm -p $(PORT):$(PORT) --memory=$(MEMORY_LIMIT) $(DOCKER_IMAGE_19)

build112: ## Build Go 1.12.1
	docker build -t $(DOCKER_IMAGE_112) \
		--build-arg GO_VERSION=$(GO_VERSION_112) \
		-f $(DOCKERFILE) .

run112: build112 ## Build and run with Go 1.12.1
	docker run --rm -p $(PORT):$(PORT) --memory=$(MEMORY_LIMIT) $(DOCKER_IMAGE_112)

clean19: ## Remove Docker image for Go 1.9.7
	docker rmi $(DOCKER_IMAGE_19) || true

clean112: ## Remove Docker image for Go 1.12.1
	docker rmi $(DOCKER_IMAGE_112) || true

clean: clean19 clean112 ## Remove all Docker images
