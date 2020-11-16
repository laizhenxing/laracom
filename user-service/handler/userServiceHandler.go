package handler

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	pb "github.com/laizhenxing/laracom/user-service/proto/user"
	"github.com/laizhenxing/laracom/user-service/repo"
)

type UserService struct {
	Repo repo.Repository
}

func (u *UserService) Create(ctx context.Context, req *pb.User, res *pb.UserResponse) error {
	// 加密密码
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 16)
	if err != nil {
		return err
	}
	req.Password = string(hashedPass)
	if err := u.Repo.Create(req); err != nil {
		return err
	}
	res.User = req
	return nil
}

func (u *UserService) Get(ctx context.Context, req *pb.User, res *pb.UserResponse) error {
	user, err := u.Repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (u *UserService) GetAll(ctx context.Context, req *pb.UserRequest, res *pb.UserResponse) error {
	users, err := u.Repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}



