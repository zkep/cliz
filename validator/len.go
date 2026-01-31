package validator

import "fmt"

type LenValidator struct {
	FieldName    string
	Length       int
	ErrorMessage string
}

func (v *LenValidator) Validate(value any) error {
	errMsg := getErrorMessage(defaultLenMsg, v.ErrorMessage)

	switch val := value.(type) {
	case string:
		if len(val) != v.Length {
			return createValidatorError(v.FieldName, errMsg, fmt.Sprintf("%d characters", v.Length))
		}
	case []string:
		for _, item := range val {
			if len(item) != v.Length {
				return createValidatorError(v.FieldName, errMsg, fmt.Sprintf("%d characters", v.Length))
			}
		}
	case []int, []int8, []int16, []int32, []int64, []uint, []uint8, []uint16, []uint32, []uint64:
		converted := asInt64s(val)
		if len(converted) == 0 {
			return createValidatorError(v.FieldName, errMsg, fmt.Sprintf("%d digits", v.Length))
		}
		for _, item := range converted {
			if len(fmt.Sprintf("%d", item)) != v.Length {
				return createValidatorError(v.FieldName, errMsg, fmt.Sprintf("%d digits", v.Length))
			}
		}
	case []float32, []float64:
		converted := asFloat64s(val)
		if len(converted) == 0 {
			return createValidatorError(v.FieldName, errMsg, fmt.Sprintf("%d digits", v.Length))
		}
		for _, item := range converted {
			if len(fmt.Sprintf("%f", item)) != v.Length {
				return createValidatorError(v.FieldName, errMsg, fmt.Sprintf("%d digits", v.Length))
			}
		}
	}
	return nil
}
