package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	//"github.com/gin-contrib/cors"

	"nexus/auth-service/internal/core/ports" 
	"nexus/auth-service/internal/config"
    "nexus/auth-service/internal/adapters/repository/memory"
	"nexus/auth-service/internal/adapters/handler/http"
	repository "nexus/auth-service/internal/adapters/repository/postgres"
	"nexus/auth-service/internal/core/domain"
	"nexus/auth-service/internal/core/services"
)

func main() {
    cfg := config.Load()

    var userRepo ports.UserRepository

    if cfg.APP_MODE == "dev" {
        log.Println("Mode: Development")
        userRepo = memory.NewMemoryUserRepository() 
    } else {
        log.Println("Mode: Production")
        db := connectDB(cfg.DATABASE_URL)
        userRepo = repository.NewPostgresRepository(db)
    }

    authService := services.NewAuthService(userRepo)
    authHandler := http.NewAuthHandler(authService)

    r := setupRouter(authHandler)
    
    log.Printf("Server started on port :%s", cfg.Port)
    r.Run(":" + cfg.Port)
}

// AUX

func connectDB(dsn string) *gorm.DB {
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

func setupRouter(h *http.AuthHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status": "ok",
            "service": "auth-service",
        })
    })

	r.POST("/auth/login", h.Login)
	r.POST("/auth/register", h.Register)

	return r
}
