package repository

import (
    "log"
    "gorm.io/gorm"
	"gorm.io/driver/postgres"
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

func ConnectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	log.Println("Running database migrations...")
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		log.Fatal("Database migration failed: ", err)
	}
	return db
}
