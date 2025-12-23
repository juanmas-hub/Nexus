package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-contrib/cors"

	"nexus/auth-service/internal/adapters/handler/http"
	"nexus/auth-service/internal/adapters/repository/postgres"
	"nexus/auth-service/internal/core/domain"
	"nexus/auth-service/internal/core/services"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Info: .env file not found, using system environment variables")
	}

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	log.Println("Running database migrations...")
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatal("Database migration failed: ", err)
	}
	log.Println("Database migration completed successfully.")

	userRepo := repository.NewPostgresRepository(db)
	authService := services.NewAuthService(userRepo)
	authHandler := http.NewAuthHandler(authService)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"https://nexus-b6b.pages.dev",
		},
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.POST("/auth/login", authHandler.Login)
	r.POST("/auth/register", authHandler.Register)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server started on port :%s", port)
	r.Run(":" + port)
}