package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenClaims struct {
	Mobile     string
	ExpireTime time.Time
	jwt.StandardClaims
}
