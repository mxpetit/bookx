package checkers

import (
	"errors"
)

const (
	SHELVE_NAME_MAXIMUM_SIZE = 64
)

var (
	ErrInvalidShelveName = errors.New("shelve_name_invalid")
	ErrShelveNameMissing = errors.New("shelve_name_missing")
)

func shelveNameCheckFunction(name string, optional bool) error {
	if optional && name == "" {
		return nil
	}

	if name == "" {
		return ErrShelveNameMissing
	}

	if len(name) > SHELVE_NAME_MAXIMUM_SIZE {
		return ErrInvalidShelveName
	}

	return nil
}
