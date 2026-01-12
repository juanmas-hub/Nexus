package main

import(
	"github.com/go-chi/chi/v5"
	
	"github.com/juanmas-hub/nexus/backend/api-gateway/internal/config"
	"github.com/juanmas-hub/nexus/backend/api-gateway/internal/adapters/clients"
	"github.com/juanmas-hub/nexus/backend/api-gateway/internal/core/services"

	httpHandler "github.com/juanmas-hub/nexus/backend/api-gateway/internal/adapters/handler/http"
)

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
