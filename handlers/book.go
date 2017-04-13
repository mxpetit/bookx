package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mxpetit/bookx/model"
	"github.com/mxpetit/bookx/store"
	"github.com/mxpetit/bookx/validators"
	"net/http"
)

func GetAllBooks(c *gin.Context) {
	parameters := GetParameters(c, "last_token", "offset")
	validator := validators.New(parameters)
	validator.AddRules(validators.Pagination{}, validators.UUID{})
	result := validator.Validate()

	if result.Code == http.StatusOK {
		result = store.GetAllBooks(c, parameters)
	}

	translateAndWriteResponse(c, &result)
}

func GetBook(c *gin.Context) {
	result := store.GetBook(c, c.Param("id"))
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
