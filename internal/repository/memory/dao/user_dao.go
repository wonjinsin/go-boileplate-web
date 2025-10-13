package dao

import (
	"errors"
	"sort"
	"sync"
	"time"

	"github.com/wonjinsin/go-boilerplate/internal/domain"
)

// User is the internal data model for persistence
type User struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
}

// userDAO handles data access operations
type userDAO struct {
	mu     sync.RWMutex
	byID   map[string]*User
	byMail map[string]*User
}

// NewUserDAO creates a new user DAO instance
func NewUserDAO() *userDAO {
	return &userDAO{
		byID:   make(map[string]*User),
		byMail: make(map[string]*User),
	}
}

// toDomainUser converts dao.User to domain.User
func toDomainUser(d *User) *domain.User {
	return &domain.User{
		ID:        d.ID,
		Name:      d.Name,
		Email:     d.Email,
		CreatedAt: d.CreatedAt,
	}
}

// toDAOUser converts domain.User to dao.User
func toDAOUser(u *domain.User) *User {
	return &User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	}
}

// Insert saves a user to storage
func (d *userDAO) Insert(u *domain.User) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	dao := toDAOUser(u)

	if _, ok := d.byMail[dao.Email]; ok {
		return errors.New("duplicate email")
	}

	d.byID[dao.ID] = dao
	d.byMail[dao.Email] = dao
	return nil
}

// SelectByID finds a user by ID
func (d *userDAO) SelectByID(id string) (*domain.User, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if dao, ok := d.byID[id]; ok {
		return toDomainUser(dao), nil
	}
	return nil, errors.New("not found")
}

// SelectByEmail finds a user by email
func (d *userDAO) SelectByEmail(email string) (*domain.User, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if dao, ok := d.byMail[email]; ok {
		return toDomainUser(dao), nil
	}
	return nil, errors.New("not found")
}

// SelectAll retrieves all users with pagination
func (d *userDAO) SelectAll(offset, limit int) ([]*domain.User, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	// Create array of dao users
	daoUsers := make([]*User, 0, len(d.byID))
	for _, dao := range d.byID {
		daoUsers = append(daoUsers, dao)
	}

	// Sort by CreatedAt
	sort.Slice(daoUsers, func(i, j int) bool {
		return daoUsers[i].CreatedAt.Before(daoUsers[j].CreatedAt)
	})

	if offset >= len(daoUsers) {
		return []*domain.User{}, nil
	}

	end := offset + limit
	if end > len(daoUsers) {
		end = len(daoUsers)
	}

	// Convert to domain users
	result := make([]*domain.User, 0, end-offset)
	for _, dao := range daoUsers[offset:end] {
		result = append(result, toDomainUser(dao))
	}

	return result, nil
}
