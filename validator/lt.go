package validator

import "fmt"

type LtValidator struct {
	FieldName    string
	Value        float64
	ErrorMessage string
}

func (v *LtValidator) Validate(value any) error {
	errMsg := getErrorMessage(defaultLtMsg, v.ErrorMessage)

	switch val := value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		if asFloat64(val) >= v.Value {
			return createValidatorError(v.FieldName, errMsg, []any{fmt.Sprintf(formatInteger, v.Value)})
		}
	case float32, float64:
		if asFloat64(val) >= v.Value {
			return createValidatorError(v.FieldName, errMsg, []any{fmt.Sprintf(formatFloat, v.Value)})
		}
	case []int, []int8, []int16, []int32, []int64, []uint, []uint8, []uint16, []uint32, []uint64:
		for _, item := range asInt64s(val) {
			if float64(item) >= v.Value {
				return createValidatorError(v.FieldName, errMsg, []any{fmt.Sprintf(formatInteger, v.Value)})
			}
		}
	case []float32, []float64:
		float64s := asFloat64s(val)
		if len(float64s) == 0 {
			return createValidatorError(v.FieldName, errMsg, []any{fmt.Sprintf(formatItems, v.Value)})
		}
		for _, item := range float64s {
			if item >= v.Value {
				return createValidatorError(v.FieldName, errMsg, []any{fmt.Sprintf(formatFloat, v.Value)})
			}
		}
	}
	return nil
}
