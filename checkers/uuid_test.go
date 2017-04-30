package checkers

import (
	"github.com/franela/goblin"
	"testing"
)

func TestUUIDChecker(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("package checkers > uuid", func() {
		g.Describe("function uuidCheckFunction", func() {
			g.Describe("when parameter is optional", func() {
				g.It("shouldn't return an error when the parameter is empty", func() {
					err := uuidCheckFunction("", true)
					g.Assert(err == nil).IsTrue()
				})

				g.It("should return an error since the token is not valid", func() {
					err := uuidCheckFunction("bad_parameter", true)
					g.Assert(err == ErrInvalidUUID).IsTrue()
				})

				g.It("shouldn't return an error since the parameter is valid", func() {
					// refering to https://tools.ietf.org/html/rfc4122#page-4
					err := uuidCheckFunction("f81d4fae-7dec-11d0-a765-00a0c91e6bf6", true)
					g.Assert(err == nil).IsTrue()
				})
			})

			g.Describe("when parameter is required", func() {
				g.It("should return an error when the parameter is empty", func() {
					err := uuidCheckFunction("", false)
					g.Assert(err == ErrMissingUUID).IsTrue()
				})

				g.It("should return an error since the parameter is not parsable to an int", func() {
					err := uuidCheckFunction("bad_parameter", false)
					g.Assert(err == ErrInvalidUUID).IsTrue()
				})

				g.It("shouldn't return an error since the parameter is valid", func() {
					err := uuidCheckFunction("f81d4fae-7dec-11d0-a765-00a0c91e6bf6", false)
					g.Assert(err == nil).IsTrue()
				})
			})
		})
	})
}
