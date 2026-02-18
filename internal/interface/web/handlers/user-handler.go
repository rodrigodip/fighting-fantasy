package webhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/rodrigodip/fighting-fantasy/internal/application/user"
)

type UserWebHandler struct {
	usecase      *userapp.UserUseCase
	sessionStore sessions.Store
}

func NewUserWebHandler(usecase *userapp.UserUseCase, store sessions.Store) *UserWebHandler {
	return &UserWebHandler{
		usecase:      usecase,
		sessionStore: store,
	}
}

// Show landing page
func (h *UserWebHandler) LandingPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "Fighting Fantasy",
	})
}

// Show the registration form
func (h *UserWebHandler) ShowRegisterForm(c *gin.Context) {
	c.HTML(http.StatusOK, "register-form.html", gin.H{
		"Title": "Register",
	})
}

// Show login form (HTMX partial)
func (h *UserWebHandler) ShowLoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "login-form.html", nil)
}

// Handle login submission
func (h *UserWebHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `form:"email" binding:"required,email"`
		Password string `form:"password" binding:"required"`
	}

	if err := c.ShouldBind(&req); err != nil {
		c.HTML(http.StatusBadRequest, "login-error.html", gin.H{
			"Error": err.Error(),
		})
		return
	}

	// Call usecase - returns JWT token
	token, err := h.usecase.Login(req.Email, req.Password)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login-error.html", gin.H{
			"Error": "Invalid credentials",
		})
		return
	}

	// Store token in session
	session, _ := h.sessionStore.Get(c.Request, "session-name")
	session.Values["token"] = token
	// You can also decode the JWT to store user_id/role directly:
	// session.Values["user_id"] = userID
	// session.Values["role"] = role
	if err := session.Save(c.Request, c.Writer); err != nil {
		c.HTML(http.StatusInternalServerError, "login-error.html", gin.H{
			"Error": "Failed to create session",
		})
		return
	}

	// Redirect to dashboard or return HTMX response
	c.Header("HX-Redirect", "/dashboard")
	c.Status(http.StatusOK)
}

func (h *UserWebHandler) Logout(c *gin.Context) {
	session, _ := h.sessionStore.Get(c.Request, "session-name")
	session.Values["token"] = ""
	session.Options.MaxAge = -1 // Delete cookie
	session.Save(c.Request, c.Writer)

	c.Redirect(http.StatusSeeOther, "/")
}
