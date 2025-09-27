package constants

// Validation limits
const (
	// User validation
	MinNameLength     = 1
	MaxNameLength     = 200
	MinEmailLength    = 3
	MaxEmailLength    = 320 // RFC 5321 limit
	MinPasswordLength = 8
	MaxPasswordLength = 128
)

// General validation
const (
	MinIDLength     = 1
	MaxIDLength     = 64
	MinStringLength = 0
	MaxStringLength = 1000
)

// Regex patterns
const (
	EmailPattern    = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	UUIDPattern     = `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`
	AlphanumPattern = `^[a-zA-Z0-9]+$`
)

// ID generation
const (
	IDAlphabet      = "0123456789abcdefghijklmnopqrstuvwxyz"
	DefaultIDLength = 12
	UUIDLength      = 36
)
