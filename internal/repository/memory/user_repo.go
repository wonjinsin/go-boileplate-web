package memory

import (
	"errors"
	"sort"
	"sync"

	"example/internal/domain"
	"example/internal/interfaces"
)

type userRepo struct {
	mu     sync.RWMutex
	byID   map[string]*domain.User
	byMail map[string]*domain.User
}

func NewUserRepository() interfaces.UserRepository {
	return &userRepo{
		byID:   make(map[string]*domain.User),
		byMail: make(map[string]*domain.User),
	}
}

func (r *userRepo) Save(u *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.byMail[u.Email]; ok {
		return domain.ErrDuplicateEmail
	}

	r.byID[u.ID] = u
	r.byMail[u.Email] = u
	return nil
}

func (r *userRepo) FindByID(id string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if user, ok := r.byID[id]; ok {
		return user, nil
	}
	return nil, domain.ErrUserNotFound
}

func (r *userRepo) FindByEmail(email string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if user, ok := r.byMail[email]; ok {
		return user, nil
	}
	return nil, errors.New("not found")
}

func (r *userRepo) List(offset, limit int) ([]*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Create array of users
	users := make([]*domain.User, 0, len(r.byID))
	for _, user := range r.byID {
		users = append(users, user)
	}

	// Sort by CreatedAt
	sort.Slice(users, func(i, j int) bool {
		return users[i].CreatedAt.Before(users[j].CreatedAt)
	})

	if offset >= len(users) {
		return []*domain.User{}, nil
	}

	end := offset + limit
	if end > len(users) {
		end = len(users)
	}

	return users[offset:end], nil
}
