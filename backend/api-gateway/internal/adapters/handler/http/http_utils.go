package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func DecodeJSONBody(responseWriter http.ResponseWriter, httpRequest *http.Request, target interface{}) bool {
	if err := json.NewDecoder(httpRequest.Body).Decode(target); err != nil {
		RespondWithError(responseWriter, http.StatusBadRequest, "Formato de datos inválido")
		return false
	}
	return true
}

func RespondWithJSON(responseWriter http.ResponseWriter, statusCode int, payload interface{}) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(statusCode)
	json.NewEncoder(responseWriter).Encode(payload)
}

func RespondWithError(responseWriter http.ResponseWriter, statusCode int, message string) {
	RespondWithJSON(responseWriter, statusCode, map[string]string{"error": message})
}

func ApplyInfrastructureMiddlewares(gatewayRouter chi.Router) {
	gatewayRouter.Use(middleware.Logger)
	gatewayRouter.Use(middleware.Recoverer)
}