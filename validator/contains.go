package validator

import (
	"strings"
)

type ContainsValidator struct {
	FieldName    string
	Substring    string
	ErrorMessage string
}

func (v *ContainsValidator) Validate(value any) error {
	errMsg := getErrorMessage(defaultContainsMsg, v.ErrorMessage)
	switch val := value.(type) {
	case string:
		if !strings.Contains(val, v.Substring) {
			return createValidatorError(v.FieldName, errMsg, v.Substring)
		}
	}
	return nil
}
