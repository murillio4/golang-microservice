package main

import (
	context "context"
	"os"

	microclient "github.com/micro/go-micro/client"
	pb "github.com/murillio4/golang-microservice/user-service/proto/user"
	log "github.com/sirupsen/logrus"
)

func main() {
	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	name := "Ewan Valentine"
	email := "ewan.valentine89@gmail.com"
	password := "test123"
	company := "BBC"

	r, err := client.Create(context.TODO(), &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})
	if err != nil {
		log.WithError(err).Fatal("Could not create")
	}
	log.WithField("user", r.GetUser().GetId()).Info("Created")

	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.WithError(err).Fatal("Could not list users")
	}
	for _, v := range getAll.Users {
		log.Info(v)
	}

	authResponse, err := client.Auth(context.TODO(), &pb.User{
		Email:    email,
		Password: password,
	})

	if err != nil {
		log.WithError(err).WithField("email", email).Fatal("Could not authenticate user")
	}

	log.WithField("token", authResponse.GetToken()).Info("Acces granted")

	// let's just exit because
	os.Exit(0)
}
