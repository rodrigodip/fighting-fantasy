package userapp

import (
	"fmt"

	"github.com/rodrigodip/fighting-fantasy/internal/domain/user"
	IDgenerator "github.com/rodrigodip/fighting-fantasy/internal/pkg/id_generator"
	"github.com/rodrigodip/fighting-fantasy/internal/pkg/security"
)

type UserUseCase struct {
	UserService *user.Service
}

func NewUserUseCase(service *user.Service) *UserUseCase {
	return &UserUseCase{UserService: service}
}

// TODO: REFACTOR: Use case deve ser responsável por orquestrar os serviços
func (uc *UserUseCase) CreateUser(name, email, password string) (*user.User, error) {
	if err := uc.UserService.ValidadeUserInput(name, email, password); err != nil {
		return &user.User{}, fmt.Errorf("CreateUser: %v", err)
	}
	if err := uc.UserService.ValidatePassword(password); err != nil {
		return &user.User{}, fmt.Errorf("CreateUser: %v", err)
	}
	encodedPassword, err := security.EncodePw(password)
	if err != nil {
		return &user.User{}, fmt.Errorf("CreateUser: %v", err)
	}
	generator := IDgenerator.NewTimestampIDGenerator()
	newUser := &user.User{
		UserID:   generator.NewID(),
		Name:     name,
		Email:    email,
		Password: encodedPassword,
		Role:     "USER",
		Status:   "UNVERIFIED",
	}
	if err := uc.UserService.Save(*newUser); err != nil {
		return &user.User{}, fmt.Errorf("CreateUser: %v", err)
	}
	if err := uc.UserService.SendVerifyEmail(
		newUser.UserID, newUser.Name, newUser.Email, newUser.Role,
	); err != nil {
		return &user.User{}, fmt.Errorf("CreateUser: %v", err)
	}
	return newUser, nil
}
func (uc *UserUseCase) Login(email, password string) (string, error) {
	if err := security.EmailFormatValidation(email); err != nil {
		return "", user.InvalidCredentials("LOGIN-0001")
	}
	foundUser, err := uc.UserService.GetEmail(email)
	if err != nil || foundUser == nil {
		return "", user.NotFound("LOGIN-0010")
	}
	if uc.UserService.IsUserVerified(foundUser.Status) {
		return "", user.NotVerified("LOGIN-0100")
	}
	if err := security.CheckPasswordHash(foundUser.Password, password); err != nil {
		return "", user.InvalidCredentials("LOGIN-1000")
	}
	token, err := uc.UserService.GetToken(foundUser.UserID, foundUser.Role)
	if err != nil {
		return "", user.InvalidToken("LOGIN-1001")
	}
	return token, nil
}
func (uc *UserUseCase) VerifyEmail(token string) error {
	//NOTE: JWT na camada de Domínio configura acoplamento? Devo implementar na camada de interface?
	t, err := security.JWTService.ValidateToken(token)
	if err != nil || !t.Valid {
		return InvalidToken("SRV-1010")
	}

	claims := t.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	foundUser, err := uc.userRepository.FindById(userID)
	if err != nil || foundUser == nil {
		return NotFound("SRV-1001")
	}
	if err = uc.userRepository.SetVerified(userID); err != nil {
		return err
	}
	return nil
}
