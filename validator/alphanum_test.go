package validator

import (
	"testing"
)

func TestAlphanum(t *testing.T) {
	v := Alphanum()
	err := v.Validate("abc123")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestAlphanumWithSymbols(t *testing.T) {
	v := Alphanum()
	err := v.Validate("abc123!")
	if err == nil {
		t.Fatal("Expected error for symbols")
	}
}
