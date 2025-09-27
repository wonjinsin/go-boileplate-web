package shared

import (
	"fmt"
	"regexp"
	"strings"

	"example/internal/constants"
	pkgConstants "example/pkg/constants"
)

var (
	emailRegex = regexp.MustCompile(pkgConstants.EmailPattern)
)

// Validator provides validation utilities for the application
type Validator struct {
	errors ValidationErrors
}

// NewValidator creates a new validator instance
func NewValidator() *Validator {
	return &Validator{
		errors: make(ValidationErrors, 0),
	}
}

// ValidateRequired validates that a field is not empty
func (v *Validator) ValidateRequired(field, value, fieldName string) *Validator {
	if strings.TrimSpace(value) == "" {
		v.errors = append(v.errors, NewValidationError(field, fieldName+" is required", constants.ErrCodeRequired))
	}
	return v
}

// ValidateEmail validates email format
func (v *Validator) ValidateEmail(field, email string) *Validator {
	email = strings.TrimSpace(email)
	if email != "" && !emailRegex.MatchString(email) {
		v.errors = append(v.errors, NewValidationError(field, "invalid email format", constants.ErrCodeInvalidEmail))
	}
	return v
}

// ValidateLength validates string length
func (v *Validator) ValidateLength(field, value string, min, max int, fieldName string) *Validator {
	length := len(strings.TrimSpace(value))
	if length < min {
		v.errors = append(v.errors, NewValidationError(field,
			fmt.Sprintf("%s must be at least %d characters", fieldName, min), constants.ErrCodeMinLength))
	}
	if max > 0 && length > max {
		v.errors = append(v.errors, NewValidationError(field,
			fmt.Sprintf("%s must not exceed %d characters", fieldName, max), constants.ErrCodeMaxLength))
	}
	return v
}

// ValidateRange validates numeric range
func (v *Validator) ValidateRange(field string, value, min, max int, fieldName string) *Validator {
	if value < min {
		v.errors = append(v.errors, NewValidationError(field,
			fmt.Sprintf("%s must be at least %d", fieldName, min), constants.ErrCodeMinValue))
	}
	if max > 0 && value > max {
		v.errors = append(v.errors, NewValidationError(field,
			fmt.Sprintf("%s must not exceed %d", fieldName, max), constants.ErrCodeMaxValue))
	}
	return v
}

// HasErrors returns true if there are validation errors
func (v *Validator) HasErrors() bool {
	return len(v.errors) > 0
}

// Errors returns all validation errors
func (v *Validator) Errors() ValidationErrors {
	return v.errors
}

// Error returns validation errors as a single error
func (v *Validator) Error() error {
	if len(v.errors) == 0 {
		return nil
	}
	return v.errors
}
