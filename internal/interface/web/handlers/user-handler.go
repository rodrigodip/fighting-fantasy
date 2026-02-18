package webhandlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/fighting-fantasy/internal/application/user"
	"net/http"
)

type UserWebHandler struct {
	usecase *userapp.UserUseCase
}

func NewUserWebHandler(usecase *userapp.UserUseCase) *UserWebHandler {
	return &UserWebHandler{usecase: usecase}
}

// Show the registration form
func (h *UserWebHandler) ShowRegisterForm(c *gin.Context) {
	c.HTML(http.StatusOK, "register-form.html", gin.H{
		"Title": "Register",
	})
}

// Handle form submission
func (h *UserWebHandler) CreateUserFromWeb(c *gin.Context) {
	// Bind form data (not JSON)
	var req struct {
		Name     string `form:"name" binding:"required"`
		Email    string `form:"email" binding:"required,email"`
		Password string `form:"password" binding:"required"`
	}

	if err := c.ShouldBind(&req); err != nil {
		c.HTML(http.StatusBadRequest, "register-error.html", gin.H{
			"Error": err.Error(),
		})
		return
	}

	// Call shared usecase
	newUser, err := h.usecase.CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		c.HTML(http.StatusBadRequest, "register-error.html", gin.H{
			"Error": err.Error(),
		})
		return
	}

	// Render success template
	c.HTML(http.StatusCreated, "register-success.html", gin.H{
		"UserName": newUser.Name,
		"UserID":   newUser.UserID,
	})
}
