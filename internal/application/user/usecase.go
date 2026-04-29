package userapp

import (
	"fmt"

	"github.com/rodrigodip/fighting-fantasy/internal/domain/user"
	authErr "github.com/rodrigodip/fighting-fantasy/internal/pkg/errors/auth_errors"
	IDgenerator "github.com/rodrigodip/fighting-fantasy/internal/pkg/id_generator"
	"github.com/rodrigodip/fighting-fantasy/internal/pkg/security"
)

type UserRepository interface {
	RegisterUser(u usr.User) error
	FindByEmail(email string) (*usr.User, error)
	FindById(id string) (*usr.User, error)
	Update(id, field, newData string) error

	//TODO:
	//ResendVeriyEmail()
	//FindHero(userID string) (*Hero, error)
}

type Authentication interface {
	GenerateToken(userID, role string) (string, error)
	ValidateToken(token string) (string, error)
}

type UserUseCase struct {
	service *usr.Service
	repo    UserRepository
	auth    Authentication
}

func NewusrUseCase(s *usr.Service, r UserRepository, a Authentication) *UserUseCase {
	return &UserUseCase{service: s, repo: r, auth: a}
}
func (uc *UserUseCase) FindById(id string) (*usr.User, error) {
	return uc.repo.FindById(id)
}

func (uc *UserUseCase) CreateUser(name, email, password string) (*usr.User, error) {
	if err := uc.service.ValidadeUserInput(name, email, password); err != nil {
		return &usr.User{}, fmt.Errorf("CreateUser: %v", err)
	}
	if err := uc.service.ValidatePassword(password); err != nil {
		return &usr.User{}, fmt.Errorf("CreateUser: %v", err)
	}
	encodedPassword, err := security.EncodePw(password)
	if err != nil {
		return &usr.User{}, fmt.Errorf("CreateUser: %v", err)
	}
	generator := IDgenerator.NewTimestampIDGenerator()
	newUser := &usr.User{
		UserID:   generator.NewID(),
		Name:     name,
		Email:    email,
		Password: encodedPassword,
		Role:     string(usr.RoleUser),
		Status:   usr.StatusUnverified,
	}
	if err := uc.repo.RegisterUser(*newUser); err != nil {
		return &usr.User{}, fmt.Errorf("CreateUser: %v", err)
	}
	token, err := uc.auth.GenerateToken(newUser.UserID, newUser.Role)
	if err != nil {
		return &usr.User{}, authErr.InvalidToken("LOGIN-1001")
	}
	if err := security.SendEmail(newUser.Name, newUser.Email, token); err != nil {
		return &usr.User{}, fmt.Errorf("CreateUser: %v", err)
	}
	return newUser, nil
}

// Login check User account integity and returns a tokin
func (uc *UserUseCase) Login(email, password string) (string, error) {
	if err := security.EmailFormatValidation(email); err != nil {
		return "", authErr.InvalidCredentials("LOGIN-0001")
	}
	foundUser, err := uc.repo.FindByEmail(email)
	if err != nil || foundUser == nil {
		return "", authErr.NotFound("LOGIN-0010")
	}
	if uc.service.IsUserVerified(foundUser.Status) {
		return "", authErr.NotVerified("LOGIN-0100")
	}
	if err := security.CheckPasswordHash(foundUser.Password, password); err != nil {
		return "", authErr.InvalidCredentials("LOGIN-1000")
	}
	token, err := uc.auth.GenerateToken(foundUser.UserID, foundUser.Role)
	if err != nil {
		return "", authErr.InvalidToken("LOGIN-1001")
	}
	return token, nil
}

// VerifyEmail check if User has a verified e-mail
func (uc *UserUseCase) VerifyEmail(token string) error {
	userID, err := uc.auth.ValidateToken(token)
	if err != nil {
		return authErr.InvalidToken("SRV-1010")
	}
	foundUser, err := uc.repo.FindById(userID)
	if err != nil || foundUser == nil {
		return authErr.NotFound("SRV-1001")
	}
	if err = uc.repo.Update(userID, "status", string(usr.StatusVerified)); err != nil {
		return fmt.Errorf("VerifyEmail: %v", err)
	}
	return nil
}
