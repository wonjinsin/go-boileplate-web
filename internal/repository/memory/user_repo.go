package memory

import (
	"github.com/wonjinsin/go-boilerplate/internal/domain"
	"github.com/wonjinsin/go-boilerplate/internal/repository"
	"github.com/wonjinsin/go-boilerplate/internal/repository/memory/dao"
)

var userDAO = dao.NewUserDAO()

type userRepo struct{}

func NewUserRepository() repository.UserRepository {
	return &userRepo{}
}

func (r *userRepo) Save(u *domain.User) error {
	// Check for duplicate email first
	if existing, _ := userDAO.SelectByEmail(u.Email); existing != nil {
		return domain.ErrDuplicateEmail
	}
	return userDAO.Insert(u)
}

func (r *userRepo) FindByID(id string) (*domain.User, error) {
	user, err := userDAO.SelectByID(id)
	if err != nil {
		return nil, domain.ErrUserNotFound
	}
	return user, nil
}

func (r *userRepo) FindByEmail(email string) (*domain.User, error) {
	return userDAO.SelectByEmail(email)
}

func (r *userRepo) List(offset, limit int) ([]*domain.User, error) {
	return userDAO.SelectAll(offset, limit)
}
