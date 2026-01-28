package main

import(
	"log"
	"github.com/gin-gonic/gin"
	"nexus/auth-service/internal/config"
	"nexus/auth-service/internal/core/ports"
	"nexus/auth-service/internal/core/services"
    "nexus/auth-service/internal/adapters/repository/memory"
	repository "nexus/auth-service/internal/adapters/repository/postgres"
	httpHandler "nexus/auth-service/internal/adapters/handler/http"
	
)

func setupDependencyInjection(cfg *config.Config) *gin.Engine {
    var userRepo ports.UserRepository

    if cfg.APP_MODE == "prod" {
        db := repository.ConnectDB(cfg.DATABASE_URL)
        userRepo = repository.NewPostgresRepository(db)
    } else {
        log.Println("Corriendo en modo: Development (Memory Storage)")
        userRepo = memory.NewMemoryUserRepository()
    }

    authService := services.NewAuthService(userRepo)

    authHandler := httpHandler.NewAuthHandler(authService)

    router := gin.Default()
    
    router.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok", "service": "auth-service"})
    })

    authHandler.SetupRoutes(router)

    return router
}