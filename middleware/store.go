package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mxpetit/bookx/store"
	"github.com/mxpetit/bookx/store/datastore"
)

// Store is a middleware that initializes the Datastore and attaches to
// the context of every http.Request.
func Store() gin.HandlerFunc {
	v := datastore.New()

	return func(c *gin.Context) {
		store.ToContext(c, v)
		c.Next()
	}
}
