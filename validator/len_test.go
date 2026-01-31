package validator

import (
	"testing"
)

func TestLen(t *testing.T) {
	v := Len(5)
	err := v.Validate("hello")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestLenStringShort(t *testing.T) {
	v := Len(5)
	err := v.Validate("hi")
	if err == nil {
		t.Fatal("Expected error for short string")
	}
}

func TestLenInt(t *testing.T) {
	v := Len(3)
	err := v.Validate(123)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestLenSlice(t *testing.T) {
	v := Len(5)
	if err := v.Validate("hello"); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := v.Validate("hi"); err == nil {
		t.Fatal("Expected error for short string")
	}
}

func TestLenStructValidation(t *testing.T) {
	v := &LenValidator{
		FieldName:    "code",
		Length:       5,
		ErrorMessage: "code must be %v characters long",
	}
	if err := v.Validate("abcde"); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := v.Validate("abcd"); err == nil {
		t.Fatal("Expected error for short string")
	}
}

func TestLenStringSlice(t *testing.T) {
	v := Len(5)
	if err := v.Validate([]string{"hello", "world", "apple"}); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := v.Validate([]string{"hello", "hi", "world"}); err == nil {
		t.Fatal("Expected error for slice with short string")
	}
}

func TestLenIntSlice(t *testing.T) {
	v := Len(3)
	if err := v.Validate([]int{123, 456, 789}); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := v.Validate([]int{123, 12, 456}); err == nil {
		t.Fatal("Expected error for slice with wrong length int")
	}
}

func TestLenFloatSlice(t *testing.T) {
	v := Len(9)
	if err := v.Validate([]float64{12.3456}); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}
