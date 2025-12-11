package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"nexus/auth-service/internal/core/ports"
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
	// 1. Buscar usuario
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", errors.New("credenciales inválidas")
	}

	// 2. Verificar contraseña
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("credenciales inválidas")
	}

	// 3. Crear Claims del Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // 24 horas
	})

	// 4. Obtener el secreto de la variable de entorno
	secretKeyString := os.Getenv("JWT_SECRET")

	// Es importante validar que la variable exista. 
	// Si se te olvidó ponerla, esto evitará que la app firme tokens vacíos o inseguros.
	if secretKeyString == "" {
		return "", errors.New("error interno: configuración de seguridad faltante")
	}

	// 5. Firmar el token
	tokenString, err := token.SignedString([]byte(secretKeyString))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}