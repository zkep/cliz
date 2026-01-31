package validator

type RequiredValidator struct {
	FieldName    string
	ErrorMessage string
}

func (v *RequiredValidator) Validate(value any) error {
	errMsg := getErrorMessage(defaultRequiredMsg, v.ErrorMessage)
	err := createValidatorError(v.FieldName, errMsg)

	switch val := value.(type) {
	case string:
		if val == "" {
			return err
		}
	case bool:
		return nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		if asZero(val) {
			return err
		}
	case float32, float64:
		if asZeroFloat(val) {
			return err
		}
	case []bool:
		if len(val) == 0 {
			return err
		}
		return nil
	case []string:
		if len(val) == 0 || containsEmptyString(val) {
			return err
		}
	case []int, []int8, []int16, []int32, []int64, []uint, []uint8, []uint16, []uint32, []uint64:
		converted := asInt64s(val)
		if len(converted) == 0 {
			return err
		}
		if containsZeroInt64(converted) {
			return err
		}
	case []float64, []float32:
		converted := asFloat64s(val)
		if len(converted) == 0 {
			return err
		}
		if containsZeroFloat64(converted) {
			return err
		}
	default:
		return err
	}
	return nil
}
