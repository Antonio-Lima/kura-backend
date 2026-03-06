package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type CustomClaims struct {
	jwt.RegisteredClaims
}

func GenerateToken(userID uuid.UUID) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")

	if jwtSecret == "" {
		return "", fmt.Errorf("JWT_SECRET não configurado")
	}

	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: userID.String(),
			// ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			// Voltar para 24 horas, ao subir para produção.
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(1, 0, 0)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtSecret))
}
