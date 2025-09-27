package constants

// Error codes for API responses
const (
	// General error codes
	ErrCodeInvalidJSON     = "INVALID_JSON"
	ErrCodeInternalError   = "INTERNAL_ERROR"
	ErrCodeBadRequest      = "BAD_REQUEST"
	ErrCodeUnauthorized    = "UNAUTHORIZED"
	ErrCodeForbidden       = "FORBIDDEN"
	ErrCodeNotFound        = "NOT_FOUND"
	ErrCodeConflict        = "CONFLICT"
	ErrCodeTooManyRequests = "TOO_MANY_REQUESTS"
)

// User-specific error codes
const (
	ErrCodeUserNotFound   = "USER_NOT_FOUND"
	ErrCodeDuplicateEmail = "DUPLICATE_EMAIL"
	ErrCodeInvalidName    = "INVALID_NAME"
	ErrCodeInvalidEmail   = "INVALID_EMAIL"
	ErrCodeMissingUserID  = "MISSING_USER_ID"
	ErrCodeWeakPassword   = "WEAK_PASSWORD"
	ErrCodeEmailExists    = "EMAIL_EXISTS"
)

// Validation error codes
const (
	ErrCodeRequired      = "REQUIRED"
	ErrCodeMinLength     = "MIN_LENGTH"
	ErrCodeMaxLength     = "MAX_LENGTH"
	ErrCodeMinValue      = "MIN_VALUE"
	ErrCodeMaxValue      = "MAX_VALUE"
	ErrCodeInvalidFormat = "INVALID_FORMAT"
)
