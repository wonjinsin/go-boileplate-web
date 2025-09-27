package shared

import (
	"errors"
	"fmt"
)

// Common application errors
var (
	ErrInvalidInput   = errors.New("invalid input")
	ErrNotFound       = errors.New("not found")
	ErrAlreadyExists  = errors.New("already exists")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("forbidden")
	ErrInternalServer = errors.New("internal server error")
)

// ValidationError represents a validation error with field details
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}

// ValidationErrors represents multiple validation errors
type ValidationErrors []ValidationError

func (e ValidationErrors) Error() string {
	if len(e) == 0 {
		return "validation errors"
	}
	if len(e) == 1 {
		return e[0].Error()
	}
	return fmt.Sprintf("validation errors: %d fields failed", len(e))
}

// NewValidationError creates a new validation error
func NewValidationError(field, message, code string) ValidationError {
	return ValidationError{
		Field:   field,
		Message: message,
		Code:    code,
	}
}

// WrapError wraps an error with additional context
func WrapError(err error, message string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", message, err)
}

// IsValidationError checks if error is a validation error
func IsValidationError(err error) bool {
	var validationErr ValidationError
	var validationErrs ValidationErrors
	return errors.As(err, &validationErr) || errors.As(err, &validationErrs)
}
