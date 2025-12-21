.PHONY: build19 run19 build124 run124 benchmark_1

build19:
	docker build -t kill-app -f Dockerfile19 .

run19:
	docker run -p 8080:8080 kill-app

build124:
	docker build -t normal-app -f Dockerfile124 .

run124:
	docker run -p 8080:8080 normal-app

benchmark_1:
	go test -bench=. -benchmem ./recomendation/first/

benchmark_2:
	go test -bench=. -benchmem ./recomendation/second/

benchmark: benchmark_1 benchmark_2

reflect_2:
	go test ./recomendation/second -run=^$ -bench=^$ -gcflags="-m"