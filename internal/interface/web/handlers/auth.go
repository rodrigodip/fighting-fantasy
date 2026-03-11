package webhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/rodrigodip/fighting-fantasy/internal/application/user"
	"github.com/rodrigodip/fighting-fantasy/internal/pkg/errors/web_errors"
)

type UserWebHandlerRepo interface {
	AuthLoginHandler(c *gin.Context)
	AuthSignUpHandler(c *gin.Context)
	AuthLogoutHandle(c *gin.Context)
	AuthVerifyEmailHandler(c *gin.Context)
	// FindByEmail(c *gin.Context)
}

type UserWebHandler struct {
	usecase      *userapp.UserUseCase
	sessionStore sessions.Store
}

func NewUserWebHandler(uc *userapp.UserUseCase, store sessions.Store) *UserWebHandler {
	return &UserWebHandler{
		usecase:      uc,
		sessionStore: store,
	}
}

// AuthLoginHandler handles user login
// POST /auth/login
// Accepts form data: email (string), password (string)
// Returns: Sets session cookie and redirects to /dashboard
func (uc *UserWebHandler) AuthLoginHandler(c *gin.Context) {
	// TODO: Use DTO here
	var req struct {
		Email    string `form:"email"`
		Password string `form:"password"`
	}
	if err := c.ShouldBind(&req); err != nil {
		c.HTML(http.StatusOK, "auth-error.html", gin.H{
			"Error": weberrors.ParseErrorForWeb(err.Error()),
		})
		return
	}
	token, err := uc.usecase.Login(req.Email, req.Password)
	if err != nil {
		c.HTML(http.StatusOK, "auth-error.html", gin.H{
			"Error": weberrors.ParseErrorForWeb(err.Error()),
		})
		return
	}
	// Store token in session
	session, _ := uc.sessionStore.Get(c.Request, "user-session")
	session.Values["token"] = token
	// Decode the JWT to store user_id/role directly:
	// session.Values["user_id"] = userID
	// session.Values["role"] = role
	if err := session.Save(c.Request, c.Writer); err != nil {
		c.HTML(http.StatusInternalServerError, "auth-error.html", gin.H{
			"Error": "Failed to create session",
		})
		return
	}
	// Redirect to dashboard or return HTMX response
	c.Header("HX-Redirect", "/dashboard")
	c.Status(http.StatusOK)
}

// AuthSignUpHandler handles new user registration
// POST /auth/signup
// Accepts form data: name (string), email (string), password (string)
// Returns: Sets session cookie and redirects to /dashboard
func (uc *UserWebHandler) AuthSignUpHandler(c *gin.Context) {

	// TODO: Use DTO here
	var req struct {
		Name     string `form:"name"`
		Email    string `form:"email"`
		Password string `form:"password"`
	}
	if err := c.ShouldBind(&req); err != nil {
		c.HTML(http.StatusOK, "auth-error.html", gin.H{

			"Error": weberrors.ParseErrorForWeb(err.Error()),
		})
		return
	}
	_, err := uc.usecase.
		CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		c.HTML(http.StatusOK, "auth-error.html", gin.H{
			"Error": weberrors.ParseErrorForWeb(err.Error()),
		})
		return
	}
	// Redirect to Login or return HTMX response
	// TODO: Create Verify page
	c.Header("HX-Redirect", "/auth/verify")
	c.Status(http.StatusOK)
}

// AuthLogoutHandler handles user logout
// POST /auth/logout
// Returns: Clears session cookie and redirects to /
func (uc *UserWebHandler) AuthLogoutHandle(c *gin.Context) {
	session, _ := uc.sessionStore.Get(c.Request, "user-session")
	session.Values["token"] = ""
	session.Options.MaxAge = -1 // Delete cookie
	session.Save(c.Request, c.Writer)

	c.Redirect(http.StatusOK, "/")
}

// AuthLogoutHandler handles user logout
// POST /auth/logout
// Returns: Clears session cookie and redirects to /
func (uc *UserWebHandler) AuthVerifyEmailHandler(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.HTML(http.StatusOK, "auth-error.html", gin.H{
			"Error": weberrors.ParseErrorForWeb("missing token"),
		})
		return
	}
	err := uc.usecase.VerifyEmail(token)
	if err != nil {
		c.HTML(http.StatusOK, "auth-error.html", gin.H{
			"Error": weberrors.ParseErrorForWeb(err.Error()),
		})
		return
	}

	// c.HTML(http.StatusOK, "index.html", gin.H{
	// 	"Title": "Fighting Fantasy",
	// })

	c.Header("HX-Redirect", "/")
	c.Status(http.StatusOK)
}
