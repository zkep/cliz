package validator

import (
	"testing"
)

func TestEq(t *testing.T) {
	v := Eq(42)
	err := v.Validate(42)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestEqNotEqual(t *testing.T) {
	v := Eq(42)
	err := v.Validate(10)
	if err == nil {
		t.Fatal("Expected error for not equal value")
	}
}

func TestEqFloat(t *testing.T) {
	v := Eq(3.14)
	err := v.Validate(3.14)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestEqVariousTypes(t *testing.T) {
	v := Eq(10)
	if err := v.Validate(10); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := v.Validate(10.0); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := v.Validate(11); err == nil {
		t.Fatal("Expected error for not equal value")
	}
}

func TestEqAllTypes(t *testing.T) {
	v := Eq(10)
	if err := v.Validate(10); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := v.Validate(10.0); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := v.Validate(10.5); err == nil {
		t.Fatal("Expected error for not equal value")
	}
}

func TestEqString(t *testing.T) {
	v := Eq("hello")
	if err := v.Validate("hello"); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := v.Validate("world"); err == nil {
		t.Fatal("Expected error for not equal string")
	}
}

func TestEqBoolean(t *testing.T) {
	v := Eq(true)
	if err := v.Validate(true); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := v.Validate(false); err == nil {
		t.Fatal("Expected error for not equal boolean")
	}
}

func TestEqSlice(t *testing.T) {
	v := Eq("value")
	if err := v.Validate([]string{"value", "value", "value"}); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := v.Validate([]string{"value", "wrong", "value"}); err == nil {
		t.Fatal("Expected error for slice with wrong value")
	}
}

func TestEqNil(t *testing.T) {
	v := Eq(nil)
	if err := v.Validate(nil); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestEqEmptyStringSlice(t *testing.T) {
	v := Eq("value")
	if err := v.Validate([]string{}); err == nil {
		t.Fatal("Expected error for empty string slice")
	}
}

func TestEqEmptyBoolSlice(t *testing.T) {
	v := Eq(true)
	if err := v.Validate([]bool{}); err == nil {
		t.Fatal("Expected error for empty bool slice")
	}
}

func TestEqEmptyFloatSlice(t *testing.T) {
	v := Eq(10.5)
	if err := v.Validate([]float64{}); err == nil {
		t.Fatal("Expected error for empty float slice")
	}
}

func TestEqIntSlice(t *testing.T) {
	v := Eq(5)
	if err := v.Validate([]int{5, 5, 5}); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := v.Validate([]int{5, 6, 5}); err == nil {
		t.Fatal("Expected error for int slice with wrong value")
	}
}

func TestEqFloatSlice(t *testing.T) {
	v := Eq(3.14)
	if err := v.Validate([]float64{3.14, 3.14}); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := v.Validate([]float64{3.14, 3.15}); err == nil {
		t.Fatal("Expected error for float slice with wrong value")
	}
}

func TestEqBoolSlice(t *testing.T) {
	v := Eq(true)
	if err := v.Validate([]bool{true, true, true}); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := v.Validate([]bool{true, false, true}); err == nil {
		t.Fatal("Expected error for bool slice with wrong value")
	}
}

func TestEqWithCustomErrorMessage(t *testing.T) {
	v := &EqValidator{
		FieldName:    "age",
		Value:        18,
		ErrorMessage: "age must be exactly %v",
	}
	if err := v.Validate(18); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := v.Validate(20); err == nil {
		t.Fatal("Expected error for not equal value")
	}
}

func TestEqDifferentTypes(t *testing.T) {
	v := Eq("10")
	if err := v.Validate(10); err != nil {
		t.Fatalf("Unexpected error when comparing string with int: %v", err)
	}
}

func TestEqStructValidation(t *testing.T) {
	v := &EqValidator{
		FieldName:    "status",
		Value:        "active",
		ErrorMessage: "status must be %v",
	}
	if err := v.Validate("active"); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := v.Validate("inactive"); err == nil {
		t.Fatal("Expected error for not equal value")
	}
}

func TestEqUint(t *testing.T) {
	v := Eq(uint(10))
	if err := v.Validate(uint(10)); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := v.Validate(uint(20)); err == nil {
		t.Fatal("Expected error for not equal uint")
	}
}

func TestEqInt8(t *testing.T) {
	v := Eq(int8(10))
	if err := v.Validate(int8(10)); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestEqUint64(t *testing.T) {
	v := Eq(uint64(100000000000))
	if err := v.Validate(uint64(100000000000)); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestEqMixedNumericTypes(t *testing.T) {
	v := Eq(int64(10))
	if err := v.Validate(float64(10.0)); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := v.Validate(float64(10.5)); err == nil {
		t.Fatal("Expected error for not equal float")
	}
}

func TestEqStringSliceAllValues(t *testing.T) {
	v := Eq("test")
	tests := []struct {
		name  string
		value []string
		want  bool
	}{
		{"all equal", []string{"test", "test"}, true},
		{"one different", []string{"test", "wrong"}, false},
		{"empty slice", []string{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v.Validate(tt.value)
			if tt.want {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
			} else {
				if err == nil {
					t.Fatal("Expected error")
				}
			}
		})
	}
}

func TestEqMultipleTypes(t *testing.T) {
	tests := []struct {
		name  string
		value any
		other any
		want  bool
	}{
		{"int equals int", 10, 10, true},
		{"int not equals int", 10, 20, false},
		{"string equals string", "hello", "hello", true},
		{"string not equals string", "hello", "world", false},
		{"bool equals bool", true, true, true},
		{"bool not equals bool", true, false, false},
		{"int equals float", 10, 10.0, true},
		{"float not equals int", 10.5, 10, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Eq(tt.other)
			err := v.Validate(tt.value)
			if tt.want {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
			} else {
				if err == nil {
					t.Fatal("Expected error")
				}
			}
		})
	}
}

func TestEqBooleanType(t *testing.T) {
	validators := ValidateTags("eq=true", "field")
	if len(validators) != 1 {
		t.Fatalf("Expected 1 validator, got %d", len(validators))
	}
	err := validators[0].Validate(true)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	err = validators[0].Validate(false)
	if err == nil {
		t.Fatal("Expected error for not equal boolean")
	}
}

func TestEqFloatType(t *testing.T) {
	validators := ValidateTags("eq=3.14", "field")
	if len(validators) != 1 {
		t.Fatalf("Expected 1 validator, got %d", len(validators))
	}
	err := validators[0].Validate(3.14)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	err = validators[0].Validate(3.15)
	if err == nil {
		t.Fatal("Expected error for not equal float")
	}
}

func TestEqStringType(t *testing.T) {
	validators := ValidateTags("eq=test", "field")
	if len(validators) != 1 {
		t.Fatalf("Expected 1 validator, got %d", len(validators))
	}
	err := validators[0].Validate("test")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	err = validators[0].Validate("testing")
	if err == nil {
		t.Fatal("Expected error for not equal string")
	}
}

func TestEqWithErrorMessages(t *testing.T) {
	validators := ValidateTags("eq=10,error_eq=Value must be 10", "field")
	if len(validators) != 1 {
		t.Fatalf("Expected 1 validator, got %d", len(validators))
	}
	err := validators[0].Validate(5)
	if err == nil {
		t.Fatal("Expected error for not equal value")
	}
}
