package memory

import (
	"errors"
	"sync"
	"nexus/auth-service/internal/core/domain"
	"nexus/auth-service/internal/core/ports" // Asegúrate de importar esto
)

type memoryUserRepository struct {
	sync.RWMutex
	// IMPORTANTE: El mapa debe guardar punteros (*domain.User) como pide tu interfaz
	users map[string]*domain.User 
}

// Retornamos la interfaz del paquete ports para que main lo acepte
func NewMemoryUserRepository() ports.UserRepository {
	return &memoryUserRepository{
		users: make(map[string]*domain.User),
	}
}

// Método Save con puntero *domain.User para cumplir la interfaz
func (r *memoryUserRepository) Save(user *domain.User) error {
	r.Lock()
	defer r.Unlock()
	r.users[user.Email] = user
	return nil
}

// Método GetByEmail con retorno *domain.User para cumplir la interfaz
func (r *memoryUserRepository) GetByEmail(email string) (*domain.User, error) {
	r.RLock()
	defer r.RUnlock()
	user, ok := r.users[email]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}