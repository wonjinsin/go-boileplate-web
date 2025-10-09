package domain

import (
	"errors"
	"time"

	"example/pkg/constants"
	"example/pkg/utils"
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
	name = utils.NormalizeName(name)
	email = utils.NormalizeEmail(email)

	if utils.IsEmptyOrWhitespace(name) || len(name) > constants.MaxNameLength {
		return nil, ErrInvalidName
	}
	if !utils.IsValidEmail(email) {
		return nil, errors.New("invalid email format")
	}
	return &User{ID: id, Name: name, Email: email, CreatedAt: now}, nil
}
