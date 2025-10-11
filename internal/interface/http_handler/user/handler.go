package userhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/fighting-fantasy/internal/application/user"
)

type UserHandlerRepo interface {
	RegisterUser(c *gin.Context)
	FindByEmail(c *gin.Context)
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
