FROM golang:1.10 AS build-env

WORKDIR /go/src/github.uio.no/mattisbr/golang-microservice/consignment-service
COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.uio.no/mattisbr/golang-microservice/consignment-service /app/.
ENTRYPOINT ./main