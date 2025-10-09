package http

import (
	"net/http"
)

// HealthController handles health check endpoints
type HealthController struct{}

// NewHealthController creates a new health controller
func NewHealthController() *HealthController {
	return &HealthController{}
}

// Check handles health check requests
func (c *HealthController) Check(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

