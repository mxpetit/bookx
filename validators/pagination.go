package validators

import (
	"errors"
	"strconv"
)

type Pagination struct{}

var (
	ErrLimitInvalid      = errors.New("limit_invalid")
	ErrLimitIsNotANumber = errors.New("limit_not_number")
)

var (
	acceptedLimits = []int{10, 25, 50, 100}
)

func (p Pagination) validate(parameters map[string]string) error {
	limit, _ := parameters["limit"]

	// No limit provided
	if limit == "" {
		return nil
	}

	count, err := strconv.Atoi(limit)

	if err != nil {
		return ErrLimitIsNotANumber
	}

	if !isLimitAccepted(count) {
		return ErrLimitInvalid
	}

	return nil
}

func isLimitAccepted(limit int) bool {
	ok := false

	for i := 0; i < len(acceptedLimits) && !ok; i++ {
		if acceptedLimits[i] == limit {
			ok = true
		}
	}

	return ok
}
