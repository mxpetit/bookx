package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mxpetit/bookx/checkers"
	"github.com/mxpetit/bookx/model"
	"github.com/mxpetit/bookx/store"
	"net/http"
)

func GetAllBooks(c *gin.Context) {
	parameters := GetParameters(c, "limit")

	syntaxCheckerGroup := checkers.NewSyntaxCheckerGroup(parameters)
	syntaxCheckerGroup.AddOptional(checkers.PAGINATION_CHECK_FUNCTION, "limit")
	result := syntaxCheckerGroup.Validate()

	if result.Code == http.StatusOK {
		result = store.GetAllBooks(c, parameters)
	}

	translateAndWriteResponse(c, result)
}

func GetNextBooks(c *gin.Context) {
	parameters := GetParameters(c, "id", "limit")

	syntaxCheckerGroup := checkers.NewSyntaxCheckerGroup(parameters)
	syntaxCheckerGroup.Add(checkers.UUID_CHECK_FUNCTION, "id")
	syntaxCheckerGroup.AddOptional(checkers.PAGINATION_CHECK_FUNCTION, "limit")
	result := syntaxCheckerGroup.Validate()

	if result.Code == http.StatusOK {
		result = store.GetNextBooks(c, parameters)
	}

	translateAndWriteResponse(c, result)
}

func GetPreviousBooks(c *gin.Context) {
	parameters := GetParameters(c, "id", "limit")

	syntaxCheckerGroup := checkers.NewSyntaxCheckerGroup(parameters)
	syntaxCheckerGroup.Add(checkers.UUID_CHECK_FUNCTION, "id")
	syntaxCheckerGroup.AddOptional(checkers.PAGINATION_CHECK_FUNCTION, "limit")
	result := syntaxCheckerGroup.Validate()

	if result.Code == http.StatusOK {
		result = store.GetPreviousBooks(c, parameters)
	}

	translateAndWriteResponse(c, result)
}

func GetBook(c *gin.Context) {
	parameters := GetParameters(c, "id")

	syntaxCheckerGroup := checkers.NewSyntaxCheckerGroup(parameters)
	syntaxCheckerGroup.Add(checkers.UUID_CHECK_FUNCTION, "id")
	result := syntaxCheckerGroup.Validate()

	if result.Code == http.StatusOK {
		result = store.GetBook(c, parameters)
	}

	translateAndWriteResponse(c, result)
}

func CreateBook(c *gin.Context) {
	in := &model.Book{}
	err := c.BindJSON(in)

	if err != nil {
		c.String(http.StatusBadRequest, "json_invalid")

		return
	}

	err = in.Validate()

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())

		return
	}

	result := store.CreateBook(c, in.Title, in.NumberOfPages)
	translateAndWriteResponse(c, result)
}
