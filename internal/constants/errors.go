package constants

// Error codes for API responses
const (
	// General error codes
	ErrCodeInvalidJSON     = "INVALID_JSON"
	ErrCodeInternalError   = "INTERNAL_ERROR"
	ErrCodeBadRequest      = "BAD_REQUEST"
	ErrCodeUnauthorized    = "UNAUTHORIZED"
	ErrCodeTooManyRequests = "TOO_MANY_REQUESTS"
)

// User-specific error codes
const (
	ErrCodeUserNotFound   = "USER_NOT_FOUND"
	ErrCodeDuplicateEmail = "DUPLICATE_EMAIL"
	ErrCodeInvalidName    = "INVALID_NAME"
	ErrCodeMissingUserID  = "MISSING_USER_ID"
)
