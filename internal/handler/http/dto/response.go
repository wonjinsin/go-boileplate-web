package dto

// StandardResponse represents the standard HTTP response format
type StandardResponse struct {
	TrID   string `json:"trid"`
	Code   string `json:"code"`
	Result any    `json:"result,omitempty"`
}

// ErrorResult represents the error result structure
type ErrorResult struct {
	Msg string `json:"msg"`
}
