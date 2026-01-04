package memory

import (
	"errors"
	"sync"
	"nexus/auth-service/internal/core/domain"
)

type memoryUserRepository struct {
	sync.RWMutex
	users map[string]domain.User
}

func NewMemoryUserRepository() *memoryUserRepository {
	return &memoryUserRepository{
		users: make(map[string]domain.User),
	}
}

func (r *memoryUserRepository) Save(user domain.User) error {
	r.Lock()
	defer r.Unlock()
	r.users[user.Email] = user
	return nil
}

func (r *memoryUserRepository) GetByEmail(email string) (domain.User, error) {
	r.RLock()
	defer r.RUnlock()
	user, ok := r.users[email]
	if !ok {
		return domain.User{}, errors.New("usuario no encontrado")
	}
	return user, nil
}