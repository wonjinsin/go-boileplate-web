package repository

import "github.com/wonjinsin/go-boilerplate/internal/domain"

// UserRepository defines the interface for user data access
type UserRepository interface {
	Save(*domain.User) error
	FindByID(id string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	List(offset, limit int) ([]*domain.User, error)
}
