package middleware

import (
	"github.com/gin-gonic/gin"
)

const (
	DEFAULT_LANGUAGE = "fr-FR"
)

// Localisation ensure that the requested language is available. Otherwise
// it'll use the default language as specified by the const DEFAULT_LANGUAGE.
func Localisation() gin.HandlerFunc {
	return func(c *gin.Context) {
		language := c.Request.Header.Get("Accept-Language")

		if !isLanguageAccepted(language) {
			c.Request.Header.Set("Accept-Language", DEFAULT_LANGUAGE)
		}

		c.Next()
	}
}

// isLanguageValid checks if the given language is valid and available.
func isLanguageAccepted(language string) bool {
	switch language {
	case DEFAULT_LANGUAGE, "en-US":
		return true
	default:
		return false
	}
}
