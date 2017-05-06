package datastore

import (
	"strconv"
)

// getCqlLimit returns the limit of a cql statement. It will be adjusted
// if it exceeds DEFAULT_MIN_LIMIT or DEFAULT_MAX_LIMIT.
func getCqlLimit(limit string) int {
	if limit == "" {
		return DEFAULT_MIN_LIMIT
	}

	parsedLimit, err := strconv.Atoi(limit)

	if err != nil || parsedLimit < DEFAULT_MIN_LIMIT {
		return DEFAULT_MIN_LIMIT
	}

	if parsedLimit > DEFAULT_MAX_LIMIT {
		return DEFAULT_MAX_LIMIT
	}

	return parsedLimit
}
