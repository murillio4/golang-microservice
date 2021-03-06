package main

import (
	"context"

	pb "github.com/murillio4/golang-microservice/vessel-service/proto/vessel"
	mgo "gopkg.in/mgo.v2"
)

type service struct {
	session *mgo.Session
}

func (s *service) GetRepo() Repository {
	return &VesselRepository{s.session.Clone()}
}

func (s *service) CreateVessel(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	if err := repo.Create(req); err != nil {
		return err
	}

	res.Vessel = req
	res.Created = true
	return nil
}

func (s *service) FindAvailableVessel(ctx context.Context, req *pb.Spesification, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	vessel, err := repo.FindAvailable(req)
	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}
