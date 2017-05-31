package middleware

import (
	"github.com/gin-gonic/gin"
)

// Cors allows cross-origin requests to be proccessed.
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://ui.book.xyz")
		c.Next()
	}
}
