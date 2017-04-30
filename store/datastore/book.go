package datastore

import (
	"github.com/gocql/gocql"
	"github.com/mxpetit/bookx/model"
	"net/http"
	"strconv"
)

const (
	getAllPagedBooks = "SELECT * FROM book WHERE token(id) >= token(?) LIMIT ?"
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

	DEFAULT_MIN_LIMIT = 10
	DEFAULT_MAX_LIMIT = 100
)

// GetAllBooks returns the number of books between uuid and limit. If uuid isn't valid,
// it'll get the first books as specified by limit parameter.
func (db *datastore) GetAllBooks(uuid, limit string) ([]*model.Book, error) {
	var title string
	var numberOfPages int
	var id gocql.UUID
	var results []*model.Book
	var query *gocql.Query

	parsedLimit, err := strconv.Atoi(limit)

	if err != nil {
		parsedLimit = DEFAULT_MIN_LIMIT
	}

	if parsedLimit > DEFAULT_MAX_LIMIT {
		parsedLimit = DEFAULT_MAX_LIMIT
	}

	if parsedLimit < DEFAULT_MIN_LIMIT {
		parsedLimit = DEFAULT_MIN_LIMIT
	}

	if parsedId, err := gocql.ParseUUID(uuid); err != nil {
		query = db.Query(getAllBooks, parsedLimit)
	} else {
		query = db.Query(getAllPagedBooks, parsedId, parsedLimit)
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

// GetBook returns book's details given its uuid.
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

// CreateBook returns the book's uuid that was created.
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
