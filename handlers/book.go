package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/mxpetit/bookx/model"
	"github.com/mxpetit/bookx/store"
	"net/http"
	"strconv"
)

// convertParameter return the first parameter as an int. If it cannot
// be parsed, it returns the default value provided.
func convertParameter(parameter string, onError int) int {
	if parameter == "" {
		return onError
	}

	result, err := strconv.Atoi(parameter)

	if err != nil {
		return onError
	}

	return result
}

func GetAllBooks(c *gin.Context) {
	last := c.Query("last")
	lastId, err := gocql.ParseUUID(last)

	if err != nil {
		lastId = gocql.UUID{}
	}

	count := convertParameter(c.Query("count"), 10)
	result, err := store.GetAllBooks(c, lastId, count)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Unable to retrieves books.",
		})

		return
	}

	c.JSON(http.StatusOK, result)
}

func GetBook(c *gin.Context) {
	id, err := gocql.ParseUUID(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "The UUID provided is not in a valid form.",
		})

		return
	}

	book, err := store.GetBook(c, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "The ressource requested does not exist.",
		})

		return
	}

	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	in := &model.Book{}
	err := c.BindJSON(in)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Unable to bind the book provided to JSON.",
		})

		return
	}

	if in.NumberOfPages < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "The number of pages provided can not be lower than 0.",
		})

		return
	}

	if in.NumberOfPages > 2000 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "The number of pages provided is too large.",
		})

		return
	}

	if in.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "The title provided can not be empty.",
		})

		return
	}

	if len(in.Title) > 256 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "The title provided is too large.",
		})

		return
	}

	id, err := store.CreateBook(c, in.Title, in.NumberOfPages)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Unable to create book.",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "The book was created.",
		"link":    id.String(),
	})
}
