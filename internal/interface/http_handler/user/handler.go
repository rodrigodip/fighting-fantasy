package userhandler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/fighting-fantasy/internal/application/user"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/user"
)

type UserHandlerRepo interface {
	RegisterUser(c *gin.Context)
	Login(c *gin.Context)
	FindByEmail(c *gin.Context)
	VerifyEmail(c *gin.Context)
}

type UserHandler struct {
	usecase *userapp.UserUseCase
}

func NewUserHandler(uc *userapp.UserUseCase) *UserHandler {
	return &UserHandler{usecase: uc}
}

// RegisterUser persistes a new user
func (uh *UserHandler) RegisterUser(c *gin.Context) {
	var req UserCreateRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Handler": err.Error()})
		return
	}
	newUser, err := uh.usecase.
		CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Handler": err.Error()})
		return
	}
	output := UserCreateResponse{
		UserID: newUser.UserID,
		Name:   newUser.Name,
		Email:  newUser.Email,
		Role:   newUser.Role,
	}
	c.JSON(http.StatusCreated, output)
}
func (uh *UserHandler) Login(c *gin.Context) {
	var req AuthLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"JSONerror": err})
		return
	}
	token, err := uh.usecase.Login(req.Email, req.Password)
	var authErr *usr.AuthErr
	if errors.As(err, &authErr) {
		c.JSON(http.StatusUnauthorized, gin.H{authErr.Producer: authErr.Err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
func (uh *UserHandler) VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing token"})
		return
	}
	err := uh.usecase.VerifyEmail(token)
	var authErr *usr.AuthErr
	if errors.As(err, &authErr) {
		c.JSON(http.StatusBadRequest, gin.H{authErr.Producer: authErr.Err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "Account verified"})
}
