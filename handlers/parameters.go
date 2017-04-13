package handlers

import (
	"github.com/gin-gonic/gin"
)

// GetParameters returns expectedParameters associated with their values,
// if any.
// Example :
// 		GET /path?foo=bar&baz=qux&quux=corge
//		parameters := GetParameters(context, "foo", "quux", "grault")
//		parameters -> map[foo:bar quux:corge]
func GetParameters(c *gin.Context, expectedParameters ...string) map[string]string {
	parameters := map[string]string{}

	for _, value := range expectedParameters {
		param := c.Query(value)

		if param != "" {
			parameters[value] = param
		}
	}

	return parameters
}
