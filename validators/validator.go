package validators

import (
	"github.com/mxpetit/bookx/renderer"
	"net/http"
)

// ValidatorRule represents a rule that will semantically checks
// the correctness of a portion of parameters.
type ValidatorRule interface {
	validate(parameters *Parameters) error
}

// Parameters wraps a map of string / interface.
type Parameters map[string]interface{}

// Validator represents a set of ValidatorRule and it's associated
// parameters.
type Validator struct {
	rules      []ValidatorRule
	parameters *Parameters
}

// New instantiates a new Validator.
func New(parameters *Parameters) *Validator {
	return &Validator{
		rules:      []ValidatorRule{},
		parameters: parameters,
	}
}

// Validate validates every rules of a Validator.
func (s *Validator) Validate() renderer.Response {
	var errs []string

	for _, rule := range s.rules {
		errs = appendError(errs, rule.validate(s.parameters))
	}

	if len(errs) == 0 {
		return renderer.Response{
			Code: http.StatusOK,
		}
	}

	return renderer.Response{
		Code: http.StatusBadRequest,
		Data: map[string]interface{}{
			"messages": errs,
			"length":   len(errs),
		},
	}
}

// appendError appends error to an array of error only if the error does
// exists.
func appendError(errs []string, err error) []string {
	if err != nil {
		return append(errs, err.Error())
	}

	return errs
}

// AddRules adds one or more ValidatorRule to a Validator without duplicates.
func (s *Validator) AddRules(rules ...ValidatorRule) {
	for _, value := range rules {
		if !s.containsRule(value) {
			s.rules = append(s.rules, value)
		}
	}
}

// containsRule checks wether the rule that will be added is
// already registered.
func (s *Validator) containsRule(ruleToAdd ValidatorRule) bool {
	contains := false

	for i := 0; i < len(s.rules) && !contains; i++ {
		if s.rules[i] == ruleToAdd {
			contains = true
		}
	}

	return contains
}
