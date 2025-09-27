package domain

// User domain constants
const (
	// User status
	UserStatusActive   = "active"
	UserStatusInactive = "inactive"
	UserStatusPending  = "pending"
	UserStatusBanned   = "banned"
)

// User roles (if needed)
const (
	RoleAdmin = "admin"
	RoleUser  = "user"
	RoleGuest = "guest"
)

// User validation specific to domain
const (
	MinUserAge = 13
	MaxUserAge = 120
)

// Business rules
const (
	MaxUsersPerOrganization = 1000
	MaxLoginAttempts        = 5
	AccountLockoutDuration  = 1800 // seconds (30 minutes)
)
