package store

import (
	"github.com/gocql/gocql"
	"github.com/mxpetit/bookx/model"
	"github.com/mxpetit/bookx/model/response"
	"golang.org/x/net/context"
)

type Store interface {
	GetAllBooks(gocql.UUID, int) (response.Multiple, error)
	GetBook(id gocql.UUID) (model.Book, error)
	CreateBook(title string, numberOfPages int) (gocql.UUID, error)
}

func GetAllBooks(c context.Context, oldId gocql.UUID, count int) (response.Multiple, error) {
	return FromContext(c).GetAllBooks(oldId, count)
}

func GetBook(c context.Context, id gocql.UUID) (model.Book, error) {
	return FromContext(c).GetBook(id)
}

func CreateBook(c context.Context, title string, numberOfPages int) (gocql.UUID, error) {
	return FromContext(c).CreateBook(title, numberOfPages)
}
