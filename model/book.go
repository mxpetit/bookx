package model

import (
	"github.com/gocql/gocql"
)

type Book struct {
	Id            gocql.UUID `json:"id"`
	Title         string     `json:"title"`
	NumberOfPages int        `json:"number_of_pages"`
}
