package main

import (
	context "context"
	"errors"

	pb "github.com/murillio4/golang-microservice/user-service/proto/user"
)

type service struct {
	repo Repository
}

func (srv *service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	if err := srv.repo.Create(req); err != nil {
		return err
	}

	res.User = req
	return nil
}

func (srv *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := srv.repo.Get(req)
	if err != nil {
		return err
	}

	res.User = user
	return nil
}

func (srv *service) GetAll(ctx context.Context, req *pb.User, res *pb.Response) error {
	users, err := srv.repo.GetAll()
	if err != nil {
		return err
	}

	res.Users = users
	return nil
}

func (srv *service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	if req.GetEmail() == "" || req.GetPassword() == "" {
		return errors.New("Only email and password")
	}

	_, err := srv.repo.Get(req)
	if err != nil {
		return err
	}

	res.Token = "huhu"
	return nil
}

func (srv *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}
