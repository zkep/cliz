package validator

import (
	"testing"
)

func TestRequired(t *testing.T) {
	v := Required()
	err := v.Validate("test")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestRequiredEmptyString(t *testing.T) {
	v := Required()
	err := v.Validate("")
	if err == nil {
		t.Fatal("Expected error for empty string")
	}
}

func TestRequiredEmptySlice(t *testing.T) {
	v := Required()
	err := v.Validate([]int{})
	if err == nil {
		t.Fatal("Expected error for empty slice")
	}
}

func TestRequiredNilSlice(t *testing.T) {
	v := Required()
	err := v.Validate([]string(nil))
	if err == nil {
		t.Fatal("Expected error for nil slice")
	}
}

func TestRequiredBoolean(t *testing.T) {
	v := Required()
	if err := v.Validate(true); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := v.Validate(false); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}
