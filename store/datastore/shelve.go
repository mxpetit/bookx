package datastore

import (
	"github.com/gocql/gocql"
	"github.com/mxpetit/bookx/model"
	"net/http"
)

const (
	createShelve    = "INSERT INTO shelve (id, name) VALUES (?, ?)"
	getShelve       = "SELECT * FROM shelve WHERE id = ?"
	addBookToShelve = "INSERT INTO book_by_shelve (shelve_id, book_id, book_name) VALUES (?, ?, ?)"
)

var (
	ErrUnableToCreateShelve = model.NewDatastoreError(http.StatusInternalServerError, "unable_create_shelve")
	ErrShelveDoesNotExists  = model.NewDatastoreError(http.StatusNotFound, "shelve_doesnt_exists")
	ErrUnableToAddBook      = model.NewDatastoreError(http.StatusInternalServerError, "unable_add_book_shelve")
	ErrNotImplemented       = model.NewDatastoreError(http.StatusNotImplemented, "method_not_implemented")
)

// CreateShelve creates an empty shelve.
func (db *datastore) CreateShelve(name string) (string, error) {
	uuid, err := gocql.RandomUUID()

	if err != nil {
		return gocql.UUID{}.String(), ErrUnableToGenerateUUID
	}

	if err = db.Query(createShelve, uuid, name).Exec(); err != nil {
		return gocql.UUID{}.String(), ErrUnableToCreateShelve
	}

	return uuid.String(), nil
}

// GetShelve retrieves shelve's details given an uuid.
func (db *datastore) GetShelve(uuid string) (*model.Shelve, error) {
	var name string

	parsedId, err := gocql.ParseUUID(uuid)

	if err != nil {
		return &model.Shelve{}, ErrUUIDInvalid
	}

	err = db.Query(getShelve, parsedId).Scan(&parsedId, &name)

	if err == gocql.ErrNotFound {
		return &model.Shelve{}, ErrShelveDoesNotExists
	}

	queryResult := &model.Shelve{
		Id:   parsedId,
		Name: name,
	}

	return queryResult, nil
}
