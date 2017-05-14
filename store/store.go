package store

import (
	"github.com/mxpetit/bookx/model"
	"github.com/mxpetit/bookx/renderer"
	"golang.org/x/net/context"
)

type Store interface {
	CreateBook(string, int) (string, error)
	CreateShelve(string) (string, error)
	GetAllBooks(string) ([]*model.Book, error)
	GetBook(string) (*model.Book, error)
	GetNextBooks(string, string) ([]*model.Book, error)
	GetPreviousBooks(string, string) ([]*model.Book, error)
	GetShelve(string) (*model.Shelve, error)
}

func CreateBook(c context.Context, title string, numberOfPages int) *renderer.Response {
	result, err := FromContext(c).CreateBook(title, numberOfPages)

	return renderer.RenderCreateBook(result, err)
}

func CreateShelve(c context.Context, parameters map[string]string) *renderer.Response {
	result, err := FromContext(c).CreateShelve(parameters["name"])

	return renderer.RenderCreateShelve(result, err)
}

func GetAllBooks(c context.Context, parameters map[string]string) *renderer.Response {
	result, err := FromContext(c).GetAllBooks(parameters["limit"])

	return renderer.RenderGetAllBooks(result, err)
}

func GetBook(c context.Context, parameters map[string]string) *renderer.Response {
	result, err := FromContext(c).GetBook(parameters["id"])

	return renderer.RenderGetBook(result, err)
}

func GetNextBooks(c context.Context, parameters map[string]string) *renderer.Response {
	result, err := FromContext(c).GetNextBooks(parameters["id"], parameters["limit"])

	return renderer.RenderGetAllBooks(result, err)
}

func GetPreviousBooks(c context.Context, parameters map[string]string) *renderer.Response {
	result, err := FromContext(c).GetPreviousBooks(parameters["id"], parameters["limit"])

	return renderer.RenderGetAllBooks(result, err)
}

func GetShelve(c context.Context, parameters map[string]string) *renderer.Response {
	result, err := FromContext(c).GetShelve(parameters["uuid"])

	return renderer.RenderGetShelve(result, err)
}
