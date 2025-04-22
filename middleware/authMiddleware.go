package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rest-api/config/auth"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No Authorization"})
			c.Abort()
			return
		}
		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		token, err := auth.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		log.Printf("Succes.claims: %+v\\\\n", token.Claims)
		c.Next()
	}
}
