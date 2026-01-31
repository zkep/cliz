package validator

import (
	"net/url"
)

type URLValidator struct {
	FieldName    string
	ErrorMessage string
}

func (v *URLValidator) Validate(value any) error {
	errMsg := getErrorMessage(defaultURLMsg, v.ErrorMessage)

	switch val := value.(type) {
	case string:
		if len(val) == 0 {
			return createValidatorError(v.FieldName, errMsg)
		}
		if u, err := url.Parse(val); err != nil || u.Scheme == "" || u.Host == "" {
			return createValidatorError(v.FieldName, errMsg)
		}
	case []string:
		if len(val) == 0 {
			return createValidatorError(v.FieldName, errMsg)
		}
		for _, item := range val {
			if len(item) == 0 {
				return createValidatorError(v.FieldName, errMsg, item)
			}
			if u, err := url.Parse(item); err != nil || u.Scheme == "" || u.Host == "" {
				return createValidatorError(v.FieldName, errMsg, item)
			}
		}
	default:
		return createValidatorError(v.FieldName, errMsg)
	}
	return nil
}
