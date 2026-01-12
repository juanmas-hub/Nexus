package http

import (
    "net/http"
    "github.com/gin-gonic/gin"

    "nexus/auth-service/internal/core/ports"
    "nexus/auth-service/internal/core/domain"
    "nexus/auth-service/internal/core/services"
)

type AuthHandler struct {
	authService ports.AuthService
}

func NewAuthHandler(s *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: s}
}

func (h *AuthHandler) SetupRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", h.Login)
		auth.POST("/register", h.Register)
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req domain.LoginRequest

	if !DecodeJSON(c, &req) {
		return
	}

	user, token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		RespondWithError(c, http.StatusUnauthorized, "Credenciales inválidas")
		return
	}

	RespondWithJSON(c, http.StatusOK, domain.LoginResponse{
		Token: token,
		User:  user.ToDTO(),
	})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req domain.RegisterRequest

	if !DecodeJSON(c, &req) {
		return
	}

	err := h.authService.Register(req.ToDomain()) 
	if err != nil {
		RespondWithError(c, http.StatusConflict, err.Error())
		return
	}

	RespondWithJSON(c, http.StatusCreated, gin.H{"message": "Usuario creado exitosamente"})
}