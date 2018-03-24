FROM golang:1.10 AS build-env

WORKDIR /go/src/github.com/murillio4/golang-microservice/user-service
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/murillio4/golang-microservice/user-service /app/.
ENTRYPOINT ./user-service