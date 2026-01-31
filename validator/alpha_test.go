package validator

import (
	"testing"
)

func TestAlpha(t *testing.T) {
	v := Alpha()
	err := v.Validate("hello")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestAlphaWithNumbers(t *testing.T) {
	v := Alpha()
	err := v.Validate("hello123")
	if err == nil {
		t.Fatal("Expected error for string with numbers")
	}
}
