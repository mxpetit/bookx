package checkers

import (
	"github.com/franela/goblin"
	"testing"
)

func TestPaginationChecker(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("package checkers > pagination", func() {
		g.Describe("function paginationCheckFunction", func() {
			g.Describe("when parameter is optional", func() {
				g.It("shouldn't return an error when the parameter is empty", func() {
					err := paginationCheckFunction("", true)
					g.Assert(err == nil).IsTrue()
				})

				g.It("should return an error since the parameter is not parsable to an int", func() {
					err := paginationCheckFunction("bad_parameter", true)
					g.Assert(err == ErrLimitIsNotANumber).IsTrue()
				})

				g.It("shouldn't return an error since the parameter is valid", func() {
					err := paginationCheckFunction("10", true)
					g.Assert(err == nil).IsTrue()
				})
			})

			g.Describe("when parameter is required", func() {
				g.It("should return an error when the parameter is empty", func() {
					err := paginationCheckFunction("", false)
					g.Assert(err == ErrLimitIsMissing).IsTrue()
				})

				g.It("should return an error since the parameter is not parsable to an int", func() {
					err := paginationCheckFunction("bad_parameter", false)
					g.Assert(err == ErrLimitIsNotANumber).IsTrue()
				})

				g.It("shouldn't return an error since the parameter is valid", func() {
					err := paginationCheckFunction("10", false)
					g.Assert(err == nil).IsTrue()
				})
			})
		})
	})
}
