package validator

import "regexp"

// AlphaValidator implements the Validator interface for alphabetic check
type AlphaValidator struct {
	FieldName    string
	ErrorMessage string
}

var alphaRegex = regexp.MustCompile(`^[a-zA-Z]+$`)

func (v *AlphaValidator) Validate(value any) error {
	errMsg := getErrorMessage(defaultAlphaMsg, v.ErrorMessage)
	switch val := value.(type) {
	case string:
		if !alphaRegex.MatchString(val) {
			return &ValidatorError{Field: v.FieldName, MessageTemplate: errMsg}
		}
	case []string:
		if len(val) == 0 {
			return createValidatorError(v.FieldName, errMsg, val)
		}
		for _, item := range val {
			if !alphaRegex.MatchString(item) {
				return createValidatorError(v.FieldName, errMsg, item)
			}
		}
	}
	return nil
}
