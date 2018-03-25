package main

import (
	micro "github.com/micro/go-micro"
	pb "github.com/murillio4/golang-microservice/user-service/proto/user"
	log "github.com/sirupsen/logrus"
)

func main() {
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.WithError(err).Fatal("Could not connect to db")
	}

	db.AutoMigrate(&pb.User{})
	repo := &UserRepository{db}
	token := &TokenService{repo}

	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, token})

	if err := srv.Run(); err != nil {
		log.WithError(err).Fatal("Failed to serve")
	}
}
