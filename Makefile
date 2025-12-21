.PHONY: build19 run19 build124 run124 benchmark

build19:
	docker build -t kill-app -f Dockerfile19 .

run19:
	docker run -p 8080:8080 kill-app

build124:
	docker build -t normal-app -f Dockerfile124 .

run124:
	docker run -p 8080:8080 normal-app

benchmark:
	go test -bench=. -benchmem ./recomendation/first/

