package http

import (
	"net/http"

	//"github.com/juanmas-hub/nexus/backend/api-gateway/internal/config"
	"github.com/juanmas-hub/nexus/backend/api-gateway/internal/core/domain"
	"github.com/juanmas-hub/nexus/backend/api-gateway/internal/core/services"
)

type GatewayHandler struct {
	service *services.GatewayService
	isProd	bool
}

func NewGatewayHandler(s *services.GatewayService, isProd bool) *GatewayHandler {
    return &GatewayHandler{
        service: s,
        isProd:  isProd,
    }
}

func (handler *GatewayHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	healthResponse := domain.HealthResponse{
		Status:  "ok",
		Service: "gateway",
	}
	RespondWithJSON(w, http.StatusOK, healthResponse)
}