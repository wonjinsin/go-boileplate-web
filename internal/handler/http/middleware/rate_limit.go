package middleware

import (
	"net/http"
	"sync"
	"time"

	"example/internal/constants"
	"example/pkg/utils"
)

// RateLimitConfig holds rate limiting configuration
type RateLimitConfig struct {
	RequestsPerMinute int
	BurstSize         int
	KeyGenerator      func(*http.Request) string
	SkipPaths         []string
}

// DefaultRateLimitConfig returns default rate limit configuration
func DefaultRateLimitConfig() RateLimitConfig {
	return RateLimitConfig{
		RequestsPerMinute: 60,
		BurstSize:         10,
		KeyGenerator: func(r *http.Request) string {
			return r.RemoteAddr
		},
		SkipPaths: []string{"/healthz"},
	}
}

// bucket represents a token bucket for rate limiting
type bucket struct {
	tokens     float64
	lastUpdate time.Time
	mu         sync.Mutex
}

// rateLimiter manages rate limiting buckets
type rateLimiter struct {
	buckets map[string]*bucket
	config  RateLimitConfig
	mu      sync.RWMutex
}

// newRateLimiter creates a new rate limiter
func newRateLimiter(config RateLimitConfig) *rateLimiter {
	return &rateLimiter{
		buckets: make(map[string]*bucket),
		config:  config,
	}
}

// allow checks if a request should be allowed
func (rl *rateLimiter) allow(key string) bool {
	rl.mu.RLock()
	b, exists := rl.buckets[key]
	rl.mu.RUnlock()

	if !exists {
		rl.mu.Lock()
		// Double-check after acquiring write lock
		if b, exists = rl.buckets[key]; !exists {
			b = &bucket{
				tokens:     float64(rl.config.BurstSize),
				lastUpdate: time.Now(),
			}
			rl.buckets[key] = b
		}
		rl.mu.Unlock()
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(b.lastUpdate).Minutes()

	// Add tokens based on elapsed time
	tokensToAdd := elapsed * float64(rl.config.RequestsPerMinute)
	b.tokens = min(float64(rl.config.BurstSize), b.tokens+tokensToAdd)
	b.lastUpdate = now

	// Check if we have tokens available
	if b.tokens >= 1 {
		b.tokens--
		return true
	}

	return false
}

// RateLimit returns a middleware that implements rate limiting
func RateLimit(config ...RateLimitConfig) func(http.Handler) http.Handler {
	cfg := DefaultRateLimitConfig()
	if len(config) > 0 {
		cfg = config[0]
	}

	limiter := newRateLimiter(cfg)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip rate limiting for specified paths
			for _, path := range cfg.SkipPaths {
				if r.URL.Path == path {
					next.ServeHTTP(w, r)
					return
				}
			}

			key := cfg.KeyGenerator(r)
			if !limiter.allow(key) {
				errorResponse := map[string]interface{}{
					"error": "Rate limit exceeded",
					"code":  constants.ErrCodeTooManyRequests,
				}
				utils.WriteJSON(w, http.StatusTooManyRequests, errorResponse)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
