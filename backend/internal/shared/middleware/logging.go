package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Printf("Request: %s | Status: %d | Latency: %s",
			c.Request.URL.Path,
			c.Writer.Status(),
			time.Since(start))
	}
}
