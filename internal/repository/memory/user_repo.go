package memory

import (
	"errors"
	"sort"
	"sync"

	"example/internal/domain"
	"example/internal/repository/dao"
)

type userRepo struct {
	mu     sync.RWMutex
	byID   map[string]*dao.UserDAO
	byMail map[string]*dao.UserDAO
}

func NewUserRepository() domain.UserRepository {
	return &userRepo{byID: make(map[string]*dao.UserDAO), byMail: make(map[string]*dao.UserDAO)}
}

func (r *userRepo) Save(u *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Convert domain to DAO
	userDAO := dao.NewUserDAO(u)

	if _, ok := r.byMail[u.Email]; ok {
		return domain.ErrDuplicateEmail
	}

	r.byID[u.ID] = userDAO
	r.byMail[u.Email] = userDAO
	return nil
}

func (r *userRepo) FindByID(id string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if userDAO, ok := r.byID[id]; ok {
		// Convert DAO to domain
		return userDAO.ToDomain(), nil
	}
	return nil, domain.ErrUserNotFound
}

func (r *userRepo) FindByEmail(email string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if userDAO, ok := r.byMail[email]; ok {
		// Convert DAO to domain
		return userDAO.ToDomain(), nil
	}
	return nil, errors.New("not found")
}

func (r *userRepo) List(offset, limit int) ([]*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Create array of DAOs first
	daoArr := make([]*dao.UserDAO, 0, len(r.byID))
	for _, userDAO := range r.byID {
		daoArr = append(daoArr, userDAO)
	}

	// Sort DAOs by CreatedAt
	sort.Slice(daoArr, func(i, j int) bool {
		return daoArr[i].CreatedAt.Before(daoArr[j].CreatedAt)
	})

	if offset >= len(daoArr) {
		return []*domain.User{}, nil
	}

	end := offset + limit
	if end > len(daoArr) {
		end = len(daoArr)
	}

	// Convert DAOs to domain objects
	result := make([]*domain.User, end-offset)
	for i, userDAO := range daoArr[offset:end] {
		result[i] = userDAO.ToDomain()
	}

	return result, nil
}
