package model

import (
	"errors"
	"github.com/gocql/gocql"
)

type Book struct {
	Id            gocql.UUID `json:"id"`
	Title         string     `json:"title"`
	NumberOfPages int        `json:"number_of_pages"`
}

func (book *Book) Validate() error {
	if err := book.validateNumberOfPages(); err != "" {
		return errors.New(err)
	}

	if err := book.validateTitle(); err != "" {
		return errors.New(err)
	}

	return nil
}

func (book *Book) validateNumberOfPages() string {
	if book.NumberOfPages < 1 {
		return "number_of_pages_under_1"
	}

	if book.NumberOfPages > 5000 {
		return "number_of_pages_upper_5000"
	}

	return ""
}

func (book *Book) validateTitle() string {
	if book.Title == "" {
		return "title_empty"
	}

	if len(book.Title) > 256 {
		return "title_upper_256"
	}

	return ""
}
