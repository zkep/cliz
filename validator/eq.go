package validator

import "fmt"

type EqValidator struct {
	FieldName    string
	Value        any
	ErrorMessage string
}

func (v *EqValidator) Validate(value any) error {
	errMsg := getErrorMessage(defaultEqMsg, v.ErrorMessage)
	switch val := value.(type) {
	case string:
		if eqStr, ok := v.Value.(string); ok && val != eqStr {
			return createValidatorError(v.FieldName, errMsg, val)
		}
	case bool:
		if eqBool, ok := v.Value.(bool); ok && val != eqBool {
			return createValidatorError(v.FieldName, errMsg, fmt.Sprintf(formatBoolean, val))
		}
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		if asFloat64(val) != asFloat64(v.Value) {
			return createValidatorError(v.FieldName, errMsg, val)
		}
	case []string:
		if len(val) == 0 {
			return createValidatorError(v.FieldName, errMsg, v.Value)
		}
		for _, item := range val {
			if item != v.Value.(string) {
				return createValidatorError(v.FieldName, errMsg, item)
			}
		}
	case []bool:
		if len(val) == 0 {
			return createValidatorError(v.FieldName, errMsg, fmt.Sprintf(formatBoolean, v.Value))
		}
		for _, item := range val {
			if item != v.Value.(bool) {
				return createValidatorError(v.FieldName, errMsg, fmt.Sprintf(formatBoolean, item))
			}
		}
	case []int, []int8, []int16, []int32, []int64, []uint, []uint8, []uint16, []uint32, []uint64, []float64, []float32:
		converted := asFloat64s(val)
		if len(converted) == 0 {
			return createValidatorError(v.FieldName, errMsg, v.Value)
		}
		for _, item := range converted {
			if item != asFloat64(v.Value) {
				return createValidatorError(v.FieldName, errMsg, item)
			}
		}
	}
	return nil
}
