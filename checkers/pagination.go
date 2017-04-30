package checkers

import (
	"errors"
	"strconv"
)

var (
	ErrLimitIsNotANumber = errors.New("limit_not_number")
	ErrLimitIsMissing    = errors.New("limit_missing")
)

func paginationCheckFunction(pagination string, optional bool) error {
	if optional && pagination == "" {
		return nil
	}

	if pagination == "" {
		return ErrLimitIsMissing
	}

	_, err := strconv.Atoi(pagination)

	if err != nil {
		return ErrLimitIsNotANumber
	}

	return nil
}
