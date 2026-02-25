package middleware

import (
	"errors"
	"kura/internal/auth"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token não informado",
			})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Formato do token inválido",
			})
		}

		tokenString := parts[1]

		token, err := jwt.ParseWithClaims(tokenString, &auth.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			if errors.Is(err, jwt.ErrTokenExpired) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expirado"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			}
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*auth.CustomClaims); ok && token.Valid {
			c.Set("userID", claims.Subject)
		}

		c.Next()
	}
}
