package validator

import (
	"testing"
)

func TestIn(t *testing.T) {
	v := In("dev", "staging", "prod")
	err := v.Validate("dev")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestInMissing(t *testing.T) {
	v := In("dev", "staging", "prod")
	err := v.Validate("testing")
	if err == nil {
		t.Fatal("Expected error for missing value")
	}
}

func TestInEmptyAllowedList(t *testing.T) {
	v := In()
	err := v.Validate("anything")
	if err == nil {
		t.Fatal("Expected error for empty allowed list")
	}
	err = v.Validate([]string{"something"})
	if err == nil {
		t.Fatal("Expected error for empty allowed list with slice")
	}
}
