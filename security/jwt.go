package security

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/zett-8/go-clean-echo/models"
)

const JWT_KEY = "hhhgfdshgfhsdgfshjgfshjdgf"

func GenToken() (string, error) {
	claims := &models.JwtCustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWT_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
