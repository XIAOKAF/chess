package serverService

import (
	"chess-room/server/dao"
	"chess-room/server/model"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("www..xyz.com")

func ParseToken(tokenString string) (*model.TokenClaims, error) {
	var tokenClaims model.TokenClaims
	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*model.TokenClaims)
	if !ok {
		return nil, errors.New("fail to parse token")
	}
	err = token.Claims.Valid()
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func GetToken(claims *model.TokenClaims) (string, error) {
	return dao.HashGet("token", claims.Mobile)
}
