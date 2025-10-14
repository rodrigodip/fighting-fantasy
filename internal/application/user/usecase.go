package userapp

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/user"
	IDgenerator "github.com/rodrigodip/fighting-fantasy/internal/pkg/id_generator"
	"github.com/rodrigodip/fighting-fantasy/internal/pkg/security"
)

type UserUseCase struct {
	UserService *usr.Service
}

func NewusrUseCase(service *usr.Service) *UserUseCase {
	return &UserUseCase{UserService: service}
}

func (uc *UserUseCase) CreateUser(name, email, password string) (*usr.User, error) {
	if err := uc.UserService.ValidadeUserInput(name, email, password); err != nil {
		return &usr.User{}, fmt.Errorf("CreateUser: %v", err)
	}
	if err := uc.UserService.ValidatePassword(password); err != nil {
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
	if err := uc.UserService.Save(*newUser); err != nil {
		return &usr.User{}, fmt.Errorf("CreateUser: %v", err)
	}
	if err := uc.UserService.SendVerifyEmail(
		newUser.UserID, newUser.Name, newUser.Email, newUser.Role,
	); err != nil {
		return &usr.User{}, fmt.Errorf("CreateUser: %v", err)
	}
	return newUser, nil
}

// Login check User account integity to use the system
func (uc *UserUseCase) Login(email, password string) (string, error) {
	if err := security.EmailFormatValidation(email); err != nil {
		return "", usr.InvalidCredentials("LOGIN-0001")
	}
	foundUser, err := uc.UserService.GetByEmail(email)
	if err != nil || foundUser == nil {
		return "", usr.NotFound("LOGIN-0010")
	}
	if uc.UserService.IsUserVerified(foundUser.Status) {
		return "", usr.NotVerified("LOGIN-0100")
	}
	if err := security.CheckPasswordHash(foundUser.Password, password); err != nil {
		return "", usr.InvalidCredentials("LOGIN-1000")
	}
	token, err := uc.UserService.GetToken(foundUser.UserID, foundUser.Role)
	if err != nil {
		return "", usr.InvalidToken("LOGIN-1001")
	}
	return token, nil
}

// VerifyEmail check if User has a verified e-mail
func (uc *UserUseCase) VerifyEmail(token string) error {
	t, err := uc.UserService.CheckToken(token)
	if err != nil {
		return usr.InvalidToken("SRV-1010")
	}
	claims := t.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	foundUser, err := uc.UserService.GetById(userID)
	if err != nil || foundUser == nil {
		return usr.NotFound("SRV-1001")
	}
	if err = uc.UserService.SetVerified(userID); err != nil {
		return fmt.Errorf("VerifyEmail: %v", err)
	}
	return nil
}
