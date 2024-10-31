package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(id int, secret string, issuer string, expttl time.Duration) string {
	claims := CustomClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		tokenString = ""
	}
	return tokenString
}
