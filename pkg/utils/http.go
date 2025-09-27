package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"example/pkg/constants"
)

// ParsePagination extracts pagination parameters from HTTP request
func ParsePagination(r *http.Request) (offset, limit int) {
	q := r.URL.Query()
	offset, _ = strconv.Atoi(q.Get("offset"))
	limit, _ = strconv.Atoi(q.Get("limit"))

	if offset < 0 {
		offset = constants.DefaultOffset
	}
	if limit <= 0 || limit > constants.MaxLimit {
		limit = constants.DefaultLimit
	}

	return offset, limit
}

// WriteJSON writes JSON response with proper headers
func WriteJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set(constants.HeaderContentType, constants.ContentTypeJSONCharset)
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("json encode error: %v", err)
	}
}

// ParseJSONBody parses JSON request body into the provided struct
func ParseJSONBody(r *http.Request, v any) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

// ExtractPathParam extracts path parameter from URL
// Example: ExtractPathParam("/users/123", "/users/") returns "123"
func ExtractPathParam(path, prefix string) string {
	if len(path) <= len(prefix) {
		return ""
	}
	return path[len(prefix):]
}

// SetCORSHeaders sets common CORS headers
func SetCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}
