package services

import (
	"errors"
	"github.com/google/uuid"

	"nexus/auth-service/internal/core/ports"
	"nexus/auth-service/internal/core/domain"
)

type AuthService struct {
	repo ports.UserRepository
}

func NewAuthService(repo ports.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.repo.GetByEmail(email)
    
    if err != nil || user == nil {
        return "", errors.New("credenciales inválidas")
    }

    err = checkPasswordHash(password, user.Password)
    if err != nil {
        return "", errors.New("credenciales inválidas")
    }

    return generateToken(user)
}

func (s *AuthService) Register(user *domain.User) error {
	existing, _ := s.repo.GetByEmail(user.Email)
	if existing != nil {
		return errors.New("el usuario ya está registrado")
	}

	user.ID = uuid.New().String()
	
	hashed, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashed

	return s.repo.Save(user)
}