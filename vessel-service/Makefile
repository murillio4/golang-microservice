all: build run

run:
	docker-compose run vessel-service

build:
	protoc -I. \
		--proto_path=$(GOPATH)/src/github.com/murillio4/golang-microservice/vessel-service \
		--go_out=. \
		--micro_out=. \
		proto/vessel/vessel.proto
