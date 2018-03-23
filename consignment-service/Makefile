all: build run

run:
	docker run -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns consignment-service

build: build_proto build_go build_docker

build_proto:
	protoc -I. \
		--proto_path=$(GOPATH)/src/github.uio.no/mattisbr/golang-microservice/consignment-service \
		--go_out=. \
		--micro_out=. \
		proto/consignment/consignment.proto

build_go:
	GOOS=linux GOARCH=amd64 go build

build_docker:
	docker build -t consignment-service .