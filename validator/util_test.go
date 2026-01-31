package validator

import (
	"testing"
)

func TestValidatorInterface(t *testing.T) {
	var v Validator = ValidatorFunc(func(value any) error {
		return nil
	})
	if err := v.Validate("test"); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestValidatorError(t *testing.T) {
	err := &ValidatorError{
		Field:           "test-field",
		Message:         "default error",
		MessageTemplate: "custom error with %v",
		Args:            []any{"arg"},
	}
	if err.Error() != "test-field: custom error with arg" {
		t.Fatalf("Expected error message 'test-field: custom error with arg', got '%s'", err.Error())
	}
}

func TestValidatorErrorWithNoArgs(t *testing.T) {
	err := &ValidatorError{
		Field:           "test-field",
		Message:         "default error",
		MessageTemplate: "custom error",
		Args:            []any{},
	}
	if err.Error() != "test-field: custom error" {
		t.Fatalf("Expected error message 'test-field: custom error', got '%s'", err.Error())
	}
}

func TestValidatorErrorWithDefaultMessage(t *testing.T) {
	err := &ValidatorError{
		Field:   "test-field",
		Message: "default error",
	}
	if err.Error() != "test-field: default error" {
		t.Fatalf("Expected error message 'test-field: default error', got '%s'", err.Error())
	}
}

func TestCountPlaceholders(t *testing.T) {
	if countPlaceholders("%v %v") != 2 {
		t.Fatalf("Expected 2 placeholders, got %d", countPlaceholders("%v %v"))
	}
	if countPlaceholders("no placeholders") != 0 {
		t.Fatalf("Expected 0 placeholders, got %d", countPlaceholders("no placeholders"))
	}
	if countPlaceholders("%d %f %s") != 3 {
		t.Fatalf("Expected 3 placeholders, got %d", countPlaceholders("%d %f %s"))
	}
}

func TestAsInt64(t *testing.T) {
	if asInt64(10) != 10 {
		t.Fatalf("Expected 10, got %d", asInt64(10))
	}
	if asInt64(10.5) != 10 {
		t.Fatalf("Expected 10, got %d", asInt64(10.5))
	}
	if asInt64("not a number") != 0 {
		t.Fatalf("Expected 0, got %d", asInt64("not a number"))
	}
}

func TestAsZero(t *testing.T) {
	if !asZero(0) {
		t.Fatal("Expected true for 0")
	}
	if asZero(1) {
		t.Fatal("Expected false for 1")
	}
	if asZero("not a number") {
		t.Fatal("Expected false for non-number")
	}
}

func TestAsZeroFloat(t *testing.T) {
	if !asZeroFloat(0.0) {
		t.Fatal("Expected true for 0.0")
	}
	if asZeroFloat(1.0) {
		t.Fatal("Expected false for 1.0")
	}
	if asZeroFloat("not a float") {
		t.Fatal("Expected false for non-float")
	}
}

func TestAsInt64s(t *testing.T) {
	if len(asInt64s([]int{1, 2, 3})[:]) != 3 {
		t.Fatalf("Expected 3 elements, got %d", len(asInt64s([]int{1, 2, 3})[:]))
	}
	if asInt64s([]uint{1, 2, 3}) == nil {
		t.Fatal("Expected non-nil slice")
	}
	if asInt64s("not a slice") != nil {
		t.Fatal("Expected nil for non-slice")
	}
}

func TestAsFloat64(t *testing.T) {
	if asFloat64(10) != 10.0 {
		t.Fatalf("Expected 10.0, got %f", asFloat64(10))
	}
	if asFloat64("10.5") != 10.5 {
		t.Fatalf("Expected 10.5, got %f", asFloat64("10.5"))
	}
	if asFloat64("not a number") != 0.0 {
		t.Fatalf("Expected 0.0, got %f", asFloat64("not a number"))
	}
}

func TestAsFloat64s(t *testing.T) {
	if len(asFloat64s([]int{1, 2, 3})[:]) != 3 {
		t.Fatalf("Expected 3 elements, got %d", len(asFloat64s([]int{1, 2, 3})[:]))
	}
	if asFloat64s([]string{"1.0", "2.0"}) == nil {
		t.Fatal("Expected non-nil slice")
	}
	if asFloat64s("not a slice") != nil {
		t.Fatal("Expected nil for non-slice")
	}
}

func TestContainsString(t *testing.T) {
	if !containsString([]string{"a", "b", "c"}, "b") {
		t.Fatal("Expected true for string 'b'")
	}
	if containsString([]string{"a", "b", "c"}, "d") {
		t.Fatal("Expected false for string 'd'")
	}
}

func TestCountPlaceholdersComplex(t *testing.T) {
	if countPlaceholders("%v%%%v") != 2 {
		t.Fatalf("Expected 2 placeholders, got %d", countPlaceholders("%v%%%v"))
	}
	if countPlaceholders("%d %s %f") != 3 {
		t.Fatalf("Expected 3 placeholders, got %d", countPlaceholders("%d %s %f"))
	}
	if countPlaceholders("no % in string") != 0 {
		t.Fatalf("Expected 0 placeholders, got %d", countPlaceholders("no % in string"))
	}
}

func TestValidatorErrorWithPartialArgs(t *testing.T) {
	err := &ValidatorError{
		Field:           "test-field",
		MessageTemplate: "%v %v %v",
		Args:            []any{"arg1", "arg2"},
	}
	if err.Error() == "" {
		t.Fatal("Expected non-empty error message")
	}
}

func TestValidatorErrorWithTooManyArgs(t *testing.T) {
	err := &ValidatorError{
		Field:           "test-field",
		MessageTemplate: "%v",
		Args:            []any{"arg1", "arg2", "arg3"},
	}
	if err.Error() == "" {
		t.Fatal("Expected non-empty error message")
	}
}

func TestCreateValidatorError(t *testing.T) {
	err := createValidatorError("field", "template %v", "value")
	if err.Field != "field" {
		t.Fatalf("Expected field 'field', got '%s'", err.Field)
	}
	if err.MessageTemplate != "template %v" {
		t.Fatalf("Expected template 'template v', got '%s'", err.MessageTemplate)
	}
	if len(err.Args) != 1 {
		t.Fatalf("Expected 1 arg, got %d", len(err.Args))
	}
	if err.Args[0] != "value" {
		t.Fatalf("Expected arg 'value', got '%v'", err.Args[0])
	}
}

func TestAsInt64AllTypes(t *testing.T) {
	if asInt64(int(1)) != 1 {
		t.Fatalf("Expected 1 for int")
	}
	if asInt64(int8(2)) != 2 {
		t.Fatalf("Expected 2 for int8")
	}
	if asInt64(int16(3)) != 3 {
		t.Fatalf("Expected 3 for int16")
	}
	if asInt64(int32(4)) != 4 {
		t.Fatalf("Expected 4 for int32")
	}
	if asInt64(int64(5)) != 5 {
		t.Fatalf("Expected 5 for int64")
	}
	if asInt64(uint(6)) != 6 {
		t.Fatalf("Expected 6 for uint")
	}
	if asInt64(uint8(7)) != 7 {
		t.Fatalf("Expected 7 for uint8")
	}
	if asInt64(uint16(8)) != 8 {
		t.Fatalf("Expected 8 for uint16")
	}
	if asInt64(uint32(9)) != 9 {
		t.Fatalf("Expected 9 for uint32")
	}
	if asInt64(uint64(10)) != 10 {
		t.Fatalf("Expected 10 for uint64")
	}
	if asInt64(float32(11.5)) != 11 {
		t.Fatalf("Expected 11 for float32")
	}
	if asInt64(float64(12.9)) != 12 {
		t.Fatalf("Expected 12 for float64")
	}
	if asInt64("not a number") != 0 {
		t.Fatalf("Expected 0 for non-number")
	}
}

func TestAsInt64sAllTypes(t *testing.T) {
	if len(asInt64s([]int8{1, 2, 3})[:]) != 3 {
		t.Fatalf("Expected 3 elements for int8 slice, got %d", len(asInt64s([]int8{1, 2, 3})[:]))
	}
	if len(asInt64s([]int16{1, 2, 3})[:]) != 3 {
		t.Fatalf("Expected 3 elements for int16 slice, got %d", len(asInt64s([]int16{1, 2, 3})[:]))
	}
	if len(asInt64s([]int32{1, 2, 3})[:]) != 3 {
		t.Fatalf("Expected 3 elements for int32 slice, got %d", len(asInt64s([]int32{1, 2, 3})[:]))
	}
	if len(asInt64s([]int64{1, 2, 3})[:]) != 3 {
		t.Fatalf("Expected 3 elements for int64 slice, got %d", len(asInt64s([]int64{1, 2, 3})[:]))
	}
	if len(asInt64s([]uint8{1, 2, 3})[:]) != 3 {
		t.Fatalf("Expected 3 elements for uint8 slice, got %d", len(asInt64s([]uint8{1, 2, 3})[:]))
	}
	if len(asInt64s([]uint16{1, 2, 3})[:]) != 3 {
		t.Fatalf("Expected 3 elements for uint16 slice, got %d", len(asInt64s([]uint16{1, 2, 3})[:]))
	}
	if len(asInt64s([]uint32{1, 2, 3})[:]) != 3 {
		t.Fatalf("Expected 3 elements for uint32 slice, got %d", len(asInt64s([]uint32{1, 2, 3})[:]))
	}
	if len(asInt64s([]uint64{1, 2, 3})[:]) != 3 {
		t.Fatalf("Expected 3 elements for uint64 slice, got %d", len(asInt64s([]uint64{1, 2, 3})[:]))
	}
}

func TestAsFloat64AllTypes(t *testing.T) {
	if asFloat64(int(1)) != 1.0 {
		t.Fatalf("Expected 1.0 for int")
	}
	if asFloat64(int8(2)) != 2.0 {
		t.Fatalf("Expected 2.0 for int8")
	}
	if asFloat64(int16(3)) != 3.0 {
		t.Fatalf("Expected 3.0 for int16")
	}
	if asFloat64(int32(4)) != 4.0 {
		t.Fatalf("Expected 4.0 for int32")
	}
	if asFloat64(int64(5)) != 5.0 {
		t.Fatalf("Expected 5.0 for int64")
	}
	if asFloat64(uint(6)) != 6.0 {
		t.Fatalf("Expected 6.0 for uint")
	}
	if asFloat64(uint8(7)) != 7.0 {
		t.Fatalf("Expected 7.0 for uint8")
	}
	if asFloat64(uint16(8)) != 8.0 {
		t.Fatalf("Expected 8.0 for uint16")
	}
	if asFloat64(uint32(9)) != 9.0 {
		t.Fatalf("Expected 9.0 for uint32")
	}
	if asFloat64(uint64(10)) != 10.0 {
		t.Fatalf("Expected 10.0 for uint64")
	}
	if asFloat64(float32(11.5)) != 11.5 {
		t.Fatalf("Expected 11.5 for float32")
	}
	if asFloat64(float64(12.9)) != 12.9 {
		t.Fatalf("Expected 12.9 for float64")
	}
	if asFloat64("13.5") != 13.5 {
		t.Fatalf("Expected 13.5 for string")
	}
	if asFloat64("not a number") != 0.0 {
		t.Fatalf("Expected 0.0 for non-number")
	}
}

func TestAsFloat64sAllTypes(t *testing.T) {
	if len(asFloat64s([]int8{1, 2, 3})[:]) != 3 {
		t.Fatalf("Expected 3 elements for int8 slice, got %d", len(asFloat64s([]int8{1, 2, 3})[:]))
	}
	if len(asFloat64s([]int16{1, 2, 3})[:]) != 3 {
		t.Fatalf("Expected 3 elements for int16 slice, got %d", len(asFloat64s([]int16{1, 2, 3})[:]))
	}
	if len(asFloat64s([]int32{1, 2, 3})[:]) != 3 {
		t.Fatalf("Expected 3 elements for int32 slice, got %d", len(asFloat64s([]int32{1, 2, 3})[:]))
	}
	if len(asFloat64s([]int64{1, 2, 3})[:]) != 3 {
		t.Fatalf("Expected 3 elements for int64 slice, got %d", len(asFloat64s([]int64{1, 2, 3})[:]))
	}
	if len(asFloat64s([]uint8{1, 2, 3})[:]) != 3 {
		t.Fatalf("Expected 3 elements for uint8 slice, got %d", len(asFloat64s([]uint8{1, 2, 3})[:]))
	}
	if len(asFloat64s([]uint16{1, 2, 3})[:]) != 3 {
		t.Fatalf("Expected 3 elements for uint16 slice, got %d", len(asFloat64s([]uint16{1, 2, 3})[:]))
	}
	if len(asFloat64s([]uint32{1, 2, 3})[:]) != 3 {
		t.Fatalf("Expected 3 elements for uint32 slice, got %d", len(asFloat64s([]uint32{1, 2, 3})[:]))
	}
	if len(asFloat64s([]uint64{1, 2, 3})[:]) != 3 {
		t.Fatalf("Expected 3 elements for uint64 slice, got %d", len(asFloat64s([]uint64{1, 2, 3})[:]))
	}
	if len(asFloat64s([]float32{1.1, 2.2, 3.3})[:]) != 3 {
		t.Fatalf("Expected 3 elements for float32 slice, got %d", len(asFloat64s([]float32{1.1, 2.2, 3.3})[:]))
	}
	if len(asFloat64s([]float64{1.1, 2.2, 3.3})[:]) != 3 {
		t.Fatalf("Expected 3 elements for float64 slice, got %d", len(asFloat64s([]float64{1.1, 2.2, 3.3})[:]))
	}
	if len(asFloat64s([]string{"1.1", "2.2", "3.3"})[:]) != 3 {
		t.Fatalf("Expected 3 elements for string slice, got %d", len(asFloat64s([]string{"1.1", "2.2", "3.3"})[:]))
	}
}

func TestAsZeroAllTypes(t *testing.T) {
	if !asZero(int(0)) {
		t.Fatal("Expected true for 0 int")
	}
	if !asZero(int8(0)) {
		t.Fatal("Expected true for 0 int8")
	}
	if !asZero(int16(0)) {
		t.Fatal("Expected true for 0 int16")
	}
	if !asZero(int32(0)) {
		t.Fatal("Expected true for 0 int32")
	}
	if !asZero(int64(0)) {
		t.Fatal("Expected true for 0 int64")
	}
	if !asZero(uint(0)) {
		t.Fatal("Expected true for 0 uint")
	}
	if !asZero(uint8(0)) {
		t.Fatal("Expected true for 0 uint8")
	}
	if !asZero(uint16(0)) {
		t.Fatal("Expected true for 0 uint16")
	}
	if !asZero(uint32(0)) {
		t.Fatal("Expected true for 0 uint32")
	}
	if !asZero(uint64(0)) {
		t.Fatal("Expected true for 0 uint64")
	}
	if asZero(int(1)) {
		t.Fatal("Expected false for 1")
	}
	if asZero("not a number") {
		t.Fatal("Expected false for non-number")
	}
}

func TestGetErrorMessage(t *testing.T) {
	if getErrorMessage("default", "") != "default" {
		t.Fatal("Expected default message")
	}
	if getErrorMessage("default", "custom") != "custom" {
		t.Fatal("Expected custom message")
	}
}
