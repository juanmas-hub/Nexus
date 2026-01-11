package http

import (
    "net/http"
    "github.com/gin-gonic/gin"

    "nexus/auth-service/internal/core/ports"
    "nexus/auth-service/internal/core/domain"
)

type AuthHandler struct {
    authService ports.AuthService
}

func NewAuthHandler(service ports.AuthService) *AuthHandler {
    return &AuthHandler{
        authService: service,
    }
}

func (handler *AuthHandler) Login(c *gin.Context) {
    var request domain.LoginRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, token, err := handler.authService.Login(request.Email, request.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
        return
    }

    c.JSON(http.StatusOK, domain.LoginResponse{
        Token: token,
        User: domain.UserDTO{
            ID:        user.ID,
            Email:     user.Email,
            FirstName: user.FirstName,
            LastName:  user.LastName,
        },
    })
}

func (h *AuthHandler) Register(c *gin.Context) {
    var req domain.RegisterRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := &domain.User{
        Email:     req.Email,
        Password:  req.Password,
        FirstName: req.FirstName,
        LastName:  req.LastName,
        Role:      domain.RoleUser,
    }

    if err := h.authService.Register(user); err != nil {
        c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Usuario creado exitosamente",
        "user": gin.H{
            "id":         user.ID,
            "email":      user.Email,
            "first_name": user.FirstName,
            "last_name":  user.LastName,
        },
    })
}