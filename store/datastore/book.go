package datastore

import (
	"github.com/gocql/gocql"
	"github.com/mxpetit/bookx/model"
	"net/http"
)

const (
	getNextBooks     = "SELECT * FROM book WHERE token(id) >= token(?) LIMIT ?"
	getPreviousBooks = "SELECT * FROM book WHERE token(id) <= token(?) LIMIT ?"
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

	DEFAULT_MIN_LIMIT = 1
	DEFAULT_MAX_LIMIT = 100
)

// scanBooks scans an iterator to retrieve a list of books.
func scanBooks(iter *gocql.Iter) ([]*model.Book, error) {
	if iter == nil {
		return []*model.Book{}, ErrUnableToRetrievesBooks
	}

	var title string
	var numberOfPages int
	var id gocql.UUID
	var results []*model.Book

	for iter.Scan(&id, &numberOfPages, &title) {
		results = append(results, &model.Book{
			Id:            id,
			Title:         title,
			NumberOfPages: numberOfPages,
		})
	}

	if err := iter.Close(); err != nil {
		return []*model.Book{}, ErrUnableToRetrievesBooks
	}

	if len(results) == 0 {
		return []*model.Book{}, ErrNoBooksAvailable
	}

	return results, nil
}

// GetAllBooks gets n first books.
func (db *datastore) GetAllBooks(limit string) ([]*model.Book, error) {
	parsedLimit := getCqlLimit(limit)
	iter := db.Query(getAllBooks, parsedLimit).Iter()

	return scanBooks(iter)
}

// GetBook gets book's details given its uuid.
func (db *datastore) GetBook(uuid string) (*model.Book, error) {
	var title string
	var numberOfPages int

	parsedId, err := gocql.ParseUUID(uuid)

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

// GetNextBooks gets the n next books to id.
func (db *datastore) GetNextBooks(id string, limit string) ([]*model.Book, error) {
	parsedLimit := getCqlLimit(limit)
	parsedId, err := gocql.ParseUUID(id)

	if err != nil {
		return []*model.Book{}, ErrUUIDInvalid
	}

	iter := db.Query(getNextBooks, parsedId, parsedLimit).Iter()

	return scanBooks(iter)
}

// GetPreviousBooks gets the n previous books to id.
func (db *datastore) GetPreviousBooks(id string, limit string) ([]*model.Book, error) {
	parsedLimit := getCqlLimit(limit)
	parsedId, err := gocql.ParseUUID(id)

	if err != nil {
		return []*model.Book{}, ErrUUIDInvalid
	}

	iter := db.Query(getPreviousBooks, parsedId, parsedLimit).Iter()

	return scanBooks(iter)
}

// CreateBook creates a new book with generated UUID.
func (db *datastore) CreateBook(title string, numberOfPages int) (string, error) {
	uuid, err := gocql.RandomUUID()

	if err != nil {
		return gocql.UUID{}.String(), ErrUnableToGenerateUUID
	}

	if err = db.Query(insertBook, uuid, numberOfPages, title).Exec(); err != nil {
		return gocql.UUID{}.String(), ErrUnableToCreateBook
	}

	return uuid.String(), nil
}
