package validators

import (
	"errors"
	"strconv"
)

type Pagination struct{}

var (
	ErrOffsetInvalid      = errors.New("offset_invalid")
	ErrUnableParseOffset  = errors.New("unable_parse_offset")
	ErrOffsetIsNotANumber = errors.New("offset_not_number")
)

var (
	acceptedOffsets = []int{10, 25, 50, 100}
)

func (p Pagination) validate(parameters *Parameters) error {
	offset, ok := (*parameters)["offset"]

	// No offset provided
	if offset == "" {
		return nil
	}

	offsetString, ok := offset.(string)

	if !ok {
		return ErrUnableParseOffset
	}

	count, err := strconv.Atoi(offsetString)

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
