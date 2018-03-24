package main

import (
	context "context"
	"os"

	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	pb "github.com/murillio4/golang-microservice/user-service/proto/user"
	log "github.com/sirupsen/logrus"
)

func main() {
	cmd.Init()

	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "Your full name",
			},
			cli.StringFlag{
				Name:  "email",
				Usage: "Your email",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "Your password",
			},
			cli.StringFlag{
				Name:  "company",
				Usage: "Your company",
			},
		),
	)

	service.Init(
		micro.Action(func(c *cli.Context) {
			name := c.String("name")
			email := c.String("email")
			password := c.String("password")
			company := c.String("company")
			log.Infof("%s %s %s %s \n", name, email, password, company)
			r, err := client.Create(context.Background(), &pb.User{
				Name:     name,
				Email:    email,
				Password: password,
				Company:  company,
			})
			if err != nil {
				log.WithError(err).Fatal("Could not create user")
			}

			log.WithField("user", r).Info("Created: ")
			getAll, err := client.GetAll(context.Background(), &pb.Request{})
			if err != nil {
				log.WithError(err).Fatal("Could not list users")
			}

			for _, v := range getAll.GetUsers() {
				log.WithField("user", v).Info("User: ")
			}

			os.Exit(0)
		}),
	)

	if err := service.Run(); err != nil {
		log.WithError(err).Fatal("Could not run service")
	}
}
