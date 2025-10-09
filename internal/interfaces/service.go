package interfaces

import "example/internal/domain"

// UserService defines the interface for user business logic
type UserService interface {
	CreateUser(name, email string) (*domain.User, error)
	GetUser(id string) (*domain.User, error)
	ListUsers(offset, limit int) ([]*domain.User, error)
}
