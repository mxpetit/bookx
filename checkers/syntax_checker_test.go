package checkers

import (
	"errors"
	"github.com/franela/goblin"
	"testing"
)

var (
	RequiredErr   = errors.New("required")
	OptionalErr   = errors.New("optional")
	UnexecptedErr = errors.New("unexcepted")
)

func TestSyntaxChecker(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("package checkers > syntaxChecker", func() {
		g.Describe("function check", func() {
			g.Describe("when parameters are optional", func() {
				g.It("shouldn't return an error since the parameter are not provided", func() {
					checker := SyntaxChecker{
						keys:     []string{"foo", "bar"},
						validate: optionalFunctionWithoutError,
						optional: true,
					}

					errs := checker.check(map[string]string{})

					g.Assert(len(errs) == 0).IsTrue()
				})

				g.It("should return an error for each parameters since they are provided and syntaxically wrong", func() {
					parameters := map[string]string{
						"foo": "qux",
						"bar": "gault",
					}

					checker := SyntaxChecker{
						keys:     []string{"foo", "bar"},
						validate: optionalFunctionWithError,
						optional: true,
					}

					errs := checker.check(parameters)

					g.Assert(len(errs) == len(checker.keys)).IsTrue()
					g.Assert(errs[0] == OptionalErr.Error()).IsTrue()
					g.Assert(errs[1] == OptionalErr.Error()).IsTrue()
				})

				g.It("shouldn't return an error for each parameters since they are provided and syntaxically correct", func() {
					parameters := map[string]string{
						"foo": "qux",
						"bar": "gault",
					}

					checker := SyntaxChecker{
						keys:     []string{"foo", "bar"},
						validate: optionalFunctionWithoutError,
						optional: true,
					}

					errs := checker.check(parameters)

					g.Assert(len(errs) == 0).IsTrue()
				})
			})

			g.Describe("when parameters are required", func() {
				g.It("should return an error since the parameter are not provided", func() {
					checker := SyntaxChecker{
						keys:     []string{"foo", "bar"},
						validate: requiredFunctionWithError,
						optional: true,
					}

					errs := checker.check(map[string]string{})

					g.Assert(len(errs) == len(checker.keys)).IsTrue()
					g.Assert(errs[0] == RequiredErr.Error()).IsTrue()
					g.Assert(errs[1] == RequiredErr.Error()).IsTrue()
				})

				g.It("should return an error for each parameters since they are provided and syntaxically wrong", func() {
					parameters := map[string]string{
						"foo": "qux",
						"bar": "gault",
					}

					checker := SyntaxChecker{
						keys:     []string{"foo", "bar"},
						validate: requiredFunctionWithError,
						optional: true,
					}

					errs := checker.check(parameters)

					g.Assert(len(errs) == len(checker.keys)).IsTrue()
					g.Assert(errs[0] == RequiredErr.Error()).IsTrue()
					g.Assert(errs[1] == RequiredErr.Error()).IsTrue()
				})

				g.It("shouldn't return an error for each parameters since they are provided and syntaxically correct", func() {
					parameters := map[string]string{
						"foo": "qux",
						"bar": "gault",
					}

					checker := SyntaxChecker{
						keys:     []string{"foo", "bar"},
						validate: requiredFunctionWithoutError,
						optional: true,
					}

					errs := checker.check(parameters)

					g.Assert(len(errs) == 0).IsTrue()
				})
			})
		})
	})
}

func optionalFunctionWithoutError(parameter string, optional bool) error {
	if optional {
		return nil
	}

	return UnexecptedErr
}

func optionalFunctionWithError(parameter string, optional bool) error {
	if optional {
		return OptionalErr
	}

	return UnexecptedErr
}

func requiredFunctionWithoutError(parameter string, optional bool) error {
	return nil
}

func requiredFunctionWithError(parameter string, optional bool) error {
	return RequiredErr
}
