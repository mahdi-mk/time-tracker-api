package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mahdi-mk/time-tracker/support/env"
)

// TokenPayload represents the payload (data) part of the JWT
type TokenPayload struct {
	UserID uint
}

// GenerateToken generates a valid JWT using the given payload
func GenerateToken(payload *TokenPayload) (string, error) {
	claims := jwt.MapClaims{
		"uid": payload.UserID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	encodedToken, err := token.SignedString([]byte(env.Get("JWT_SECRET", "")))
	if err != nil {
		return "", err
	}

	return encodedToken, nil
}
