.PHONY: build19 run19 build124 run124 benchmark benchmark_equal

build19:
	docker build --build-arg GO_VERSION=1.9.1 -t kill-app -f Dockerfile .

run19:
	docker run --rm -p 8080:8080 --memory=500m kill-app

build124:
	docker build --build-arg GO_VERSION=1.24.1 -t normal-app -f Dockerfile .

run124:
	docker run --rm -p 8080:8080 --memory=500m normal-app

benchmark:
	go test -bench=. -benchmem -benchtime=50000x ./recomendation/
