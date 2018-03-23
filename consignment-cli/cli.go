package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	pb "github.com/murillio4/golang-microservice/consignment-service/proto/consignment"
	log "github.com/sirupsen/logrus"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)

	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	cmd.Init()

	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("Could not parse file")
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("Could not greet")
	}

	log.WithFields(log.Fields{
		"con":    r.Consignment,
		"status": r.Created,
	}).Info("Created new consignment")

	r, err = client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("Could not greet")
	}

	log.WithFields(log.Fields{
		"con":    r.Consignment,
		"status": r.Created,
	}).Info("Created new consignment")
}
