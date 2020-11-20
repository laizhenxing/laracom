package repo

import (
	"github.com/jinzhu/gorm"

	pb "github.com/laizhenxing/laracom/user-service/proto/user"
)

type Repository interface {
	Create(user *pb.User) error
	Get(id string) (*pb.User, error)
	GetByEmail(email string) (*pb.User, error)
	GetAll() ([]*pb.User, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func (u *UserRepository) Create(user *pb.User) error {
	return u.DB.Create(&user).Error
}

func (u *UserRepository) Get(id string) (*pb.User, error) {
	user := &pb.User{}
	user.Id = id
	if err := u.DB.Unscoped().First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) GetByEmail(email string) (*pb.User, error) {
	user := &pb.User{}
	if err := u.DB.Debug().Unscoped().Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := u.DB.Debug().Unscoped().Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

