package validator

import (
	"regexp"
	"testing"
)

func TestPattern(t *testing.T) {
	v := Pattern(`^\d{3}-\d{3}-\d{4}$`)
	err := v.Validate("123-456-7890")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestPatternInvalid(t *testing.T) {
	v := Pattern(`^\d{3}-\d{3}-\d{4}$`)
	err := v.Validate("1234567890")
	if err == nil {
		t.Fatal("Expected error for invalid pattern")
	}
}

func TestPatternStructValidation(t *testing.T) {
	v := &PatternValidator{
		FieldName:    "phone",
		Pattern:      `^\d{3}-\d{3}-\d{4}$`,
		Regexp:       regexp.MustCompile(`^\d{3}-\d{3}-\d{4}$`),
		ErrorMessage: "phone must match pattern '%v'",
	}
	if err := v.Validate("123-456-7890"); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if err := v.Validate("1234567890"); err == nil {
		t.Fatal("Expected error for invalid pattern")
	}
}
