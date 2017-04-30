package checkers

import (
	"github.com/mxpetit/bookx/renderer"
	"net/http"
)

// checkFunctions' keys.
const (
	UUID_CHECK_FUNCTION        = "UUID"
	PAGINATION_CHECK_FUNCTION  = "PAGINATION"
	SHELVE_NAME_CHECK_FUNCTION = "SHELVE_NAME"
)

var (
	checkFunctions = map[string]CheckFunction{
		UUID_CHECK_FUNCTION:        uuidCheckFunction,
		PAGINATION_CHECK_FUNCTION:  paginationCheckFunction,
		SHELVE_NAME_CHECK_FUNCTION: shelveNameCheckFunction,
	}
)

// CheckFunction represents a function that will check if the given parameter is
// syntaxically correct. Parameter may be optional, meaning that any
// errors could be ignored.
type CheckFunction func(parameter string, optional bool) error

// SyntaxCheckerGroup represents a set of CheckFunction
// that will check parameters with their associated SyntaxChecker.
type SyntaxCheckerGroup struct {
	checkFunctions map[string]CheckFunction
	parameters     map[string]string
	checkers       []SyntaxChecker
}

// NewSyntaxCheckerGroup creates a SyntaxCheckerGroup with parameters
// and CheckFunctions given by XXX_CHECK_FUNCTION constants.
func NewSyntaxCheckerGroup(parameters map[string]string) *SyntaxCheckerGroup {
	return &SyntaxCheckerGroup{
		checkFunctions: checkFunctions,
		parameters:     parameters,
		checkers:       []SyntaxChecker{},
	}
}

// add adds a new SyntaxChecker to an existing SyntaxCheckerGroup.
func (w *SyntaxCheckerGroup) add(syntaxCheckerKey string, optional bool, keys ...string) {
	checkFunction, exists := w.checkFunctions[syntaxCheckerKey]

	if !exists {
		return
	}

	w.checkers = append(w.checkers, SyntaxChecker{
		keys:     keys,
		validate: checkFunction,
		optional: optional,
	})
}

// Add adds a SyntaxChecker where parameters are required.
func (w *SyntaxCheckerGroup) Add(syntaxCheckerKey string, keys ...string) {
	w.add(syntaxCheckerKey, false, keys...)
}

// AddOptional adds a SyntaxChecker where parameters are optional.
func (w *SyntaxCheckerGroup) AddOptional(syntaxCheckerKey string, keys ...string) {
	w.add(syntaxCheckerKey, true, keys...)
}

// Validate calls every CheckFunction associated to a SyntaxChecker
// and returns any errors that occured during validation.
func (w *SyntaxCheckerGroup) Validate() *renderer.Response {
	var errs []string

	for _, syntaxChecker := range w.checkers {
		errs = append(errs, syntaxChecker.check(w.parameters)...)
	}

	if len(errs) == 0 {
		return &renderer.Response{
			Code: http.StatusOK,
		}
	}

	return &renderer.Response{
		Code: http.StatusBadRequest,
		Data: map[string]interface{}{
			"messages": errs,
			"length":   len(errs),
		},
	}
}

// appendError appends an error to an array of error only if the error does
// exists (e.g. not nil).
func appendError(errs []string, err error) []string {
	if err != nil {
		return append(errs, err.Error())
	}

	return errs
}
