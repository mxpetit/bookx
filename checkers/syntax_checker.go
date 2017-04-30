package checkers

// SyntaxChecker represents a set of keys that will be checked
// against the CheckFunction. Keys may be optionals by setting the optional
// parameter.
type SyntaxChecker struct {
	keys     []string
	validate CheckFunction
	optional bool
}

// check checks if values given by the SyntaxChecker's keys are
// syntaxically correct.
func (s *SyntaxChecker) check(parameters map[string]string) []string {
	var errs []string

	for _, key := range s.keys {
		errs = appendError(errs, s.validate(parameters[key], s.optional))
	}

	return errs
}
