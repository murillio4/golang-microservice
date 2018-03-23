package main

import (
	context "context"

	micro "github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"

	pb "github.com/murillio4/golang-microservice/consignment-service/proto/consignment"
	vpb "github.com/murillio4/golang-microservice/vessel-service/proto/vessel"
)

// Repository repository interface
type Repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// ConsignmentRepository - Dummy repository, will simulate a datastore
type ConsignmentRepository struct {
	consignments []*pb.Consignment
}

//Create create new consignment
func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

//GetAll get all consignments
func (repo *ConsignmentRepository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// service should implement al of the methods to satisfy the service
// we defined in the protobuf definition
type service struct {
	repo    Repository
	vClient vpb.VesselServiceClient
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	vRes, err := s.vClient.FindAvailable(context.Background(), &vpb.Spesification{
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

	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	res.Consignments = s.repo.GetAll()
	return nil
}

func main() {
	repo := &ConsignmentRepository{}
	log.Info("Failed to serve")
	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vClient := vpb.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo, vClient})

	if err := srv.Run(); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("Failed to serve")
	}
}
