package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mxpetit/bookx/renderer"
)

// getLanguage returns the language associated with the request.
func getLanguage(c *gin.Context) string {
	return c.Request.Header.Get("Accept-Language")
}

// translateAndWriteResponse wraps the response's sending,
// for translation purpose.
func translateAndWriteResponse(c *gin.Context, response *renderer.Response) {
	language := getLanguage(c)
	response.Translate(language)
	c.JSON(response.Code, response.Data)
}
