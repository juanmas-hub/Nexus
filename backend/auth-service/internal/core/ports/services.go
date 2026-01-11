package ports
import(
    "nexus/auth-service/internal/core/domain"
)

type AuthService interface {
    Login(email, password string) (*domain.User, string, error)
    Register(user *domain.User) error
}