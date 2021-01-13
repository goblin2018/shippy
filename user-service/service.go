package main

import (
	"github.com/dgrijalva/jwt-go"
	pb "github.com/goblin2018/shippy/user-service/proto/user"
)

var (
	key = []byte("mySecretKey")
)

type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type TokenService struct {
	repo Repository
}

func (srv *TokenService) Decode(token string) (*CustomClaims, error) {
	tokenType, err := jwt.ParseWithClaims(string(key), &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if claims, ok := tokenType.Claims.(*CustomClaims); ok && tokenType.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (srv *TokenService) Encode(user *pb.User) (string, error) {
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "shippy.service.user",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
