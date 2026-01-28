package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespondWithJSON(c *gin.Context, code int, payload interface{}) {
	c.JSON(code, payload)
}

func RespondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"error": message})
}

func DecodeJSON(c *gin.Context, request interface{}) bool {
	if err := c.ShouldBindJSON(request); err != nil {
		RespondWithError(c, http.StatusBadRequest, "Datos de entrada inválidos: " + err.Error())
		return false
	}
	return true
}