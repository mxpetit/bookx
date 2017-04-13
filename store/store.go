package store

import (
	"github.com/mxpetit/bookx/model"
	"github.com/mxpetit/bookx/renderer"
	"golang.org/x/net/context"
)

type Store interface {
	GetAllBooks(string, string) ([]*model.Book, error)
	GetBook(string) (*model.Book, error)
	CreateBook(string, int) (string, error)
}

func GetAllBooks(c context.Context, parameters map[string]string) renderer.Response {
	result, err := FromContext(c).GetAllBooks(parameters["last_token"], parameters["offset"])

	return renderer.RenderGetAllBooks(result, err)
}

func GetBook(c context.Context, id string) renderer.Response {
	result, err := FromContext(c).GetBook(id)

	return renderer.RenderGetBook(result, err)
}

func CreateBook(c context.Context, title string, numberOfPages int) renderer.Response {
	result, err := FromContext(c).CreateBook(title, numberOfPages)

	return renderer.RenderCreateBook(result, err)
}
