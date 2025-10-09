package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggerConfig holds logger configuration
type LoggerConfig struct {
	TimeFormat string
	UTC        bool
	SkipPaths  []string
}

// DefaultLoggerConfig returns default logger configuration
func DefaultLoggerConfig() LoggerConfig {
	return LoggerConfig{
		TimeFormat: time.RFC3339,
		UTC:        true,
		SkipPaths:  []string{"/healthz", "/metrics"},
	}
}

// responseWriter wraps http.ResponseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	size       int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.size += size
	return size, err
}

// Logger returns a middleware that logs HTTP requests
func Logger(config ...LoggerConfig) func(http.Handler) http.Handler {
	cfg := DefaultLoggerConfig()
	if len(config) > 0 {
		cfg = config[0]
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip logging for specified paths
			for _, path := range cfg.SkipPaths {
				if r.URL.Path == path {
					next.ServeHTTP(w, r)
					return
				}
			}

			start := time.Now()
			if cfg.UTC {
				start = start.UTC()
			}

			// Wrap response writer
			wrapped := newResponseWriter(w)

			// Process request
			next.ServeHTTP(wrapped, r)

			// Calculate duration
			duration := time.Since(start)

			// Log request
			log.Printf(
				"%s - %s %s %s - %d %d - %v - %s",
				start.Format(cfg.TimeFormat),
				r.RemoteAddr,
				r.Method,
				r.URL.Path,
				wrapped.statusCode,
				wrapped.size,
				duration,
				r.UserAgent(),
			)
		})
	}
}
