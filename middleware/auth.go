package middleware

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		log.Printf("%s %s %s", c.Request.Method, c.Request.RequestURI, latency)
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token != os.Getenv("PRIVATE_TOKEN") {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
