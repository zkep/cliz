package validator

import (
	"regexp"
)

// Range creates a validator that checks if a value is within the specified range
func Range(minValue, maxValue float64) Validator {
	return &RangeValidator{
		Min: minValue,
		Max: maxValue,
	}
}

// Required creates a validator that checks if a value is provided
func Required() Validator {
	return &RequiredValidator{}
}

// Eq creates a validator that checks if a value equals the specified value
func Eq(value any) Validator {
	return &EqValidator{
		Value: value,
	}
}

// Len creates a validator that checks if a string length equals the specified length
func Len(length int) Validator {
	return &LenValidator{
		Length: length,
	}
}

// Gt creates a validator that checks if a value is greater than the specified value
func Gt(value float64) Validator {
	return &GtValidator{
		Value: value,
	}
}

// Lt creates a validator that checks if a value is less than the specified value
func Lt(value float64) Validator {
	return &LtValidator{
		Value: value,
	}
}

// In creates a validator that checks if a value is in the specified list of allowed values
func In(allowed ...string) Validator {
	return &InValidator{
		Allowed: allowed,
	}
}

// Contains creates a validator that checks if a string contains the specified substring
func Contains(substring string) Validator {
	return &ContainsValidator{
		Substring: substring,
	}
}

// Alpha creates a validator that checks if a string contains only alphabetic characters
func Alpha() Validator {
	return &AlphaValidator{}
}

// Alphanum creates a validator that checks if a string contains only alphanumeric characters
func Alphanum() Validator {
	return &AlphanumValidator{}
}

// Email creates a validator that checks if a string is a valid email address
func Email() Validator {
	return &EmailValidator{}
}

// URL creates a validator that checks if a string is a valid URL
func URL() Validator {
	return &URLValidator{}
}

// Pattern creates a validator that checks if a string matches the specified regex pattern
func Pattern(pattern string) Validator {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return &PatternValidator{
			Pattern: pattern,
		}
	}
	return &PatternValidator{
		Pattern: pattern,
		Regexp:  r,
	}
}
