package validator

import "regexp"

// AlphanumValidator implements the Validator interface for alphanumeric check
type AlphanumValidator struct {
	FieldName    string
	ErrorMessage string
}

var alphanumRegex = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

func (v *AlphanumValidator) Validate(value any) error {
	errMsg := getErrorMessage(defaultAlphanumMsg, v.ErrorMessage)
	switch val := value.(type) {
	case string:
		if !alphanumRegex.MatchString(val) {
			return &ValidatorError{Field: v.FieldName, MessageTemplate: errMsg}
		}
	case []string:
		if len(val) == 0 {
			return createValidatorError(v.FieldName, errMsg, val)
		}
		for _, item := range val {
			if !alphanumRegex.MatchString(item) {
				return createValidatorError(v.FieldName, errMsg, item)
			}
		}
	}
	return nil
}
