package ports

import "../domain"

type UserRepository interface {
    GetByEmail(email string) (*domain.User, error)
}