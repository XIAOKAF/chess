package server_service

import (
	"chess/server/dao"
	"chess/server/model"
	"chess/service"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("www..xyz.com")

func CreateToken(mobile string, duration time.Duration) (error, string) {
	expireTime := time.Now().Add(duration * time.Minute)
	claims := model.TokenClaims{
		Identify:   mobile,
		ExpireTime: expireTime,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return err, tokenString
	}
	return nil, tokenString
}

func StoreToken(request *service.LoginRequest, token string) error {
	fields := make(map[string]interface{})
	fields[request.Mobile] = token
	return dao.HashSet("token", fields)
}
