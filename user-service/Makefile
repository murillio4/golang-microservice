all: build run

run:
	docker run -p 50053:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns user-service

build: build_proto build_docker

build_proto:
	protoc -I. \
		--proto_path=$(GOPATH)/src/github.com/murillio4/golang-microservice/user-service \
		--go_out=. \
		--micro_out=. \
		proto/user/user.proto

build_docker:
	docker build -t user-service .