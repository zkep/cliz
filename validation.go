package cliz

import (
	"fmt"

	"github.com/zkep/cliz/validator"
)

// Validator interface for flag validation
// Implement this interface to create custom validators that can be used with flag validation
type Validator interface {
	Validate(value any) error
	WithMessage(msg string) Validator
}

// ValidatorFunc is a function type that implements the Validator interface
type ValidatorFunc func(value any) error

// Validate calls the underlying function
func (f ValidatorFunc) Validate(value any) error {
	return f(value)
}

// ValidatorError represents a validation error with field name and error message
type ValidatorError struct {
	Field           string
	Message         string // Backward compatibility with existing error messages
	MessageTemplate string // Template for custom error messages with placeholders
	Args            []any  // Arguments to populate the message template
}

func (e *ValidatorError) Error() string {
	// Use custom message template with arguments if provided
	if e.MessageTemplate != "" {
		msg := e.MessageTemplate
		if len(e.Args) > 0 {
			msg = fmt.Sprintf(msg, e.Args...)
		}
		return fmt.Sprintf("%s: %s", e.Field, msg)
	}
	// Fallback to original message for backward compatibility
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// parseValidateTags parses validate tags and creates corresponding validators
func parseValidateTags(validateTags, fieldName string) []Validator {
	// Delegate to validation package
	externalValidators := validator.ValidateTags(validateTags, fieldName)
	var validators []Validator
	for _, v := range externalValidators {
		validators = append(validators, validatorWrapper{v})
	}
	return validators
}

// validatorWrapper adapts external validation.Validator to cliz.Validator
type validatorWrapper struct {
	externalValidator validator.Validator
}

func (w validatorWrapper) Validate(value any) error {
	return w.externalValidator.Validate(value)
}

func (w validatorWrapper) WithMessage(msg string) Validator {
	switch v := w.externalValidator.(type) {
	case *validator.RangeValidator:
		v.ErrorMessage = msg
		return w
	case *validator.RequiredValidator:
		v.ErrorMessage = msg
		return w
	case *validator.LenValidator:
		v.ErrorMessage = msg
		return w
	case *validator.PatternValidator:
		v.ErrorMessage = msg
		return w
	case *validator.InValidator:
		v.ErrorMessage = msg
		return w
	case *validator.EqValidator:
		v.ErrorMessage = msg
		return w
	case *validator.GtValidator:
		v.ErrorMessage = msg
		return w
	case *validator.LtValidator:
		v.ErrorMessage = msg
		return w
	case *validator.ContainsValidator:
		v.ErrorMessage = msg
		return w
	case *validator.EmailValidator:
		v.ErrorMessage = msg
		return w
	case *validator.URLValidator:
		v.ErrorMessage = msg
		return w
	case *validator.AlphaValidator:
		v.ErrorMessage = msg
		return w
	case *validator.AlphanumValidator:
		v.ErrorMessage = msg
		return w
	default:
		return w
	}
}

// Range creates a validator that checks if a value is within the specified range
func Range(minValue, maxValue float64) Validator {
	return validatorWrapper{validator.Range(minValue, maxValue)}
}

// Required creates a validator that checks if a value is provided
func Required() Validator {
	return validatorWrapper{validator.Required()}
}

// Eq creates a validator that checks if a value equals the specified value
func Eq(value any) Validator {
	return validatorWrapper{validator.Eq(value)}
}

// Len creates a validator that checks if a string length equals the specified length
func Len(length int) Validator {
	return validatorWrapper{validator.Len(length)}
}

// Gt creates a validator that checks if a value is greater than the specified value
func Gt(value float64) Validator {
	return validatorWrapper{validator.Gt(value)}
}

// Lt creates a validator that checks if a value is less than the specified value
func Lt(value float64) Validator {
	return validatorWrapper{validator.Lt(value)}
}

// In creates a validator that checks if a value is in the specified list of allowed values
func In(allowed ...string) Validator {
	return validatorWrapper{validator.In(allowed...)}
}

// Contains creates a validator that checks if a string contains the specified substring
func Contains(substring string) Validator {
	return validatorWrapper{validator.Contains(substring)}
}

// Alpha creates a validator that checks if a string contains only alphabetic characters
func Alpha() Validator {
	return validatorWrapper{validator.Alpha()}
}

// Alphanum creates a validator that checks if a string contains only alphanumeric characters
func Alphanum() Validator {
	return validatorWrapper{validator.Alphanum()}
}

// Email creates a validator that checks if a string is a valid email address
func Email() Validator {
	return validatorWrapper{validator.Email()}
}

// URL creates a validator that checks if a string is a valid URL
func URL() Validator {
	return validatorWrapper{validator.URL()}
}

// Pattern creates a validator that checks if a string matches the specified regex pattern
func Pattern(pattern string) Validator {
	return validatorWrapper{validator.Pattern(pattern)}
}

// Custom creates a custom validator using the provided function
func Custom(validateFunc ValidatorFunc) Validator {
	return validatorWrapper{validateFunc}
}
