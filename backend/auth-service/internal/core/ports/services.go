package ports
import(
    "nexus/auth-service/internal/core/domain"
)

type AuthService interface {
    Login(email, password string) (string, error)
    Register(user *domain.User) error
}