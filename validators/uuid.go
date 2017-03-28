package validators

import (
	"errors"
	"github.com/gocql/gocql"
)

type UUID struct{}

var (
	ErrUnableParseUUID = errors.New("unable_parse_uuid")
	ErrInvalidUUID     = errors.New("uuid_invalid")
)

func (u UUID) validate(parameters *Parameters) error {
	lastToken, ok := (*parameters)["lastToken"]

	// No token provided
	if lastToken == nil {
		return nil
	}

	parsedLastToken, ok := lastToken.(string)

	if !ok {
		return ErrUnableParseUUID
	}

	_, err := gocql.ParseUUID(parsedLastToken)

	if err != nil {
		return ErrInvalidUUID
	}

	return nil
}
