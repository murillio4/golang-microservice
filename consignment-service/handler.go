package main

import (
	"context"

	pb "github.com/murillio4/golang-microservice/consignment-service/proto/consignment"
	vpb "github.com/murillio4/golang-microservice/vessel-service/proto/vessel"
	log "github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
)

type service struct {
	session *mgo.Session
	vClient vpb.VesselServiceClient
}

func (s *service) GetRepo() Repository {
	return &ConsignmentRepository{s.session.Clone()}
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	vRes, err := s.vClient.FindAvailableVessel(context.Background(), &vpb.Spesification{
		MaxWeight: req.GetWeight(),
		Capacity:  int32(len(req.GetContainers())),
	})

	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"name": vRes.GetVessel().GetName(),
	}).Info("Found vessel")

	req.VesselId = vRes.GetVessel().GetId()

	err = repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	consignments, err := repo.GetAll()
	if err != nil {
		return err
	}

	res.Consignments = consignments
	return nil
}
