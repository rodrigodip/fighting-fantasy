package webhandlers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/rodrigodip/fighting-fantasy/internal/application/user"
	"github.com/rodrigodip/fighting-fantasy/internal/interface/web/viewmodels"
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
// Returns: Sets session cookie, redirects to /dashboard or retunrs a error
func (uc *UserWebHandler) AuthLoginHandler(c *gin.Context) {
	// TODO: Use DTO here
	var req struct {
		Email    string `form:"email"`
		Password string `form:"password"`
	}
	if err := c.ShouldBind(&req); err != nil {
		tmpl := template.Must(template.New("").ParseFiles(
			"internal/interface/web/templates/partials/auth-feedback.html",
		))
		data := viewmodels.PageData{
			FeedBack: &viewmodels.Message{Error: "Invalid email or password"}}
		tmpl.ExecuteTemplate(c.Writer, "auth-feedback", data)
		return
	}
	token, err := uc.usecase.Login(req.Email, req.Password)
	if err != nil {
		tmpl := template.Must(template.New("").ParseFiles(
			"internal/interface/web/templates/partials/auth-feedback.html",
		))
		data := viewmodels.PageData{
			FeedBack: &viewmodels.Message{Error: weberrors.ParseErrorForWeb(err.Error())}}
		tmpl.ExecuteTemplate(c.Writer, "auth-feedback", data)

		return
	}
	// Store token in session
	session, _ := uc.sessionStore.Get(c.Request, "user-session")
	session.Values["token"] = token
	// Decode the JWT to store user_id/role directly:
	// session.Values["user_id"] = userID
	// session.Values["role"] = role
	if err := session.Save(c.Request, c.Writer); err != nil {
		tmpl := template.Must(template.New("").ParseFiles(
			"internal/interface/web/templates/partials/auth-feedback.html",
		))
		data := viewmodels.PageData{
			FeedBack: &viewmodels.Message{Error: "Invalid email or password"}}
		tmpl.ExecuteTemplate(c.Writer, "auth-feedback", data)
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
		tmpl := template.Must(template.New("").ParseFiles(
			"internal/interface/web/templates/partials/auth-feedback.html",
		))
		data := viewmodels.PageData{
			FeedBack: &viewmodels.Message{Error: weberrors.ParseErrorForWeb(err.Error())}}
		tmpl.ExecuteTemplate(c.Writer, "auth-feedback", data)
		return
	}
	_, err := uc.usecase.
		CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		tmpl := template.Must(template.New("").ParseFiles(
			"internal/interface/web/templates/partials/auth-feedback.html",
		))
		data := viewmodels.PageData{
			FeedBack: &viewmodels.Message{Error: weberrors.ParseErrorForWeb(err.Error())}}
		tmpl.ExecuteTemplate(c.Writer, "auth-feedback", data)
		return
	}
	tmpl := template.Must(template.New("").ParseFiles(
		"internal/interface/web/templates/partials/auth-feedback.html",
	))
	data := viewmodels.PageData{
		FeedBack: &viewmodels.Message{Success: "Your account has been successfully created"}}
	tmpl.ExecuteTemplate(c.Writer, "auth-feedback", data)
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
		tmpl := template.Must(template.New("").ParseFiles(
			"internal/interface/web/templates/partials/auth-feedback.html",
		))
		data := viewmodels.PageData{
			FeedBack: &viewmodels.Message{Error: weberrors.ParseErrorForWeb("missing token")}}
		tmpl.ExecuteTemplate(c.Writer, "auth-feedback", data)
		return
	}
	err := uc.usecase.VerifyEmail(token)
	if err != nil {
		tmpl := template.Must(template.New("").ParseFiles(
			"internal/interface/web/templates/partials/auth-feedback.html",
		))
		data := viewmodels.PageData{
			FeedBack: &viewmodels.Message{Error: weberrors.ParseErrorForWeb(err.Error())}}
		tmpl.ExecuteTemplate(c.Writer, "auth-feedback", data)
		return
	}

	c.Redirect(http.StatusFound, "/verify")
}
