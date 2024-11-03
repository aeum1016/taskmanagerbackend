package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id string, expiresAt time.Time) (string, error) {
	// Create the Claims
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	hash := os.Getenv("JWT_PHRASE")
	signingKey := []byte(hash)

	signedToken, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}