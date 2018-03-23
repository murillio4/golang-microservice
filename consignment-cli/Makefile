all: build run

run: 
	docker run -e MICRO_REGISTRY=mdns consignment-cli

build:
	docker build -t consignment-cli .