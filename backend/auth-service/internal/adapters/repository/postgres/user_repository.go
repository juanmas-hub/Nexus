package repository

import (
    "gorm.io/gorm"
    "nexus/auth-service/internal/core/domain"
    "nexus/auth-service/internal/core/ports"
)

type PostgresRepository struct {
    db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) ports.UserRepository {
    return &PostgresRepository{db: db}
}

func (r *PostgresRepository) GetByEmail(email string) (*domain.User, error) {
    var userEntity domain.User

    result := r.db.Where("email = ?", email).First(&userEntity)

    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return nil, nil 
        }
        return nil, result.Error
    }

    return &userEntity, nil
}

func (r *PostgresRepository) Save(user *domain.User) error {
    return r.db.Create(user).Error
}