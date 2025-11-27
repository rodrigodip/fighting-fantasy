package userhandler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/fighting-fantasy/internal/application/user"
	authErr "github.com/rodrigodip/fighting-fantasy/internal/pkg/errors"
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
// @Summary Persistes a new user
// @Description Create a new User account.
// @Tags User
// @Accept json
// @Produce json
// @Param UserCreateRequest body userhandler.UserCreateRequest true "User params to create a User"
// @Success 200 {object} userhandler.UserCreateResponse
// @Failure 400
// @Failure 500 {string} string "Error: Parsing JSON"
// @Router /users [post]
func (uh *UserHandler) RegisterUser(c *gin.Context) {
	var req UserCreateRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Handler": err.Error()})
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

// Login allows a user to log in and obtain an authentication token.
// @Summary User Login
// @Description Allows a user to log in and receive an authentication token.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param AuthLoginRequst body userhandler.UserCreateRequest true "User login credentials"
// @Success 200
// @Header 200 {string} Authorization "Authentication token"
// @Failure 403 {string} string "Error: Invalid login credentials"
// @Router /login [post]
func (uh *UserHandler) Login(c *gin.Context) {
	var req AuthLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"JSONerror": err.Error()})
		return
	}
	token, err := uh.usecase.Login(req.Email, req.Password)
	var authErr *authErr.AuthErr
	if errors.As(err, &authErr) {
		c.JSON(http.StatusUnauthorized, gin.H{authErr.Producer: authErr.Err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (uh *UserHandler) VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "missing token"})
		return
	}
	err := uh.usecase.VerifyEmail(token)
	var authErr *authErr.AuthErr
	if errors.As(err, &authErr) {
		c.JSON(http.StatusBadRequest, gin.H{authErr.Producer: authErr.Err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "Account verified"})
}

func (uh *UserHandler) RegisterHero(c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		c.Abort()
		return
	}
	// token := c.GetHeader("Authorization")
	// if token == "" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
	// 	c.Abort()
	// 	return
	// }
	// t, err := uc.HeroService.CheckToken(token)
	// if err != nil {
	// 	return &hero.Hero{}, fmt.Errorf("CreateHero: %v", err)
	// }
	// claims := t.Claims.(jwt.MapClaims)
	// userID := claims["user_id"].(string)

	var req HeroCreateRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"JSON_ParseError": err.Error()})
		return
	}
	newHero, err := uh.usecase.CreateHero(userID.(string), req.HeroName, req.Potion)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"userId":    newHero.UserID,
		"hero_name": newHero.HeroName,
	})
}
