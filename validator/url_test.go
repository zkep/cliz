package validator

import (
	"testing"
)

func TestURL(t *testing.T) {
	v := URL()
	err := v.Validate("http://example.com")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestURLInvalid(t *testing.T) {
	v := URL()
	err := v.Validate("invalid-url")
	if err == nil {
		t.Fatal("Expected error for invalid URL")
	}
}