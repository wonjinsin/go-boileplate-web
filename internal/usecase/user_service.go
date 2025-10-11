package usecase

import (
	"sync/atomic"
	"time"

	"github.com/wonjinsin/go-boilerplate/internal/domain"
	"github.com/wonjinsin/go-boilerplate/internal/interfaces"
	"github.com/wonjinsin/go-boilerplate/pkg/utils"
)

type userService struct {
	repo interfaces.UserRepository
	seq  atomic.Int64
}

func NewUserService(r interfaces.UserRepository) interfaces.UserService {
	return &userService{repo: r}
}

func (s *userService) nextID() string {
	c := s.seq.Add(1)
	return utils.GenerateID(c)
}

func (s *userService) CreateUser(name, email string) (*domain.User, error) {
	if existing, _ := s.repo.FindByEmail(email); existing != nil {
		return nil, domain.ErrDuplicateEmail
	}
	id := s.nextID()
	u, err := domain.NewUser(id, name, email, time.Now())
	if err != nil {
		return nil, err
	}
	if err := s.repo.Save(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *userService) GetUser(id string) (*domain.User, error) {
	u, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, domain.ErrUserNotFound
	}
	return u, nil
}

func (s *userService) ListUsers(offset, limit int) ([]*domain.User, error) {
	if limit <= 0 {
		limit = 50
	}
	return s.repo.List(offset, limit)
}
