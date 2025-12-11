package http

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "nexus/auth-service/internal/core/ports"
)

type AuthHandler struct {
    authService ports.AuthService //Input Port
}

func NewAuthHandler(service ports.AuthService) *AuthHandler {
    return &AuthHandler{
        authService: service,
    }
}

// DTO para el request (Solo vive en la capa de adaptadores)
type loginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Login(c *gin.Context) {
    var req loginRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    token, err := h.authService.Login(req.Email, req.Password)
    if err != nil {
       
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
        return
    }

    c.SetCookie("auth_token", token, 3600, "/", "localhost", false, true)
    c.JSON(http.StatusOK, gin.H{"message": "Login exitoso"})
}