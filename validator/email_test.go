package validator

import (
	"testing"
)

func TestEmail(t *testing.T) {
	v := Email()
	err := v.Validate("test@example.com")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestEmailInvalid(t *testing.T) {
	v := Email()
	err := v.Validate("invalid-email")
	if err == nil {
		t.Fatal("Expected error for invalid email")
	}
}
