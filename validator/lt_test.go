package validator

import (
	"testing"
)

func TestLt(t *testing.T) {
	v := Lt(18)
	err := v.Validate(10)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestLtEqual(t *testing.T) {
	v := Lt(18)
	err := v.Validate(18)
	if err == nil {
		t.Fatal("Expected error for equal value")
	}
}

func TestLtFloat(t *testing.T) {
	v := Lt(18.5)
	err := v.Validate(18.0)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestLtAllTypes(t *testing.T) {
	v := Lt(18.5)
	if err := v.Validate(18.0); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := v.Validate(18.5); err == nil {
		t.Fatal("Expected error for equal value")
	}
}

func TestLtStructValidation(t *testing.T) {
	v := &LtValidator{
		FieldName:    "age",
		Value:        18,
		ErrorMessage: "age must be less than %v",
	}
	if err := v.Validate(17); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := v.Validate(18); err == nil {
		t.Fatal("Expected error for equal value")
	}
}

func TestLtSlice(t *testing.T) {
	v := Lt(18)
	if err := v.Validate([]int{17, 16, 15}); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := v.Validate([]int{17, 18, 15}); err == nil {
		t.Fatal("Expected error for slice with wrong value")
	}
}
