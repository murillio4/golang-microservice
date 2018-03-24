all: build run

run:
	docker-compose run consignment-service

build:
	protoc -I. \
		--proto_path=$(GOPATH)/src/github.com/murillio4/golang-microservice/consignment-service \
		--go_out=. \
		--micro_out=. \
		proto/consignment/consignment.proto