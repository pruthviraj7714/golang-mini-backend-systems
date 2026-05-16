package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"url-shortener/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func loadJWTSecret() string {
	return config.LoadConfig().JWTSecret
}

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header is missing",
			})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		if len(token) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header is malformed",
			})
			c.Abort()
			return
		}

		parsedToken, err := jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(loadJWTSecret()), nil
		})

		if err != nil || !parsedToken.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header is invalid",
			})
			c.Abort()
			return

		}

		claims, ok := parsedToken.Claims.(*jwt.MapClaims)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header is invalid",
			})
			c.Abort()
			return
		}

		c.Set("userId", (*claims)["userId"])

		c.Next()
	}

}
