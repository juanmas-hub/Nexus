package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/juanmas-hub/nexus/backend/api-gateway/internal/adapters/clients"
	httpHandler "github.com/juanmas-hub/nexus/backend/api-gateway/internal/adapters/handler/http"
	"github.com/juanmas-hub/nexus/backend/api-gateway/internal/config"
	"github.com/juanmas-hub/nexus/backend/api-gateway/internal/core/services"
)

func main() {
	cfg := config.Load()

	gatewayRouter := setupDependencyInjection(cfg)

	startServer(cfg.Port, gatewayRouter)
}

func setupDependencyInjection(configuration *config.Config) *chi.Mux {
	// Adaptadores de Salida
	authServiceClient := clients.NewHTTPAuthClient(configuration.AuthServiceURL, configuration.AuthServiceTimeout)

	// Servicios
	gatewayService := services.NewGatewayService(authServiceClient)

	// Adaptadores de Entrada
	gatewayHandler := httpHandler.NewGatewayHandler(gatewayService)

	// Configuración del router
	router := chi.NewRouter()
	httpHandler.ApplyCORSConfiguration(router, configuration.AllowedOrigins)
	httpHandler.ApplyInfrastructureMiddlewares(router)
	gatewayHandler.SetupRoutes(router)

	return router
}


func startServer(port string, router *chi.Mux) {
	serverAddress := ":" + port
	log.Printf("[GATEWAY START] %s", serverAddress)

	if err := http.ListenAndServe(serverAddress, router); err != nil {
		log.Fatalf("[CRITICAL ERROR] El servidor falló: %v", err)
	}
}