package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		log.Printf("Start обработки %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
		log.Printf("Завершено за %v", time.Since(start))
	}
}
