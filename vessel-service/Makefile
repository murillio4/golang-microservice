all: build run

run:
	docker run -p 50052:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns vessel-service

build: build_proto build_go build_docker

build_proto:
	protoc -I. \
		--proto_path=$(GOPATH)/src/github.uio.no/mattisbr/golang-microservice/vessel-service \
		--go_out=. \
		--micro_out=. \
		proto/vessel/vessel.proto

build_go:
	GOOS=linux GOARCH=amd64 go build

build_docker:
	docker build -t vessel-service .