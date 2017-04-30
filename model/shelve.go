package model

import (
	"github.com/gocql/gocql"
)

type Shelve struct {
	Id   gocql.UUID `json:"id"`
	Name string     `json:"name"`
}
