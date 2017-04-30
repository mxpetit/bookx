package checkers

import (
	"github.com/franela/goblin"
	"testing"
)

const (
	// len(bad_shelve_name) should be superior to SHELVE_NAME_MAXIMUM_SIZE
	// constant.
	bad_shelve_name = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
)

func TestShelveNameChecker(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("package checkers > shelve_name", func() {
		g.Describe("function shelveNameCheckFunction", func() {
			g.Describe("when parameter is optional", func() {
				g.It("shouldn't return an error when shelve's name is empty", func() {
					err := shelveNameCheckFunction("", true)
					g.Assert(err == nil).IsTrue()
				})

				g.It("should return an error when shelve's name is not valid", func() {
					err := shelveNameCheckFunction(bad_shelve_name, true)
					g.Assert(err == ErrInvalidShelveName).IsTrue()
				})

				g.It("shouldn't return an error when shelve's name is valid", func() {
					err := shelveNameCheckFunction("shelve_name", true)
					g.Assert(err == nil).IsTrue()
				})
			})

			g.Describe("when parameter is required", func() {
				g.It("should return an error when shelve's name is empty", func() {
					err := shelveNameCheckFunction("", false)
					g.Assert(err == ErrShelveNameMissing).IsTrue()
				})

				g.It("should return an error when shelve's name is not valid", func() {
					err := shelveNameCheckFunction(bad_shelve_name, false)
					g.Assert(err == ErrInvalidShelveName).IsTrue()
				})

				g.It("shouldn't return an error when shelve's name is valid", func() {
					err := shelveNameCheckFunction("shelve_name", false)
					g.Assert(err == nil).IsTrue()
				})
			})
		})
	})
}
