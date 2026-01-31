package validator

import "fmt"

type InValidator struct {
	FieldName    string
	Allowed      []string
	ErrorMessage string
}

func (v *InValidator) Validate(value any) error {
	inTemplate := ""
	for k := range v.Allowed {
		if k == 0 {
			inTemplate = "%v"
		} else {
			inTemplate += ",%v"
		}
	}
	defaultMsg := fmt.Sprintf(defaultInMsg, inTemplate)
	errMsg := getErrorMessage(defaultMsg, v.ErrorMessage)
	allowedAny := make([]any, len(v.Allowed))
	for k, v := range v.Allowed {
		allowedAny[k] = v
	}
	switch val := value.(type) {
	case string:
		if !containsString(v.Allowed, val) {
			return createValidatorError(v.FieldName, errMsg, allowedAny...)
		}
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		if !containsInt64(v.Allowed, asInt64(val)) {
			return createValidatorError(v.FieldName, errMsg, allowedAny...)
		}
	case float32, float64:
		if !containsFloat64(v.Allowed, asFloat64(val)) {
			return createValidatorError(v.FieldName, errMsg, allowedAny...)
		}
	case []string:
		if len(v.Allowed) == 0 {
			return createValidatorError(v.FieldName, errMsg, allowedAny...)
		}
		for _, item := range val {
			if !containsString(v.Allowed, item) {
				return createValidatorError(v.FieldName, errMsg, allowedAny...)
			}
		}
	case []int, []int8, []int16, []int32, []int64, []uint, []uint8, []uint16, []uint32, []uint64:
		if len(v.Allowed) == 0 {
			return createValidatorError(v.FieldName, errMsg, allowedAny...)
		}
		converted := asInt64s(val)
		if len(converted) == 0 {
			return createValidatorError(v.FieldName, errMsg, allowedAny...)
		}
		for _, item := range converted {
			if !containsInt64(v.Allowed, item) {
				return createValidatorError(v.FieldName, errMsg, allowedAny...)
			}
		}
	case []float32, []float64:
		if len(v.Allowed) == 0 {
			return createValidatorError(v.FieldName, errMsg, allowedAny...)
		}
		converted := asFloat64s(val)
		if len(converted) == 0 {
			return createValidatorError(v.FieldName, errMsg, allowedAny...)
		}
		for _, item := range converted {
			if !containsFloat64(v.Allowed, item) {
				return createValidatorError(v.FieldName, errMsg, allowedAny...)
			}
		}
	}
	return nil
}
