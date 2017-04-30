package checkers

import (
	"errors"
	"github.com/franela/goblin"
	"net/http"
	"reflect"
	"testing"
)

func TestSyntaxCheckerGroup(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("package checkers > syntax_checker_group", func() {
		g.Describe("function appendError", func() {
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

		g.Describe("function Add", func() {
			g.It("should add a SyntaxChecker, where his parameters are required, to a SyntaxCheckerGroup", func() {
				paginationKeys := []string{"limit", "pagination"}

				syntaxCheckerGroup := NewSyntaxCheckerGroup(map[string]string{})
				syntaxCheckerGroup.Add(PAGINATION_CHECK_FUNCTION, paginationKeys...)

				g.Assert(len(syntaxCheckerGroup.checkers) == 1).IsTrue()
				g.Assert(reflect.DeepEqual(syntaxCheckerGroup.checkers[0].keys, paginationKeys)).IsTrue()
				g.Assert(syntaxCheckerGroup.checkers[0].optional).IsFalse()
				g.Assert(syntaxCheckerGroup.checkers[0].validate != nil).IsTrue()
			})
		})

		g.Describe("function AddOptional", func() {
			g.It("should add a SyntaxChecker, where his parameters are optional, to a SyntaxCheckerGroup", func() {
				uuidKeys := []string{"uuid", "id"}

				syntaxCheckerGroup := NewSyntaxCheckerGroup(map[string]string{})
				syntaxCheckerGroup.AddOptional(UUID_CHECK_FUNCTION, uuidKeys...)

				g.Assert(len(syntaxCheckerGroup.checkers) == 1).IsTrue()
				g.Assert(reflect.DeepEqual(syntaxCheckerGroup.checkers[0].keys, uuidKeys)).IsTrue()
				g.Assert(syntaxCheckerGroup.checkers[0].optional).IsTrue()
				g.Assert(syntaxCheckerGroup.checkers[0].validate != nil).IsTrue()
			})
		})

		g.Describe("function add", func() {
			g.It("should add a SyntaxChecker to a SyntaxCheckerGroup", func() {
				parameters := []string{"parameter1", "parameter2"}

				syntaxCheckerGroup := NewSyntaxCheckerGroup(map[string]string{})
				syntaxCheckerGroup.add(UUID_CHECK_FUNCTION, true, parameters...)
				syntaxCheckerGroup.add(PAGINATION_CHECK_FUNCTION, false, parameters...)

				g.Assert(len(syntaxCheckerGroup.checkers) == 2).IsTrue()
				g.Assert(reflect.DeepEqual(syntaxCheckerGroup.checkers[0].keys, parameters)).IsTrue()
				g.Assert(syntaxCheckerGroup.checkers[0].optional).IsTrue()
				g.Assert(syntaxCheckerGroup.checkers[0].validate != nil).IsTrue()
				g.Assert(reflect.DeepEqual(syntaxCheckerGroup.checkers[1].keys, parameters)).IsTrue()
				g.Assert(syntaxCheckerGroup.checkers[1].optional).IsFalse()
				g.Assert(syntaxCheckerGroup.checkers[1].validate != nil).IsTrue()
			})

			g.It("shouldn't add a SyntaxChecker to a SyntaxCheckerGroup since the CheckFunction does not exists", func() {
				parameters := []string{"parameter1", "parameter2"}

				syntaxCheckerGroup := NewSyntaxCheckerGroup(map[string]string{})
				syntaxCheckerGroup.add("check_function_that_doesnt_exists", true, parameters...)

				g.Assert(len(syntaxCheckerGroup.checkers) == 0).IsTrue()
			})
		})

		g.Describe("function Validate", func() {
			g.It("shouldn't return an error since there is no parameter", func() {
				syntaxCheckerGroup := NewSyntaxCheckerGroup(map[string]string{})
				response := syntaxCheckerGroup.Validate()

				g.Assert(response.Code == http.StatusOK).IsTrue()
				g.Assert(len(response.Data) == 0).IsTrue()
			})

			g.It("should return an error since required parameter is missing", func() {
				parameter := []string{"parameter"}

				syntaxCheckerGroup := NewSyntaxCheckerGroup(map[string]string{})
				syntaxCheckerGroup.Add(PAGINATION_CHECK_FUNCTION, parameter...)
				response := syntaxCheckerGroup.Validate()

				messages, ok := response.Data["messages"].([]string)

				if !ok {
					g.Fail("Cannot type cast response.Data[\"messages\"] as array of string.")
				}

				g.Assert(response.Code == http.StatusBadRequest).IsTrue()
				g.Assert(response.Data["length"] == 1).IsTrue()
				g.Assert(len(messages) == 1).IsTrue()
				g.Assert(messages[0] == ErrLimitIsMissing.Error()).IsTrue()
			})
		})
	})
}
