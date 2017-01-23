package store

import (
	"golang.org/x/net/context"
)

const key = "store"

type Setter interface {
	Set(string, interface{})
}

// FromContext returns the Store associated with the context.
func FromContext(c context.Context) Store {
	return c.Value(key).(Store)
}

// ToContext associates the Store with the context.
func ToContext(s Setter, store Store) {
	s.Set(key, store)
}
