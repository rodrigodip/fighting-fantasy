package userhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/fighting-fantasy/internal/application/user"
)

type UserHandlerRepo interface {
	RegisterUser(c *gin.Context)
	FindByEmail(c *gin.Context)
	// GetTasks(c *gin.Context)
	// GetTask(c *gin.Context)
	// DeleteTask(c *gin.Context)
	// UpdateTask(c *gin.Context)
	// SetTaskDone(c *gin.Context)
}

type UserHandler struct {
	usecase *userapp.UserUseCase
}

func NewUserHandler(uc *userapp.UserUseCase) *UserHandler {
	return &UserHandler{usecase: uc}
}
func (uh *UserHandler) FindByEmail(c *gin.Context) {
	var req UserFindByEmailRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	foundUser, err := uh.usecase.GetEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
	}
	output := UserCreateResponse{
		UserID: foundUser.UserID,
		Name:   foundUser.Name,
		Email:  foundUser.Email,
		Roles:  foundUser.Roles,
	}
	c.JSON(http.StatusOK, output)
}
func (uh *UserHandler) RegisterUser(c *gin.Context) {
	var req UserCreateRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	newUser, err := uh.usecase.
		CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	output := UserCreateResponse{
		UserID: newUser.UserID,
		Name:   newUser.Name,
		Email:  newUser.Email,
		Roles:  newUser.Roles,
	}
	c.JSON(http.StatusCreated, output)
}
