package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ExceptionMiddleware handles panics and returns a 500 Internal Server Error.
func ExceptionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				c.Abort()
			}
		}()
		c.Next()
	}
}
