package main

import (
	pb "github.com/murillio4/golang-microservice/vessel-service/proto/vessel"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName           = "shippy"
	vesselCollection = "vessels"
)

type Repository interface {
	FindAvailable(*pb.Spesification) (*pb.Vessel, error)
	Create(*pb.Vessel) error

	Close()
}

//VesselRepository dummy vessel repository
type VesselRepository struct {
	session *mgo.Session
}

//Create create a new vessel
func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}

//FindAvailable find available vessels
func (repo *VesselRepository) FindAvailable(spec *pb.Spesification) (*pb.Vessel, error) {
	var vessel *pb.Vessel

	err := repo.collection().Find(bson.M{
		"capacity":  bson.M{"$gte": spec.GetCapacity()},
		"maxweight": bson.M{"$gte": spec.GetMaxWeight()},
	}).One(&vessel)

	if err != nil {
		return nil, err
	}

	return vessel, nil
}

//Close closes db session
func (repo *VesselRepository) Close() {
	repo.session.Close()
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}
