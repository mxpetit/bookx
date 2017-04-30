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
	CreateShelve(string) (string, error)
	GetShelve(string) (*model.Shelve, error)
}

func GetAllBooks(c context.Context, parameters map[string]string) renderer.Response {
	result, err := FromContext(c).GetAllBooks(parameters["uuid"], parameters["limit"])

	return renderer.RenderGetAllBooks(result, err)
}

func GetBook(c context.Context, parameters map[string]string) renderer.Response {
	result, err := FromContext(c).GetBook(parameters["uuid"])

	return renderer.RenderGetBook(result, err)
}

func CreateBook(c context.Context, title string, numberOfPages int) renderer.Response {
	result, err := FromContext(c).CreateBook(title, numberOfPages)

	return renderer.RenderCreateBook(result, err)
}

func CreateShelve(c context.Context, parameters map[string]string) renderer.Response {
	result, err := FromContext(c).CreateShelve(parameters["name"])

	return renderer.RenderCreateShelve(result, err)
}

func GetShelve(c context.Context, parameters map[string]string) renderer.Response {
	result, err := FromContext(c).GetShelve(parameters["uuid"])

	return renderer.RenderGetShelve(result, err)
}
