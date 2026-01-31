package validator

type GtValidator struct {
	FieldName    string
	Value        float64
	ErrorMessage string
}

func (v *GtValidator) Validate(value any) error {
	errMsg := getErrorMessage(defaultGtMsg, v.ErrorMessage)

	switch val := value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		if asFloat64(val) <= v.Value {
			return createValidatorError(v.FieldName, errMsg, v.Value)
		}
	case []int, []int8, []int16, []int32, []int64, []uint, []uint8, []uint16, []uint32, []uint64, []float64, []float32:
		converted := asFloat64s(val)
		if len(converted) == 0 {
			return createValidatorError(v.FieldName, errMsg, v.Value)
		}
		for _, item := range converted {
			if item <= v.Value {
				return createValidatorError(v.FieldName, errMsg, v.Value)
			}
		}
	}
	return nil
}
