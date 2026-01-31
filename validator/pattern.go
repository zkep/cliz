package validator

import (
	"regexp"
)

type PatternValidator struct {
	FieldName    string
	Pattern      string
	ErrorMessage string
	*regexp.Regexp
}

func (v *PatternValidator) Validate(value any) error {
	errMsg := getErrorMessage(defaultPatternMsg, v.ErrorMessage)

	if v.Regexp == nil {
		return nil
	}

	switch val := value.(type) {
	case string:
		if !v.Regexp.MatchString(val) {
			return createValidatorError(v.FieldName, errMsg, v.Pattern)
		}
	}
	return nil
}
