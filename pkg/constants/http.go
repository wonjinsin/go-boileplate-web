package constants

// Context keys (use unexported type for safety)
type ContextKey string

const (
	// ContextKeyTrID is the key for storing TrID in context
	ContextKeyTrID ContextKey = "tr_id"
)

// HTTP Status related constants
const (
	// Default pagination values
	DefaultLimit  = 50
	MaxLimit      = 200
	DefaultOffset = 0
)

// HTTP Headers
const (
	HeaderContentType   = "Content-Type"
	HeaderAuthorization = "Authorization"
	HeaderAccept        = "Accept"
)

// Content Types
const (
	ContentTypeJSONCharset = "application/json; charset=utf-8"
)
