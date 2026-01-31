package validator

import (
	"testing"
)

func TestValidateTagsRange(t *testing.T) {
	validators := ValidateTags("range=1-10", "age")
	if len(validators) != 1 {
		t.Fatalf("Expected 1 validator, got %d", len(validators))
	}
	err := validators[0].Validate(5)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	err = validators[0].Validate(15)
	if err == nil {
		t.Fatal("Expected error for value outside range")
	}
}

func TestValidateTagsEq(t *testing.T) {
	validators := ValidateTags("eq=18", "age")
	if len(validators) != 1 {
		t.Fatalf("Expected 1 validator, got %d", len(validators))
	}
	err := validators[0].Validate(18)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	err = validators[0].Validate(20)
	if err == nil {
		t.Fatal("Expected error for not equal value")
	}
}

func TestValidateTagsLen(t *testing.T) {
	validators := ValidateTags("len=5", "code")
	if len(validators) != 1 {
		t.Fatalf("Expected 1 validator, got %d", len(validators))
	}
	err := validators[0].Validate("abcde")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	err = validators[0].Validate("abcd")
	if err == nil {
		t.Fatal("Expected error for short string")
	}
}

func TestValidateTagsContains(t *testing.T) {
	validators := ValidateTags("contains=hello", "message")
	if len(validators) != 1 {
		t.Fatalf("Expected 1 validator, got %d", len(validators))
	}
	err := validators[0].Validate("hello world")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	err = validators[0].Validate("world")
	if err == nil {
		t.Fatal("Expected error for missing substring")
	}
}

func TestValidateTagsAlpha(t *testing.T) {
	validators := ValidateTags("alpha", "name")
	if len(validators) != 1 {
		t.Fatalf("Expected 1 validator, got %d", len(validators))
	}
	err := validators[0].Validate("hello")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	err = validators[0].Validate("hello123")
	if err == nil {
		t.Fatal("Expected error for non-alpha string")
	}
}

func TestValidateTagsAlphanum(t *testing.T) {
	validators := ValidateTags("alphanum", "username")
	if len(validators) != 1 {
		t.Fatalf("Expected 1 validator, got %d", len(validators))
	}
	err := validators[0].Validate("abc123")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	err = validators[0].Validate("abc!@#")
	if err == nil {
		t.Fatal("Expected error for symbols")
	}
}

func TestValidateTagsPattern(t *testing.T) {
	validators := ValidateTags(`pattern=^\d{3}-\d{3}-\d{4}$`, "phone")
	if len(validators) != 1 {
		t.Fatalf("Expected 1 validator, got %d", len(validators))
	}
	err := validators[0].Validate("123-456-7890")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	err = validators[0].Validate("1234567890")
	if err == nil {
		t.Fatal("Expected error for invalid pattern")
	}
}

func TestValidateTagsURL(t *testing.T) {
	validators := ValidateTags("url", "website")
	if len(validators) != 1 {
		t.Fatalf("Expected 1 validator, got %d", len(validators))
	}
	err := validators[0].Validate("http://example.com")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	err = validators[0].Validate("invalid-url")
	if err == nil {
		t.Fatal("Expected error for invalid URL")
	}
}

func TestValidateTagsMultiple(t *testing.T) {
	validators := ValidateTags("required,email,gt=18", "user")
	if len(validators) != 3 {
		t.Fatalf("Expected 3 validators, got %d", len(validators))
	}
}

func TestParseValidateTagsWithErrorMessages(t *testing.T) {
	validators := ValidateTags("required,eq=10,error_required=Field is required,error_eq=Value must be 10", "field")
	if len(validators) != 2 {
		t.Fatalf("Expected 2 validators, got %d", len(validators))
	}
}

func TestParseValidateTagsComplex(t *testing.T) {
	validators := ValidateTags("required,range=1-10,gt=0,lt=100,len=5", "field")
	if len(validators) != 5 {
		t.Fatalf("Expected 5 validators, got %d", len(validators))
	}
}

func TestParseValidateTagsWithQuotes(t *testing.T) {
	validators := ValidateTags(`in="dev|staging|prod"`, "env")
	if len(validators) != 1 {
		t.Fatalf("Expected 1 validator, got %d", len(validators))
	}
	err := validators[0].Validate("dev")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	err = validators[0].Validate("test")
	if err == nil {
		t.Fatal("Expected error for invalid value")
	}
}

func TestParseValidateTagsInvalidPattern(t *testing.T) {
	validators := ValidateTags(`pattern=([invalid regex]`, "field")
	if len(validators) != 0 {
		t.Fatalf("Expected 0 validators for invalid regex, got %d", len(validators))
	}
}

func TestParseValidateTagsEmpty(t *testing.T) {
	validators := ValidateTags("", "field")
	if len(validators) != 0 {
		t.Fatalf("Expected 0 validators, got %d", len(validators))
	}
}

func TestParseValidateTagsWhitespace(t *testing.T) {
	validators := ValidateTags(" required , eq=10 ", "field")
	if len(validators) != 2 {
		t.Fatalf("Expected 2 validators, got %d", len(validators))
	}
}
