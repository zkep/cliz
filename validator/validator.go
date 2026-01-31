package validator

import (
	"fmt"
	"regexp"
)

// Validator interface for flag validation
// Implement this interface to create custom validators that can be used with flag validation
type Validator interface {
	Validate(value any) error
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

// countPlaceholders counts the number of placeholders in a format string
func countPlaceholders(format string) int {
	re := regexp.MustCompile(`%[0-9]*\.?[0-9]*[vTbcdefgopqstvxX]`)
	matches := re.FindAllString(format, -1)
	return len(matches)
}

func (e *ValidatorError) Error() string {
	msg := e.Message
	if e.MessageTemplate != "" {
		msg = e.MessageTemplate
		if len(e.Args) > 0 {
			if countPlaceholders(msg) == len(e.Args) {
				msg = fmt.Sprintf(msg, e.Args...)
			}
		}
	}
	return fmt.Sprintf("%s: %s", e.Field, msg)
}

// ValidateTags parses validate tags and creates corresponding validators
// This function is exported for use by the main cliz package
func ValidateTags(validateTags, fieldName string) []Validator {
	return parseValidateTags(validateTags, fieldName)
}
