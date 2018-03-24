all: build run

run:
	docker-compose run user-service

build:
	protoc -I. \
		--proto_path=$(GOPATH)/src/github.com/murillio4/golang-microservice/user-service \
		--go_out=. \
		--micro_out=. \
		proto/user/user.proto
