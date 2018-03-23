package main

import (
	"context"
	"errors"

	micro "github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"

	pb "github.com/murillio4/golang-microservice/vessel-service/proto/vessel"
)

type Repository interface {
	FindAvailable(*pb.Spesification) (*pb.Vessel, error)
}

//VesselRepository dummy vessel repository
type VesselRepository struct {
	vessels []*pb.Vessel
}

//FindAvailable find available vessels from the repository
func (repo *VesselRepository) FindAvailable(spec *pb.Spesification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.GetCapacity() <= vessel.GetCapacity() && spec.GetMaxWeight() <= vessel.GetMaxWeight() {
			return vessel, nil
		}
	}

	return nil, errors.New("No vessel found by that spec")
}

type service struct {
	repo Repository
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Spesification, res *pb.Response) error {
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}

func main() {
	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}

	repo := &VesselRepository{vessels}

	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)
	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("Failed to serve")
	}
}
