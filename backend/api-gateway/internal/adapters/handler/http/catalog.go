package http

import (
	"net/http"
	// "github.com/juanmas-hub/nexus/backend/api-gateway/internal/core/domain"
)

func (handler *GatewayHandler) GetEvents(w http.ResponseWriter, r *http.Request){
	getEventsResponse, err := handler.service.GetEvents(r.Context())

	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Credenciales inválidas o error de conexión")
		return
	}


	RespondWithJSON(w, http.StatusOK, getEventsResponse)
}