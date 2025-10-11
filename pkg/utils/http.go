package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/wonjinsin/go-boilerplate/pkg/constants"
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

// WriteStandardJSON writes a standard JSON response with TrID
func WriteStandardJSON(w http.ResponseWriter, r *http.Request, code int, result any) {
	// Get TrID from context
	trID := ""
	if ctx := r.Context(); ctx != nil {
		if id, ok := ctx.Value(constants.ContextKeyTrID).(string); ok {
			trID = id
		}
	}

	// Format code as 4-digit string (e.g., 200 -> "0200")
	codeStr := fmt.Sprintf("%04d", code)

	// Create standard response
	response := map[string]any{
		"trid": trID,
		"code": codeStr,
	}

	// Only add result if it's not nil
	if result != nil {
		response["result"] = result
	}

	w.Header().Set(constants.HeaderContentType, constants.ContentTypeJSONCharset)
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("json encode error: %v", err)
	}
}
