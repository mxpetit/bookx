package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mxpetit/bookx/checkers"
	"github.com/mxpetit/bookx/store"
	"net/http"
)

func CreateShelve(c *gin.Context) {
	parameters := GetParameters(c, "name")

	syntaxCheckerGroup := checkers.NewSyntaxCheckerGroup(parameters)
	syntaxCheckerGroup.Add(checkers.SHELVE_NAME_CHECK_FUNCTION, "name")
	result := syntaxCheckerGroup.Validate()

	if result.Code == http.StatusOK {
		result = store.CreateShelve(c, parameters)
	}

	translateAndWriteResponse(c, result)
}

func GetShelve(c *gin.Context) {
	parameters := GetParameters(c, "id")

	syntaxCheckerGroup := checkers.NewSyntaxCheckerGroup(parameters)
	syntaxCheckerGroup.Add(checkers.UUID_CHECK_FUNCTION, "id")
	result := syntaxCheckerGroup.Validate()

	if result.Code == http.StatusOK {
		result = store.GetShelve(c, parameters)
	}

	translateAndWriteResponse(c, result)
}
