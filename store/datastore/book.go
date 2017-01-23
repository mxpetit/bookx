package datastore

import (
	"errors"
	"github.com/gocql/gocql"
	"github.com/mxpetit/bookx/model"
	"github.com/mxpetit/bookx/model/response"
)

const (
	getAllPagedBooks = "SELECT * FROM book WHERE token(id) > token(?) LIMIT ?"
	getAllBooks      = "SELECT * FROM book LIMIT ?"
	getBook          = "SELECT * FROM book WHERE id = ?"
	insertBook       = "INSERT INTO book (id, number_of_pages, title) VALUES (?, ?, ?)"
)

// GetAllBooks retuns the number of books from the UUID to count. If the UUID
// is empty, it'll get the first books as specified by count parameter.
func (db *datastore) GetAllBooks(oldId gocql.UUID, count int) (response.Multiple, error) {
	if count < 1 {
		return response.Multiple{}, errors.New("Cannot get less than 1 item.")
	}

	var title string
	var numberOfPages int
	var bookId, newId gocql.UUID
	var queryResults []interface{}
	var query *gocql.Query

	// If no id provided, we get the first elements
	if oldId == (gocql.UUID{}) {
		query = db.Query(getAllBooks, count)
	} else {
		query = db.Query(getAllPagedBooks, oldId, count)
	}

	iter := query.Iter()

	for iter.Scan(&bookId, &numberOfPages, &title) {
		queryResults = append(queryResults, model.Book{
			Id:            bookId,
			Title:         title,
			NumberOfPages: numberOfPages,
		})

		newId = bookId
	}

	if err := iter.Close(); err != nil {
		return response.Multiple{}, err
	}

	multiple := response.Multiple{
		Links:   newId.String(),
		Results: queryResults,
		Length:  len(queryResults),
	}

	return multiple, nil
}

// GetBook returns the details of a book given its id.
func (db *datastore) GetBook(id gocql.UUID) (model.Book, error) {
	var title string
	var numberOfPages int

	if err := db.Query(getBook, id).Scan(&id, &numberOfPages, &title); err != nil {
		return model.Book{}, err
	}

	queryResult := model.Book{
		Id:            id,
		Title:         title,
		NumberOfPages: numberOfPages,
	}

	return queryResult, nil
}

// CreateBook returns the id of the book that was created.
func (db *datastore) CreateBook(title string, numberOfPages int) (gocql.UUID, error) {
	id, err := gocql.RandomUUID()

	if err != nil {
		return gocql.UUID{}, err
	}

	if err = db.Query(insertBook, id, numberOfPages, title).Exec(); err != nil {
		return gocql.UUID{}, err
	}

	return id, nil
}
