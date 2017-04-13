package validators

import (
	"errors"
	"github.com/franela/goblin"
	"net/http"
	"reflect"
	"testing"
)

func TestValidator(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Validators > validator", func() {
		g.Describe("appendError", func() {
			g.It("should append an error", func() {
				var strings []string

				wanted := []string{
					"validator_error",
				}

				err := errors.New("validator_error")
				strings = appendError(strings, err)

				g.Assert(reflect.DeepEqual(strings, wanted)).IsTrue()
				g.Assert(len(strings) == 1).IsTrue()
			})

			g.It("shouldn't append an error since the error is nil", func() {
				var strings []string

				strings = appendError(strings, nil)

				g.Assert(len(strings) == 0).IsTrue()
			})
		})

		g.Describe("AddRules", func() {
			g.It("should add one rule", func() {
				parameters := map[string]string{}
				validator := New(parameters)

				validator.AddRules(Pagination{})

				g.Assert(len(validator.rules) == 1).IsTrue()
				g.Assert(validator.rules[0] == Pagination{}).IsTrue()
			})

			g.It("should add two rules", func() {
				parameters := map[string]string{}
				validator := New(parameters)

				validator.AddRules(Pagination{}, UUID{})

				g.Assert(len(validator.rules) == 2).IsTrue()
				g.Assert(validator.rules[0] == Pagination{}).IsTrue()
				g.Assert(validator.rules[1] == UUID{}).IsTrue()
			})

			g.It("should add rules without duplicates", func() {
				parameters := map[string]string{}
				validator := New(parameters)

				validator.AddRules(Pagination{}, UUID{}, Pagination{}, UUID{})

				g.Assert(len(validator.rules) == 2).IsTrue()
				g.Assert(validator.rules[0] == Pagination{}).IsTrue()
				g.Assert(validator.rules[1] == UUID{}).IsTrue()
			})
		})

		g.Describe("containsRule", func() {
			g.It("should check if the given rules is registered", func() {
				parameters := map[string]string{}
				validator := New(parameters)
				validator.AddRules(Pagination{})

				contains1 := validator.containsRule(Pagination{})
				contains2 := validator.containsRule(UUID{})

				g.Assert(contains1).IsTrue()
				g.Assert(contains2).IsFalse()
			})
		})

		g.Describe("Validate", func() {
			g.It("should validate all parameters", func() {
				parameters := map[string]string{
					"foo": "bar",
				}
				validator := New(parameters)
				validator.AddRules(FooValidator{})
				response := validator.Validate()

				g.Assert(response.Code == http.StatusOK).IsTrue()
				g.Assert(len(response.Data) == 0).IsTrue()
			})

			g.It("should return \"foo_validator_error\"", func() {
				parameters := map[string]string{}
				validator := New(parameters)
				validator.AddRules(BarValidator{})
				response := validator.Validate()

				messsages, ok := response.Data["messages"].([]string)

				if !ok {
					err := errors.New("Cannot get response's messages. Failing...")
					g.Fail(err)
				}

				g.Assert(response.Code == http.StatusBadRequest).IsTrue()
				g.Assert(response.Data["length"] == 1).IsTrue()
				g.Assert(len(messsages) == 1).IsTrue()
				g.Assert(messsages[0] == "foo_validator_error").IsTrue()
			})
		})
	})
}

// Mocked validators for testing purpose
type FooValidator struct{}
type BarValidator struct{}

func (t FooValidator) validate(parameters map[string]string) error {
	return nil
}

func (t BarValidator) validate(parameters map[string]string) error {
	return errors.New("foo_validator_error")
}
