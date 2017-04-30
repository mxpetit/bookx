package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mxpetit/bookx/checkers"
	"github.com/mxpetit/bookx/model"
	"github.com/mxpetit/bookx/store"
	"net/http"
)

func GetAllBooks(c *gin.Context) {
	parameters := GetParameters(c, "uuid", "limit")

	syntaxCheckerGroup := checkers.NewSyntaxCheckerGroup(parameters)
	syntaxCheckerGroup.AddOptional(checkers.PAGINATION_CHECK_FUNCTION, "limit")
	syntaxCheckerGroup.AddOptional(checkers.UUID_CHECK_FUNCTION, "uuid")
	result := syntaxCheckerGroup.Validate()

	if result.Code == http.StatusOK {
		result = store.GetAllBooks(c, parameters)
	}

	translateAndWriteResponse(c, &result)
}

func GetBook(c *gin.Context) {
	parameters := map[string]string{
		"uuid": c.Param("id"),
	}

	syntaxCheckerGroup := checkers.NewSyntaxCheckerGroup(parameters)
	syntaxCheckerGroup.Add(checkers.UUID_CHECK_FUNCTION, "uuid")
	result := syntaxCheckerGroup.Validate()

	if result.Code == http.StatusOK {
		result = store.GetBook(c, parameters)
	}

	translateAndWriteResponse(c, &result)
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
	translateAndWriteResponse(c, &result)
}
