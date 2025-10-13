package usecase

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/wonjinsin/go-boilerplate/internal/domain"
	"github.com/wonjinsin/go-boilerplate/internal/repository"
	"github.com/wonjinsin/go-boilerplate/pkg/logger"
	"github.com/wonjinsin/go-boilerplate/pkg/utils"
)

type userService struct {
	repo repository.UserRepository
	seq  atomic.Int64
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) nextID() string {
	c := s.seq.Add(1)
	return utils.GenerateID(c)
}

func (s *userService) CreateUser(ctx context.Context, name, email string) (*domain.User, error) {
	if existing, _ := s.repo.FindByEmail(email); existing != nil {
		logger.LogWarn(ctx, "duplicate email attempted")
		return nil, domain.ErrDuplicateEmail
	}
	id := s.nextID()
	u, err := domain.NewUser(id, name, email, time.Now())
	if err != nil {
		logger.LogError(ctx, "failed to create user domain object", err)
		return nil, err
	}
	if err := s.repo.Save(u); err != nil {
		logger.LogError(ctx, "failed to save user to repository", err)
		return nil, err
	}
	logger.LogInfo(ctx, "user created successfully")
	return u, nil
}

func (s *userService) GetUser(ctx context.Context, id string) (*domain.User, error) {
	u, err := s.repo.FindByID(id)
	if err != nil {
		logger.LogError(ctx, "failed to find user by ID", err)
		return nil, err
	}
	if u == nil {
		logger.LogWarn(ctx, "user not found")
		return nil, domain.ErrUserNotFound
	}
	logger.LogInfo(ctx, "user found successfully")
	return u, nil
}

func (s *userService) ListUsers(ctx context.Context, offset, limit int) ([]*domain.User, error) {
	if limit <= 0 {
		limit = 50
	}
	users, err := s.repo.List(offset, limit)
	if err != nil {
		logger.LogError(ctx, "failed to list users", err)
		return nil, err
	}
	return users, nil
}
