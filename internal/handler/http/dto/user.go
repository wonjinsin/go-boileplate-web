package dto

import (
	"time"

	"github.com/wonjinsin/go-boilerplate/internal/domain"
)

// CreateUserRequest represents the request payload for creating a user
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserResponse represents the response payload for user data
type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// UserListResponse represents the response payload for user list
type UserListResponse struct {
	Users  []UserResponse `json:"users"`
	Total  int            `json:"total"`
	Offset int            `json:"offset"`
	Limit  int            `json:"limit"`
}

// ErrorResponse represents the error response payload
type ErrorResponse struct {
	Error   string            `json:"error"`
	Code    string            `json:"code,omitempty"`
	Details map[string]string `json:"details,omitempty"`
}

// ToUserResponse converts domain.User to UserResponse
func ToUserResponse(user *domain.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

// ToUserListResponse converts slice of domain.User to UserListResponse
func ToUserListResponse(users []*domain.User, total, offset, limit int) UserListResponse {
	userResponses := make([]UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = ToUserResponse(user)
	}

	return UserListResponse{
		Users:  userResponses,
		Total:  total,
		Offset: offset,
		Limit:  limit,
	}
}
