package validator

import (
	"testing"
)

func TestGt(t *testing.T) {
	v := Gt(18)
	err := v.Validate(20)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGtEqual(t *testing.T) {
	v := Gt(18)
	err := v.Validate(18)
	if err == nil {
		t.Fatal("Expected error for equal value")
	}
}

func TestGtFloat(t *testing.T) {
	v := Gt(18.5)
	err := v.Validate(19.0)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGtAllTypes(t *testing.T) {
	v := Gt(18.5)
	if err := v.Validate(19.0); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := v.Validate(18); err == nil {
		t.Fatal("Expected error for equal value")
	}
}

func TestGtSlice(t *testing.T) {
	v := Gt(18)
	if err := v.Validate([]int{19, 20, 21}); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := v.Validate([]int{19, 18, 20}); err == nil {
		t.Fatal("Expected error for slice with wrong value")
	}
}
