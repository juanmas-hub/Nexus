package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"nexus/auth-service/internal/adapters/handler/http"
	"nexus/auth-service/internal/adapters/repository/postgres"
	"nexus/auth-service/internal/core/domain"
	"nexus/auth-service/internal/core/services"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: No se encontró archivo .env")
	}

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("No se pudo conectar a la DB:", err)
	}

	log.Println("Ejecutando migraciones...")
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatal("Error en la migración de la base de datos:", err)
	}
	log.Println("Migración completada con éxito.")


	userRepo := repository.NewPostgresRepository(db)
	authService := services.NewAuthService(userRepo)
	authHandler := http.NewAuthHandler(authService)

	r := gin.Default()

	r.POST("/auth/login", authHandler.Login)

	log.Println("Servidor iniciado en el puerto :8080")
	r.Run(":8080")
}