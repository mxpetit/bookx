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
	lastToken, _ := parameters["lastToken"]

	// No token provided
	if lastToken == "" {
		return nil
	}

	_, err := gocql.ParseUUID(lastToken)

	if err != nil {
		return ErrInvalidUUID
	}

	return nil
}
