package authhandler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/fighting-fantasy/internal/application/auth"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/auth"
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
		c.JSON(http.StatusBadRequest, gin.H{"JSONerror": err})
		return
	}

	token, err := h.usecase.Login(req.Email, req.Password)
	var authErr *auth.AuthErr
	if errors.As(err, &authErr) {
		c.JSON(http.StatusUnauthorized, gin.H{authErr.Producer: authErr.Err})
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

	err := h.usecase.VerifyEmail(token)
	var authErr *auth.AuthErr
	if errors.As(err, &authErr) {
		c.JSON(http.StatusBadRequest, gin.H{authErr.Producer: authErr.Err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Account verified"})
}
