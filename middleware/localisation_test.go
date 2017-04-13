package middleware

import (
	"github.com/franela/goblin"
	"testing"
)

func TestLocalisation(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("isLanguageAccepted", func() {
		g.It("should check if a given language is accepted", func() {
			g.Assert(isLanguageAccepted("fr-FR")).IsTrue()
			g.Assert(isLanguageAccepted("en-US")).IsTrue()
			g.Assert(isLanguageAccepted("language_that_doesnt_exists")).IsFalse()
		})
	})
}
