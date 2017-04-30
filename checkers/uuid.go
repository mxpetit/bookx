package checkers

import (
	"errors"
	"github.com/gocql/gocql"
)

var (
	ErrInvalidUUID = errors.New("uuid_invalid")
	ErrMissingUUID = errors.New("uuid_missing")
)

func uuidCheckFunction(uuid string, optional bool) error {
	if optional && uuid == "" {
		return nil
	}

	if uuid == "" {
		return ErrMissingUUID
	}

	_, err := gocql.ParseUUID(uuid)

	if err != nil {
		return ErrInvalidUUID
	}

	return nil
}
