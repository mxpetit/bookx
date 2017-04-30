package renderer

import (
	"github.com/mxpetit/bookx/model"
	"net/http"
)

// Response wraps an HTTP response.
type Response struct {
	Code int
	Data map[string]interface{}
}

func RenderGetAllBooks(results []*model.Book, err error) Response {
	datastoreError, _ := err.(model.DatastoreError)

	if !datastoreError.IsNil() {
		return Response{
			Code: datastoreError.Code(),
			Data: map[string]interface{}{
				"message": datastoreError.Error(),
			},
		}
	} else {
		resultsLength := len(results)

		return Response{
			Code: http.StatusOK,
			Data: map[string]interface{}{
				"_links":  results[resultsLength-1].Id,
				"results": results,
				"length":  resultsLength,
			},
		}
	}
}

func RenderGetBook(result *model.Book, err error) Response {
	datastoreError, _ := err.(model.DatastoreError)

	if !datastoreError.IsNil() {
		return Response{
			Code: datastoreError.Code(),
			Data: map[string]interface{}{
				"message": datastoreError.Error(),
			},
		}
	} else {
		return Response{
			Code: http.StatusOK,
			Data: map[string]interface{}{
				"result": result,
			},
		}
	}
}

func RenderCreateBook(result string, err error) Response {
	datastoreError, _ := err.(model.DatastoreError)

	if !datastoreError.IsNil() {
		return Response{
			Code: datastoreError.Code(),
			Data: map[string]interface{}{
				"message": datastoreError.Error(),
			},
		}
	} else {
		return Response{
			Code: http.StatusOK,
			Data: map[string]interface{}{
				"_links":  result,
				"message": "book_created",
			},
		}
	}
}

func RenderCreateShelve(result string, err error) Response {
	datastoreError, _ := err.(model.DatastoreError)

	if !datastoreError.IsNil() {
		return Response{
			Code: datastoreError.Code(),
			Data: map[string]interface{}{
				"message": datastoreError.Error(),
			},
		}
	} else {
		return Response{
			Code: http.StatusOK,
			Data: map[string]interface{}{
				"_links":  result,
				"message": "shelve_created",
			},
		}
	}
}

func RenderGetShelve(result *model.Shelve, err error) Response {
	datastoreError, _ := err.(model.DatastoreError)

	if !datastoreError.IsNil() {
		return Response{
			Code: datastoreError.Code(),
			Data: map[string]interface{}{
				"message": datastoreError.Error(),
			},
		}
	} else {
		return Response{
			Code: http.StatusOK,
			Data: map[string]interface{}{
				"result": result,
			},
		}
	}
}
