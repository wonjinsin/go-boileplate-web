package middleware

import (
	"log"
	"net/http"
	"runtime/debug"

	"example/pkg/utils"
)

// RecoveryConfig holds recovery configuration
type RecoveryConfig struct {
	EnableStackTrace bool
	LogStackTrace    bool
}

// DefaultRecoveryConfig returns default recovery configuration
func DefaultRecoveryConfig() RecoveryConfig {
	return RecoveryConfig{
		EnableStackTrace: true,
		LogStackTrace:    true,
	}
}

// Recovery returns a middleware that recovers from panics
func Recovery(config ...RecoveryConfig) func(http.Handler) http.Handler {
	cfg := DefaultRecoveryConfig()
	if len(config) > 0 {
		cfg = config[0]
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					// Log the panic
					log.Printf("PANIC: %v", err)

					if cfg.LogStackTrace && cfg.EnableStackTrace {
						log.Printf("Stack trace:\n%s", debug.Stack())
					}

					// Return 500 error
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusInternalServerError)

					errorResponse := map[string]interface{}{
						"error": "Internal server error",
						"code":  "INTERNAL_SERVER_ERROR",
					}

					if cfg.EnableStackTrace {
						errorResponse["stack_trace"] = string(debug.Stack())
					}

					utils.WriteJSON(w, http.StatusInternalServerError, errorResponse)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
