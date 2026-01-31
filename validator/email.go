package validator

import (
	"regexp"
)

var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

type EmailValidator struct {
	FieldName    string
	ErrorMessage string
}

func (v *EmailValidator) Validate(value any) error {
	errMsg := getErrorMessage(defaultEmailMsg, v.ErrorMessage)
	switch val := value.(type) {
	case string:
		if !emailRegex.MatchString(val) {
			return createValidatorError(v.FieldName, errMsg)
		}
	case []string:
		if len(val) == 0 {
			return createValidatorError(v.FieldName, errMsg)
		}
		for _, item := range val {
			if !emailRegex.MatchString(item) {
				return createValidatorError(v.FieldName, errMsg, item)
			}
		}
	default:
		return createValidatorError(v.FieldName, errMsg)
	}
	return nil
}
