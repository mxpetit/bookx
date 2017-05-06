package handlers

import (
	"github.com/gin-gonic/gin"
)

// GetParameters gets parameters associated with their values,
// from query and param. If the parameter is both in the query and param,
// the value associated with param is taken. Missing values will not be in
// the result.
func GetParameters(c *gin.Context, expectedParameters ...string) map[string]string {
	parameters := map[string]string{}

	for _, key := range expectedParameters {
		param := c.Param(key)

		if param == "" {
			param = c.Query(key)
		}

		if param != "" {
			parameters[key] = param
		}
	}

	return parameters
}
