package util

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func GenerateJwt(issuer string) (string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.RegisteredClaims{
		Issuer: issuer,
		ExpiresAt: &jwt.NumericDate{
			Time: time.Now().Add(time.Hour * 12),
		},
	})
	secretKey := []byte(os.Getenv("SECRETKEY"))
	return claims.SignedString(secretKey)
}
