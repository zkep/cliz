package validator

type RangeValidator struct {
	FieldName    string
	Min          float64
	Max          float64
	ErrorMessage string
}

func (r *RangeValidator) Validate(value any) error {
	errMsg := getErrorMessage(defaultRangeMsg, r.ErrorMessage)

	switch val := value.(type) {
	case string:
		num := asFloat64(val)
		if num < r.Min || num > r.Max {
			return createValidatorError(r.FieldName, errMsg, r.Min, r.Max)
		}
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		num := asFloat64(val)
		if num < r.Min || num > r.Max {
			return createValidatorError(r.FieldName, errMsg, r.Min, r.Max)
		}
	case []string:
		if len(val) == 0 {
			return createValidatorError(r.FieldName, errMsg, r.Min, r.Max)
		}
		for _, item := range val {
			num := asFloat64(item)
			if num < r.Min || num > r.Max {
				return createValidatorError(r.FieldName, errMsg, r.Min, r.Max)
			}
		}
	case []int, []int8, []int16, []int32, []int64, []uint, []uint8, []uint16, []uint32, []uint64:
		converted := asInt64s(val)
		if len(converted) == 0 {
			return createValidatorError(r.FieldName, errMsg, r.Min, r.Max)
		}
		for _, item := range converted {
			num := float64(item)
			if num < r.Min || num > r.Max {
				return createValidatorError(r.FieldName, errMsg, r.Min, r.Max)
			}
		}
	case []float32, []float64:
		converted := asFloat64s(val)
		if len(converted) == 0 {
			return createValidatorError(r.FieldName, errMsg, r.Min, r.Max)
		}
		for _, item := range converted {
			if item < r.Min || item > r.Max {
				return createValidatorError(r.FieldName, errMsg, r.Min, r.Max)
			}
		}
	}
	return nil
}
