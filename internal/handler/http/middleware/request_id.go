package middleware

import (
	"context"
	"net/http"

	"github.com/wonjinsin/go-boilerplate/internal/constants"
	"github.com/wonjinsin/go-boilerplate/pkg/utils"
)

// RequestIDConfig holds request ID configuration
type RequestIDConfig struct {
	Header    string
	Generator func() string
}

// DefaultRequestIDConfig returns default request ID configuration
func DefaultRequestIDConfig() RequestIDConfig {
	return RequestIDConfig{
		Header: "X-Request-ID",
		Generator: func() string {
			id, _ := utils.GenerateRandomID(16)
			return id
		},
	}
}

// RequestID returns a middleware that generates or extracts request IDs
func RequestID(config ...RequestIDConfig) func(http.Handler) http.Handler {
	cfg := DefaultRequestIDConfig()
	if len(config) > 0 {
		cfg = config[0]
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Try to get request ID from header
			requestID := r.Header.Get(cfg.Header)

			// Generate new ID if not present
			if requestID == "" {
				requestID = cfg.Generator()
			}

			// Set response header
			w.Header().Set(cfg.Header, requestID)

			// Add to context
			ctx := context.WithValue(r.Context(), constants.ContextKeyRequestID, requestID)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

// GetRequestID extracts request ID from context
func GetRequestID(ctx context.Context) string {
	if requestID, ok := ctx.Value(constants.ContextKeyRequestID).(string); ok {
		return requestID
	}
	return ""
}
