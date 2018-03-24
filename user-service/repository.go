package main

import (
	"github.com/jinzhu/gorm"
	pb "github.com/murillio4/golang-microservice/user-service/proto/user"
)

type Repository interface {
	Create(user *pb.User) error
	GetAll() ([]*pb.User, error)
	Get(user *pb.User) (*pb.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) Create(user *pb.User) error {
	return repo.db.Create(user).Error
}

func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *UserRepository) Get(user *pb.User) (*pb.User, error) {
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
