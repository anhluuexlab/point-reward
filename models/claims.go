package models

import "github.com/dgrijalva/jwt-go"

type JwtCustomClaims struct {
	jwt.StandardClaims
}
