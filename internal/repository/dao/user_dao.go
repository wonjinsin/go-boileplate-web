package dao

import (
	"time"

	"example/internal/domain"
)

// UserDAO represents the data access object for user data
type UserDAO struct {
	ID        string     `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Email     string     `json:"email" db:"email"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// TableName returns the table name for UserDAO
func (UserDAO) TableName() string {
	return "users"
}

// ToDomain converts UserDAO to domain.User
func (dao *UserDAO) ToDomain() *domain.User {
	return &domain.User{
		ID:        dao.ID,
		Name:      dao.Name,
		Email:     dao.Email,
		CreatedAt: dao.CreatedAt,
	}
}

// FromDomain converts domain.User to UserDAO
func (dao *UserDAO) FromDomain(user *domain.User) {
	dao.ID = user.ID
	dao.Name = user.Name
	dao.Email = user.Email
	dao.CreatedAt = user.CreatedAt
	dao.UpdatedAt = time.Now()
}

// NewUserDAO creates a new UserDAO from domain.User
func NewUserDAO(user *domain.User) *UserDAO {
	dao := &UserDAO{}
	dao.FromDomain(user)
	return dao
}

// UserDAOList represents a slice of UserDAO
type UserDAOList []*UserDAO

// ToDomainList converts UserDAOList to slice of domain.User
func (list UserDAOList) ToDomainList() []*domain.User {
	users := make([]*domain.User, len(list))
	for i, dao := range list {
		users[i] = dao.ToDomain()
	}
	return users
}

// UserFilter represents filter criteria for user queries
type UserFilter struct {
	IDs       []string   `json:"ids,omitempty"`
	Emails    []string   `json:"emails,omitempty"`
	Name      *string    `json:"name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Limit     int        `json:"limit"`
	Offset    int        `json:"offset"`
}

// UserSort represents sorting criteria for user queries
type UserSort struct {
	Field string `json:"field"` // id, name, email, created_at
	Order string `json:"order"` // asc, desc
}

// DefaultUserSort returns default sorting criteria
func DefaultUserSort() UserSort {
	return UserSort{
		Field: "created_at",
		Order: "asc",
	}
}
