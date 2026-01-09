package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	//"github.com/gin-contrib/cors"

	"nexus/auth-service/internal/core/ports" 
    "nexus/auth-service/internal/adapters/repository/memory"
	"nexus/auth-service/internal/adapters/handler/http"
	repository "nexus/auth-service/internal/adapters/repository/postgres"
	"nexus/auth-service/internal/core/domain"
	"nexus/auth-service/internal/core/services"
)

func main() {
    loadEnv()

    var userRepo ports.UserRepository 

    appMode := getEnv("APP_MODE", "prod")

    if appMode == "dev" {
        log.Println("Mode: Development")
        userRepo = memory.NewMemoryUserRepository() 
    } else {
        log.Println("Mode: Production")
        db := connectDB(os.Getenv("DATABASE_URL"))
        userRepo = repository.NewPostgresRepository(db)
    }

    authService := services.NewAuthService(userRepo)
    authHandler := http.NewAuthHandler(authService)

    r := setupRouter(authHandler)
    
    port := getEnv("PORT", "8081")
    log.Printf("Server started on port :%s", port)
    r.Run(":" + port)
}

// AUX

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Info: .env file not found, using system environment variables")
	}
}

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

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}