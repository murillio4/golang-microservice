package main

import (
	"os"

	micro "github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"

	pb "github.com/murillio4/golang-microservice/vessel-service/proto/vessel"
)

const (
	defaultHost = "localhost:27017"
)

func main() {

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)
	defer session.Close()

	if err != nil {
		log.WithFields(log.Fields{
			"err":  err,
			"host": host,
		}).Panic("Could not connect to datastore with host")
	}

	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)
	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), &service{session})

	if err := srv.Run(); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("Failed to serve")
	}
}
