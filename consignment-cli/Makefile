all: build run

run: 
	docker run -e MICRO_REGISTRY=mdns consignment-cli

build: build_go build_docker

build_go:
	GOOS=linux GOARCH=amd64 go build

build_docker:
	docker build -t consignment-cli .