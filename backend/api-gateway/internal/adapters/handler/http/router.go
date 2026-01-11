package http

import (
	"net/http"
	"github.com/go-chi/chi/v5"

    "os"

	"github.com/juanmas-hub/nexus/backend/api-gateway/internal/core/services"
	"github.com/juanmas-hub/nexus/backend/api-gateway/internal/core/domain"
)

type GatewayHandler struct {
	service *services.GatewayService
}

func NewGatewayHandler(s *services.GatewayService) *GatewayHandler {
	return &GatewayHandler{service: s}
}

func (handler *GatewayHandler) SetupRoutes(router chi.Router) {
    ApplyInfrastructureMiddlewares(router)

    router.Get("/health", handler.HealthCheck)
    router.Post("/login", handler.Login)
    
    router.Post("/register", handler.Register)
    // router.Get("/events", handler.GetEvents)
}

func (handler *GatewayHandler) HealthCheck(responseWriter http.ResponseWriter, httpRequest *http.Request) {
    healthResponse := domain.HealthResponse{
        Status:  "ok",
        Service: "gateway",
    }
    
    RespondWithJSON(responseWriter, http.StatusOK, healthResponse)
}

func (handler *GatewayHandler) Login(responseWriter http.ResponseWriter, httpRequest *http.Request) {
    var loginRequest domain.LoginRequest

    if !DecodeJSONBody(responseWriter, httpRequest, &loginRequest) {
        return
    }

    loginResponse, err := handler.service.Login(httpRequest.Context(), loginRequest)
    if err != nil {
        RespondWithError(responseWriter, http.StatusUnauthorized, "Credenciales inválidas o error de conexión")
        return
    }

    isProd := os.Getenv("APP_ENV") == "prod"

    http.SetCookie(responseWriter, &http.Cookie{
        Name:     "auth_token",
        Value:    loginResponse.Token,
        Path:     "/",
        HttpOnly: true,
        Secure:   isProd,
        SameSite: http.SameSiteLaxMode,
        MaxAge:   3600,
    })

    RespondWithJSON(responseWriter, http.StatusOK, loginResponse)
}

func (handler *GatewayHandler) Register(responseWriter http.ResponseWriter, httpRequest *http.Request) {
    var registerRequest domain.RegisterRequest

    if !DecodeJSONBody(responseWriter, httpRequest, &registerRequest) {
        return
    }

    registerResponse, err := handler.service.Register(httpRequest.Context(), registerRequest)
    if err != nil {
        RespondWithError(responseWriter, http.StatusConflict, "No se pudo completar el registro: el usuario ya existe")
        return
    }

    RespondWithJSON(responseWriter, http.StatusOK, registerResponse)
}