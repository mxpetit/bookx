package validators

import (
	"errors"
	"github.com/gocql/gocql"
)

type UUID struct{}

var (
	ErrInvalidUUID = errors.New("uuid_invalid")
)

func (u UUID) validate(parameters map[string]string) error {
	uuid, _ := parameters["uuid"]

	// No token provided
	if uuid == "" {
		return nil
	}

	_, err := gocql.ParseUUID(uuid)

	if err != nil {
		return ErrInvalidUUID
	}

	return nil
}
