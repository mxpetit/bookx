package validators

import (
	"errors"
	"github.com/franela/goblin"
	"net/http"
	"testing"
)

func TestUUIDValidator(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Validators > uuid", func() {
		g.Describe("validate", func() {
			g.It("shouldn't return an error since the token is optionnal", func() {
				parameters := map[string]string{}
				validator := New(parameters)
				validator.AddRules(UUID{})
				response := validator.Validate()

				g.Assert(response.Code == http.StatusOK).IsTrue()
				g.Assert(response.Data == nil).IsTrue()
			})

			g.It("should return an error since the token is not valid", func() {
				parameters := map[string]string{
					"uuid": "bad_token",
				}

				validator := New(parameters)
				validator.AddRules(UUID{})
				response := validator.Validate()
				messsages, ok := response.Data["messages"].([]string)

				if !ok {
					err := errors.New("Cannot get response's messages. Failing...")
					g.Fail(err)
				}

				g.Assert(response.Code == http.StatusBadRequest).IsTrue()
				g.Assert(response.Data["length"] == 1).IsTrue()
				g.Assert(len(messsages) == 1).IsTrue()
				g.Assert(messsages[0] == "uuid_invalid").IsTrue()
			})

			g.It("shouldn't return an error since the token is valid", func() {
				parameters := map[string]string{
					// refering to https://tools.ietf.org/html/rfc4122#page-4
					"uuid": "f81d4fae-7dec-11d0-a765-00a0c91e6bf6",
				}

				validator := New(parameters)
				validator.AddRules(UUID{})
				response := validator.Validate()

				g.Assert(response.Code == http.StatusOK).IsTrue()
				g.Assert(response.Data == nil).IsTrue()
			})
		})
	})
}
