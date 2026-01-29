package http

import (
	"net/http"
	"github.com/juanmas-hub/nexus/backend/api-gateway/internal/core/domain"
)

func (handler *GatewayHandler) GetEvents(w http.ResponseWriter, r *http.Request){
	getEventsResponse, err := handler.service.GetEvents()
}