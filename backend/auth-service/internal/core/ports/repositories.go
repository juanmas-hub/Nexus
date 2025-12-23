package ports

import "nexus/auth-service/internal/core/domain"

type UserRepository interface {
    GetByEmail(email string) (*domain.User, error)
    Save(user *domain.User) error
}