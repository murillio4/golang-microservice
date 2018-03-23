package main

import (
	pb "github.com/murillio4/golang-microservice/consignment-service/proto/consignment"
	"gopkg.in/mgo.v2"
)

const (
	dbName                = "shippy"
	consignmentCollection = "consignments"
)

// Repository repository interface
type Repository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)

	Close()
}

// ConsignmentRepository - Dummy repository, will simulate a datastore
type ConsignmentRepository struct {
	session *mgo.Session
}

//Create create new consignment
func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) error {
	return repo.collection().Insert(consignment)
}

//GetAll get all consignments
func (repo *ConsignmentRepository) GetAll() ([]*pb.Consignment, error) {
	var consigments []*pb.Consignment

	err := repo.collection().Find(nil).All(consigments)
	return consigments, err
}

//Close closes db
func (repo *ConsignmentRepository) Close() {
	repo.session.Close()
}

func (repo *ConsignmentRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(consignmentCollection)
}
