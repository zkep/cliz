package validator

import (
	"testing"
)

func TestRange(t *testing.T) {
	v := Range(1, 10)
	err := v.Validate(5)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestRangeWithMessage(t *testing.T) {
	rv := &RangeValidator{
		Min:          1,
		Max:          10,
		ErrorMessage: "value must be between {0} and {1}",
	}
	err := rv.Validate(15)
	if err == nil {
		t.Fatal("Expected error for value outside range")
	}
	if err.Error() != ": value must be between {0} and {1}" {
		t.Fatalf("Expected error message ': value must be between {0} and {1}', got '%s'", err.Error())
	}
}

func TestRangeIntSlice(t *testing.T) {
	v := Range(1, 10)
	err := v.Validate([]int{5, 7, 3})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestRangeFloatSlice(t *testing.T) {
	v := Range(1.5, 5.5)
	err := v.Validate([]float64{2.0, 3.5, 5.0})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestRangeLowerBound(t *testing.T) {
	v := Range(1, 10)
	err := v.Validate(1)
	if err != nil {
		t.Fatalf("Expected no error for lower bound, got %v", err)
	}
	err = v.Validate(0)
	if err == nil {
		t.Fatal("Expected error for value below lower bound")
	}
}

func TestRangeUpperBound(t *testing.T) {
	v := Range(1, 10)
	err := v.Validate(10)
	if err != nil {
		t.Fatalf("Expected no error for upper bound, got %v", err)
	}
	err = v.Validate(11)
	if err == nil {
		t.Fatal("Expected error for value above upper bound")
	}
}

func TestRangeSliceEmpty(t *testing.T) {
	v := Range(1, 10)
	err := v.Validate([]int{})
	if err == nil {
		t.Fatal("Expected error for empty slice")
	}
	err = v.Validate([]float64{})
	if err == nil {
		t.Fatal("Expected error for empty float slice")
	}
	err = v.Validate([]string{})
	if err == nil {
		t.Fatal("Expected error for empty string slice")
	}
}
