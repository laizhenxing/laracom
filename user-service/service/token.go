package service

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	pb "github.com/laizhenxing/laracom/user-service/proto/user"
	"github.com/laizhenxing/laracom/user-service/repo"
)

var (
	key = []byte(os.Getenv("JWT_SECRET"))
)

type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type Authable interface {
	Encode(user *pb.User) (string, error)
	Decode(tokenStr string) (*CustomClaims, error)
}

type TokenService struct {
	Repo repo.Repository
}

func (t *TokenService) Encode(user *pb.User) (string, error) {
	// 设置过期时间
	expire := time.Now().Add(time.Hour * 72).Unix()
	// create the claims
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expire,
			Issuer:    "laracom.user.service",
		},
	}
	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// sign token and return
	return token.SignedString(key)
}

func (t *TokenService) Decode(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	// 验证 token
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

