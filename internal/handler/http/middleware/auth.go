package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"example/internal/constants"
	"example/pkg/utils"
)

// AuthConfig holds authentication configuration
type AuthConfig struct {
	SkipPaths []string
	TokenType string // "Bearer", "Basic", etc.
}

// DefaultAuthConfig returns default auth configuration
func DefaultAuthConfig() AuthConfig {
	return AuthConfig{
		SkipPaths: []string{"/healthz", "/api/v1/auth/login", "/api/v1/auth/register"},
		TokenType: "Bearer",
	}
}

// Auth returns a middleware that handles authentication
func Auth(config ...AuthConfig) func(http.Handler) http.Handler {
	cfg := DefaultAuthConfig()
	if len(config) > 0 {
		cfg = config[0]
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip auth for specified paths
			for _, path := range cfg.SkipPaths {
				if r.URL.Path == path {
					next.ServeHTTP(w, r)
					return
				}
			}

			// Get authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				writeUnauthorizedError(w, "Missing authorization header")
				return
			}

			// Parse token
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || parts[0] != cfg.TokenType {
				writeUnauthorizedError(w, "Invalid authorization header format")
				return
			}

			token := parts[1]
			if token == "" {
				writeUnauthorizedError(w, "Missing token")
				return
			}

			// TODO: Validate token (implement your token validation logic)
			userID, err := validateToken(token)
			if err != nil {
				writeUnauthorizedError(w, "Invalid token")
				return
			}

			// Add user ID to context
			ctx := context.WithValue(r.Context(), constants.ContextKeyUserID, userID)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

// validateToken validates the JWT token and returns user ID
// TODO: Implement actual token validation logic
func validateToken(token string) (string, error) {
	// This is a placeholder implementation
	// In a real application, you would:
	// 1. Parse and validate JWT token
	// 2. Check token expiration
	// 3. Verify token signature
	// 4. Extract user ID from claims

	if token == "valid-token" {
		return "user-123", nil
	}

	return "", fmt.Errorf("invalid token")
}

func writeUnauthorizedError(w http.ResponseWriter, message string) {
	errorResponse := map[string]interface{}{
		"error": message,
		"code":  constants.ErrCodeUnauthorized,
	}
	utils.WriteJSON(w, http.StatusUnauthorized, errorResponse)
}
