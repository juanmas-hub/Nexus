package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	
	httpHandler "github.com/juanmas-hub/nexus/backend/api-gateway/internal/adapters/handler/http"
	"github.com/juanmas-hub/nexus/backend/api-gateway/internal/adapters/proxy"
	"github.com/juanmas-hub/nexus/backend/api-gateway/internal/core/services"
)

func main() {
	authServiceURL := os.Getenv("AUTH_SERVICE_URL")
	port := os.Getenv("PORT")
	if port == "" { port = "8080" }

	authProxy, err := proxy.NewHTTPProxy(authServiceURL)
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

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}