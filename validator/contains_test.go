package validator

import (
	"testing"
)

func TestContains(t *testing.T) {
	v := Contains("hello")
	err := v.Validate("hello world")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestContainsMissing(t *testing.T) {
	v := Contains("hello")
	err := v.Validate("world")
	if err == nil {
		t.Fatal("Expected error for missing substring")
	}
}

func TestContainsEmptyString(t *testing.T) {
	if !containsEmptyString([]string{"", "not empty"}) {
		t.Fatal("Expected true for empty string")
	}
	if containsEmptyString([]string{"not empty", "also not empty"}) {
		t.Fatal("Expected false for non-empty strings")
	}
}

func TestContainsZeroInt64(t *testing.T) {
	if !containsZeroInt64([]int64{0, 1, 2}) {
		t.Fatal("Expected true for zero int64")
	}
	if containsZeroInt64([]int64{1, 2, 3}) {
		t.Fatal("Expected false for non-zero int64s")
	}
}

func TestContainsZeroFloat64(t *testing.T) {
	if !containsZeroFloat64([]float64{0.0, 1.0, 2.0}) {
		t.Fatal("Expected true for zero float64")
	}
	if containsZeroFloat64([]float64{1.0, 2.0, 3.0}) {
		t.Fatal("Expected false for non-zero float64s")
	}
}

func TestContainsInt64(t *testing.T) {
	if !containsInt64([]string{"1", "2", "3"}, 2) {
		t.Fatal("Expected true for int64 2")
	}
	if containsInt64([]string{"1", "2", "3"}, 4) {
		t.Fatal("Expected false for int64 4")
	}
}

func TestContainsFloat64(t *testing.T) {
	if !containsFloat64([]string{"1.0", "2.0", "3.0"}, 2.0) {
		t.Fatal("Expected true for float64 2.0")
	}
	if containsFloat64([]string{"1.0", "2.0", "3.0"}, 4.0) {
		t.Fatal("Expected false for float64 4.0")
	}
}