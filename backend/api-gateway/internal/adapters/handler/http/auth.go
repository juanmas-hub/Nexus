package http

import (
	"net/http"
	"github.com/juanmas-hub/nexus/backend/api-gateway/internal/core/domain"
)

func (handler *GatewayHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest domain.LoginRequest

	if !DecodeJSONBody(w, r, &loginRequest) {
		return
	}

	loginResponse, err := handler.service.Login(r.Context(), loginRequest)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Credenciales inválidas o error de conexión")
		return
	}

	// Delegamos la creación de la cookie al helper
	handler.setAuthCookie(w, loginResponse.Token)

	RespondWithJSON(w, http.StatusOK, loginResponse)
}

func (handler *GatewayHandler) Register(w http.ResponseWriter, r *http.Request) {
	var registerRequest domain.RegisterRequest

	if !DecodeJSONBody(w, r, &registerRequest) {
		return
	}

	registerResponse, err := handler.service.Register(r.Context(), registerRequest)
	if err != nil {
		RespondWithError(w, http.StatusConflict, "No se pudo completar el registro: el usuario ya existe")
		return
	}

	RespondWithJSON(w, http.StatusOK, registerResponse)
}