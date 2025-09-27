package domain

import (
	"errors"
	"strings"
	"time"
)

// Domain errors
var (
	ErrUserNotFound   = errors.New("user not found")
	ErrInvalidName    = errors.New("invalid name")
	ErrDuplicateEmail = errors.New("duplicate email")
)

// User is an aggregate root.
type User struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
}

func NewUser(id, name, email string, now time.Time) (*User, error) {
	name = strings.TrimSpace(name)
	email = strings.ToLower(strings.TrimSpace(email))
	if len(name) == 0 || len(name) > 200 {
		return nil, ErrInvalidName
	}
	if !looksLikeEmail(email) {
		return nil, errors.New("invalid email format")
	}
	return &User{ID: id, Name: name, Email: email, CreatedAt: now}, nil
}

func looksLikeEmail(s string) bool {
	if len(s) < 3 {
		return false
	}
	at := strings.Count(s, "@")
	return at == 1 && strings.Contains(s, ".")
}

// Repository (domain boundary)
type UserRepository interface {
	Save(*User) error
	FindByID(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	List(offset, limit int) ([]*User, error)
}
