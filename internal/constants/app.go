package constants

// Application configuration
const (
	AppName    = "go-boiler-plate"
	AppVersion = "1.0.0"
)

// Database configuration
const (
	DefaultDBTimeout       = 30 // seconds
	DefaultMaxOpenConns    = 25
	DefaultMaxIdleConns    = 5
	DefaultConnMaxLifetime = 300 // seconds
)

// Cache configuration
const (
	DefaultCacheTTL = 3600 // seconds (1 hour)
	CacheKeyPrefix  = "app:"
)

// API configuration
const (
	APIVersion = "v1"
	APIPrefix  = "/api/" + APIVersion
)

// Default routes
const (
	RouteHealth = "/healthz"
	RouteUsers  = "/users"
	RouteUser   = "/users/"
)

// Context keys (use unexported type for safety)
type contextKey string

const (
	ContextKeyUserID    contextKey = "user_id"
	ContextKeyRequestID contextKey = "request_id"
	ContextKeyTraceID   contextKey = "trace_id"
)
