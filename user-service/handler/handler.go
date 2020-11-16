package handler

import (
	"context"
	"errors"
	"github.com/laizhenxing/laracom/user-service/util/e"
	"log"

	"golang.org/x/crypto/bcrypt"

	pb "github.com/laizhenxing/laracom/user-service/proto/user"
	"github.com/laizhenxing/laracom/user-service/repo"
	"github.com/laizhenxing/laracom/user-service/service"
)

type UserService struct {
	Repo  repo.Repository
	Token service.Authable
}

func (u *UserService) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	log.Println("Logging in withL: ", req.Email, req.Password)
	// 获取用户信息
	user, err := u.Repo.GetByEmail(req.Email)
	if err != nil {
		return err
	}
	// 校验用户输入的密码是否与数据库密码匹配
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	// 生成token
	token, err := u.Token.Encode(user)
	if err != nil {
		return err
	}
	res.Token = token
	return nil
}

func (u *UserService) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	// 校验用户请求中的token是否有效
	claims, err := u.Token.Decode(req.Token)
	if err != nil {
		return err
	}

	if claims.User.Id == "" {
		return errors.New(e.GetMsg(e.InvalidUser))
	}

	res.Valid = true
	return nil
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
