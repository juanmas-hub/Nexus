package http

import (
	"github.com/go-chi/chi/v5"
)

func (handler *GatewayHandler) SetupRoutes(router chi.Router) {
    //ApplyInfrastructureMiddlewares(router)

    router.Get("/health", handler.HealthCheck)
    router.Post("/login", handler.Login)
    
    router.Post("/register", handler.Register)
    // router.Get("/events", handler.GetEvents)
}

