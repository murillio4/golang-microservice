package main

import (
	context "context"

	"golang.org/x/crypto/bcrypt"

	pb "github.com/murillio4/golang-microservice/user-service/proto/user"
)

type service struct {
	repo         Repository
	tokenService Authable
}

func (srv *service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashed)

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

func (srv *service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := srv.repo.GetAll()
	if err != nil {
		return err
	}

	res.Users = users
	return nil
}

func (srv *service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {

	user, err := srv.repo.GetByEmail(req.GetEmail())
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, err := srv.tokenService.Encode(user)
	if err != nil {
		return err
	}

	res.Token = token
	return nil
}

func (srv *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	_, err := srv.tokenService.Decode(req.GetToken())
	return err
}
