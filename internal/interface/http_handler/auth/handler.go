package authhandler

import (
	"github.com/rodrigodip/fighting-fantasy/internal/application/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandlerRepo interface {
	Login(c *gin.Context)
	VerifyEmail(c *gin.Context)
}

type AuthHandler struct {
	usecase *authapp.AuthUseCase
}

func NewAuthHandler(uc *authapp.AuthUseCase) *AuthHandler {
	return &AuthHandler{usecase: uc}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req AuthLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.usecase.Login(req.Email, req.Password)
	if err != nil {
		//TODO: Learn how to handling multiple errors from use case (wrap error)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
func (h *AuthHandler) VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing token"})
		return
	}

	if err := h.usecase.VerifyEmail(token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Account verified"})
}
