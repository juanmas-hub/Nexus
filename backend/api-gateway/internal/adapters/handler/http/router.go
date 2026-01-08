package http

import (
	"net/http"
	"github.com/juanmas-hub/nexus/backend/api-gateway/internal/core/services"
	"github.com/go-chi/chi/v5"
)

type GatewayHandler struct {
	service *services.GatewayService
}

func NewGatewayHandler(s *services.GatewayService) *GatewayHandler {
	return &GatewayHandler{service: s}
}

func (h *GatewayHandler) SetupRoutes(r chi.Router) {
	r.Post("/login", h.Login)
	r.Post("/register", h.Register)

	r.Get("/events", h.GetEvents)
}

func (h *GatewayHandler) Login(w http.ResponseWriter, r *http.Request) {
	h.service.Login(w, r)
}

func (h *GatewayHandler) Register(w http.ResponseWriter, r *http.Request) {
	h.service.Register(w, r)
}

func (h *GatewayHandler) GetEvents(w http.ResponseWriter, r *http.Request) {
    h.service.GetEvents(w, r)
}