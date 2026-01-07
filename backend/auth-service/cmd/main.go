package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/gin-contrib/cors"

	"nexus/auth-service/internal/core/ports" 
    "nexus/auth-service/internal/adapters/repository/memory"
	"nexus/auth-service/internal/adapters/handler/http"
	repository "nexus/auth-service/internal/adapters/repository/postgres"
	"nexus/auth-service/internal/core/domain"
	"nexus/auth-service/internal/core/services"
)

func main() {
    loadEnv()

    // 1. Declaramos la variable usando la INTERFAZ (Puerto)
    var userRepo ports.UserRepository 

    // 2. Elegimos el adaptador según una variable de entorno
    appMode := getEnv("APP_MODE", "prod")

    if appMode == "dev" {
        log.Println("🔧 Mode: Development (Using In-Memory Database)")
        // Aquí inicializas tu adaptador de memoria que creamos antes
        userRepo = memory.NewMemoryUserRepository() 
    } else {
        log.Println("🚀 Mode: Production (Using Postgres)")
        db := connectDB(os.Getenv("DATABASE_URL"))
        userRepo = repository.NewPostgresRepository(db)
    }

    // 3. El resto del flujo sigue IGUAL, gracias a la inyección de dependencias
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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://nexus-b6b.pages.dev"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

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