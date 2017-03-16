package datastore

import (
	"github.com/gocql/gocql"
	"github.com/mxpetit/bookx/model"
	"net/http"
	"strconv"
)

const (
	getAllPagedBooks = "SELECT * FROM book WHERE token(id) > token(?) LIMIT ?"
	getAllBooks      = "SELECT * FROM book LIMIT ?"
	getBook          = "SELECT * FROM book WHERE id = ?"
	insertBook       = "INSERT INTO book (id, number_of_pages, title) VALUES (?, ?, ?)"
)

var (
	ErrNoBooksAvailable       = model.NewDatastoreError(http.StatusNotFound, "no_books_available")
	ErrUnableToRetrievesBooks = model.NewDatastoreError(http.StatusInternalServerError, "unable_retrieves_books")
	ErrBookDoesNotExists      = model.NewDatastoreError(http.StatusNotFound, "book_doesnt_exists")
	ErrUnableToGenerateUUID   = model.NewDatastoreError(http.StatusInternalServerError, "unable_generate_uuid")
	ErrUnableToCreateBook     = model.NewDatastoreError(http.StatusInternalServerError, "unable_create_book")
	ErrUUIDInvalid            = model.NewDatastoreError(http.StatusBadRequest, "uuid_invalid")
)

// GetAllBooks retuns the number of books between lastToken and offset. If lastToken isn't valid,
// it'll get the first books as specified by offset parameter.
func (db *datastore) GetAllBooks(lastToken, offset string) ([]*model.Book, error) {
	var title string
	var numberOfPages int
	var id gocql.UUID
	var results []*model.Book
	var query *gocql.Query

	parsedOffset, err := strconv.Atoi(offset)

	if err != nil {
		parsedOffset = 10
	}

	if parsedId, err := gocql.ParseUUID(lastToken); err != nil {
		query = db.Query(getAllBooks, parsedOffset)
	} else {
		query = db.Query(getAllPagedBooks, parsedId, parsedOffset)
	}

	iter := query.Iter()

	for iter.Scan(&id, &numberOfPages, &title) {
		results = append(results, &model.Book{
			Id:            id,
			Title:         title,
			NumberOfPages: numberOfPages,
		})
	}

	if err = iter.Close(); err != nil {
		return []*model.Book{}, ErrUnableToRetrievesBooks
	}

	if len(results) == 0 {
		return []*model.Book{}, ErrNoBooksAvailable
	}

	return results, nil
}

// GetBook returns book's details given its id.
func (db *datastore) GetBook(id string) (*model.Book, error) {
	var title string
	var numberOfPages int

	parsedId, err := gocql.ParseUUID(id)

	if err != nil {
		return &model.Book{}, ErrUUIDInvalid
	}

	err = db.Query(getBook, parsedId).Scan(&parsedId, &numberOfPages, &title)

	if err == gocql.ErrNotFound {
		return &model.Book{}, ErrBookDoesNotExists
	}

	queryResult := &model.Book{
		Id:            parsedId,
		Title:         title,
		NumberOfPages: numberOfPages,
	}

	return queryResult, nil
}

// CreateBook returns the book's id that was created.
func (db *datastore) CreateBook(title string, numberOfPages int) (string, error) {
	id, err := gocql.RandomUUID()

	if err != nil {
		return gocql.UUID{}.String(), ErrUnableToGenerateUUID
	}

	if err = db.Query(insertBook, id, numberOfPages, title).Exec(); err != nil {
		return gocql.UUID{}.String(), ErrUnableToCreateBook
	}

	return id.String(), nil
}
