FROM golang:1.10 AS build-env

WORKDIR /go/src/github.com/murillio4/golang-microservice/consignment-service
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo main.go repository.go handler.go datastore.go

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/murillio4/golang-microservice/consignment-service /app/.
ENTRYPOINT ./consignment-service