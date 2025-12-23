package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

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
	if err != nil {
		return "", errors.New("credenciales inválidas")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("credenciales inválidas")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKeyString := os.Getenv("JWT_SECRET")

	if secretKeyString == "" {
		return "", errors.New("error interno: configuración de seguridad faltante")
	}

	tokenString, err := token.SignedString([]byte(secretKeyString))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) Register(user *domain.User) error {
    existing, _ := s.repo.GetByEmail(user.Email)
    if existing != nil {
        return errors.New("el usuario ya está registrado")
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)

    return s.repo.Save(user)
}