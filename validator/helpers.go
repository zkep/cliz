package validator

import (
	"strconv"
)

func getErrorMessage(defaultMsg, customMsg string) string {
	if customMsg != "" {
		return customMsg
	}
	return defaultMsg
}

func createValidatorError(fieldName, msgTemplate string, args ...any) *ValidatorError {
	return &ValidatorError{
		Field:           fieldName,
		MessageTemplate: msgTemplate,
		Args:            args,
	}
}

func asInt64(value any) int64 {
	switch val := value.(type) {
	case int:
		return int64(val)
	case int8:
		return int64(val)
	case int16:
		return int64(val)
	case int32:
		return int64(val)
	case int64:
		return val
	case uint:
		return int64(val)
	case uint8:
		return int64(val)
	case uint16:
		return int64(val)
	case uint32:
		return int64(val)
	case uint64:
		return int64(val)
	case float32:
		return int64(val)
	case float64:
		return int64(val)
	}
	return 0
}

func asInt64s(values any) []int64 {
	switch val := values.(type) {
	case []int:
		int64s := make([]int64, len(val))
		for i, item := range val {
			int64s[i] = int64(item)
		}
		return int64s
	case []int8:
		int64s := make([]int64, len(val))
		for i, item := range val {
			int64s[i] = int64(item)
		}
		return int64s
	case []int16:
		int64s := make([]int64, len(val))
		for i, item := range val {
			int64s[i] = int64(item)
		}
		return int64s
	case []int32:
		int64s := make([]int64, len(val))
		for i, item := range val {
			int64s[i] = int64(item)
		}
		return int64s
	case []int64:
		return val
	case []uint:
		int64s := make([]int64, len(val))
		for i, item := range val {
			int64s[i] = int64(item)
		}
		return int64s
	case []uint8:
		int64s := make([]int64, len(val))
		for i, item := range val {
			int64s[i] = int64(item)
		}
		return int64s
	case []uint16:
		int64s := make([]int64, len(val))
		for i, item := range val {
			int64s[i] = int64(item)
		}
		return int64s
	case []uint32:
		int64s := make([]int64, len(val))
		for i, item := range val {
			int64s[i] = int64(item)
		}
		return int64s
	case []uint64:
		int64s := make([]int64, len(val))
		for i, item := range val {
			int64s[i] = int64(item)
		}
		return int64s
	}
	return nil
}

func asFloat64(value any) float64 {
	switch val := value.(type) {
	case int:
		return float64(val)
	case int8:
		return float64(val)
	case int16:
		return float64(val)
	case int32:
		return float64(val)
	case int64:
		return float64(val)
	case uint:
		return float64(val)
	case uint8:
		return float64(val)
	case uint16:
		return float64(val)
	case uint32:
		return float64(val)
	case uint64:
		return float64(val)
	case float32:
		return float64(val)
	case float64:
		return val
	case string:
		num, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return 0
		}
		return num
	}
	return 0
}

func asFloat64s(values any) []float64 {
	switch val := values.(type) {
	case []int:
		floats := make([]float64, len(val))
		for i, item := range val {
			floats[i] = float64(item)
		}
		return floats
	case []int8:
		floats := make([]float64, len(val))
		for i, item := range val {
			floats[i] = float64(item)
		}
		return floats
	case []int16:
		floats := make([]float64, len(val))
		for i, item := range val {
			floats[i] = float64(item)
		}
		return floats
	case []int32:
		floats := make([]float64, len(val))
		for i, item := range val {
			floats[i] = float64(item)
		}
		return floats
	case []int64:
		floats := make([]float64, len(val))
		for i, item := range val {
			floats[i] = float64(item)
		}
		return floats
	case []uint:
		floats := make([]float64, len(val))
		for i, item := range val {
			floats[i] = float64(item)
		}
		return floats
	case []uint8:
		floats := make([]float64, len(val))
		for i, item := range val {
			floats[i] = float64(item)
		}
		return floats
	case []uint16:
		floats := make([]float64, len(val))
		for i, item := range val {
			floats[i] = float64(item)
		}
		return floats
	case []uint32:
		floats := make([]float64, len(val))
		for i, item := range val {
			floats[i] = float64(item)
		}
		return floats
	case []uint64:
		floats := make([]float64, len(val))
		for i, item := range val {
			floats[i] = float64(item)
		}
		return floats
	case []float64:
		return val
	case []float32:
		floats := make([]float64, len(val))
		for i, item := range val {
			floats[i] = float64(item)
		}
		return floats
	case []string:
		floats := make([]float64, len(val))
		for i, item := range val {
			floats[i] = asFloat64(item)
		}
		return floats
	}
	return nil
}

func asZero(value any) bool {
	switch val := value.(type) {
	case int:
		return val == 0
	case int8:
		return val == 0
	case int16:
		return val == 0
	case int32:
		return val == 0
	case int64:
		return val == 0
	case uint:
		return val == 0
	case uint8:
		return val == 0
	case uint16:
		return val == 0
	case uint32:
		return val == 0
	case uint64:
		return val == 0
	}
	return false
}

func asZeroFloat(value any) bool {
	switch val := value.(type) {
	case float32:
		return val == 0.0
	case float64:
		return val == 0.0
	}
	return false
}

func containsEmptyString(arr []string) bool {
	for _, item := range arr {
		if item == "" {
			return true
		}
	}
	return false
}

func containsZeroInt64(arr []int64) bool {
	for _, item := range arr {
		if item == 0 {
			return true
		}
	}
	return false
}

func containsZeroFloat64(arr []float64) bool {
	for _, item := range arr {
		if item == 0 {
			return true
		}
	}
	return false
}

func containsString(arr []string, s string) bool {
	for _, item := range arr {
		if item == s {
			return true
		}
	}
	return false
}

func containsInt64(arr []string, val int64) bool {
	for _, allowed := range arr {
		if allowedInt, err := strconv.ParseInt(allowed, 10, 64); err == nil && val == allowedInt {
			return true
		}
	}
	return false
}

func containsFloat64(arr []string, val float64) bool {
	for _, allowed := range arr {
		if allowedFloat, err := strconv.ParseFloat(allowed, 64); err == nil && val == allowedFloat {
			return true
		}
	}
	return false
}
