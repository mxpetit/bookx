package validators

import (
	"errors"
	"strconv"
)

type Pagination struct{}

var (
	ErrOffsetInvalid      = errors.New("offset_invalid")
	ErrOffsetIsNotANumber = errors.New("offset_not_number")
)

var (
	acceptedOffsets = []int{10, 25, 50, 100}
)

func (p Pagination) validate(parameters map[string]string) error {
	offset, _ := parameters["offset"]

	// No offset provided
	if offset == "" {
		return nil
	}

	count, err := strconv.Atoi(offset)

	if err != nil {
		return ErrOffsetIsNotANumber
	}

	if !isOffsetAccepted(count) {
		return ErrOffsetInvalid
	}

	return nil
}

func isOffsetAccepted(offset int) bool {
	ok := false

	for i := 0; i < len(acceptedOffsets) && !ok; i++ {
		if acceptedOffsets[i] == offset {
			ok = true
		}
	}

	return ok
}
