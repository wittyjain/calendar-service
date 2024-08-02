package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		fmt.Printf("Method: %s, Path: %s, Status: %d, Duration: %v\n",
			c.Request.Method,
			c.Request.RequestURI,
			c.Writer.Status(),
			duration,
		)
	}
}
