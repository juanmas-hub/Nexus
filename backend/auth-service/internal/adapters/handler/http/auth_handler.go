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

type registerRequest struct {
    Email     string `json:"email" binding:"required,email"`
    Password  string `json:"password" binding:"required,min=6"`
    FirstName string `json:"first_name" binding:"required"`
    LastName  string `json:"last_name" binding:"required"`
}

func (h *AuthHandler) Register(c *gin.Context) {
    var req registerRequest

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