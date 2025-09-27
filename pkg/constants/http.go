package constants

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
	HeaderUserAgent     = "User-Agent"
	HeaderAccept        = "Accept"
	HeaderCacheControl  = "Cache-Control"
)

// Content Types
const (
	ContentTypeJSON        = "application/json"
	ContentTypeJSONCharset = "application/json; charset=utf-8"
	ContentTypeXML         = "application/xml"
	ContentTypeForm        = "application/x-www-form-urlencoded"
	ContentTypeMultipart   = "multipart/form-data"
)

// HTTP Methods (if needed beyond standard library)
const (
	MethodOptions = "OPTIONS"
	MethodHead    = "HEAD"
	MethodPatch   = "PATCH"
)
