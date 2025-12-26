package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	
	httpHandler "github.com/juanmas-hub/nexus/backend/api-gateway/internal/adapters/handler/http"
	"github.com/juanmas-hub/nexus/backend/api-gateway/internal/config"
	"github.com/juanmas-hub/nexus/backend/api-gateway/internal/adapters/proxy"
	"github.com/juanmas-hub/nexus/backend/api-gateway/internal/core/services"
)

func main() {
	cfg := config.Load()

	authProxy, err := proxy.NewHTTPProxy(cfg.AuthServiceURL)
	if err != nil {
		log.Fatalf("Error configurando Proxy de Auth: %v", err)
	}

	gatewayService := services.NewGatewayService(authProxy)

	gatewayHandler := httpHandler.NewGatewayHandler(gatewayService)

	r := chi.NewRouter()

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "https://nexus-b6b.pages.dev"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler)

	gatewayHandler.SetupRoutes(r)

	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}