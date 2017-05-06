package datastore

import (
	"github.com/franela/goblin"
	"testing"
)

func TestUtils(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("package datastore > utils", func() {
		g.Describe("function getCqlLimit", func() {
			g.It("should return the limit provided as an int", func() {
				limit := getCqlLimit("10")

				g.Assert(limit == 10).IsTrue()
			})

			g.It("should return DEFAULT_MIN_LIMIT since no limit was provided", func() {
				limit := getCqlLimit("")

				g.Assert(limit == DEFAULT_MIN_LIMIT).IsTrue()
			})

			g.It("should return DEFAULT_MIN_LIMIT since the limit was not parsable to an int", func() {
				limit := getCqlLimit("invalid_limit")

				g.Assert(limit == DEFAULT_MIN_LIMIT).IsTrue()
			})

			g.It("should return DEFAULT_MIN_LIMIT since the limit was lower than DEFAULT_MIN_LIMIT", func() {
				limit := getCqlLimit("0")

				g.Assert(limit == DEFAULT_MIN_LIMIT).IsTrue()
			})

			g.It("should return DEFAULT_MAX_LIMIT since the limit was upper than DEFAULT_MAX_LIMIT", func() {
				limit := getCqlLimit("101")

				g.Assert(limit == DEFAULT_MAX_LIMIT).IsTrue()
			})
		})
	})
}
