package usecase

import (
	"github.com/rodrigodip/fighting-fantasy/internal/domain"
	IDgenerator "github.com/rodrigodip/fighting-fantasy/pkg/id_generator"
)

func (ur *UserRepository) CreateUser(name, email, password string) (UserDtoCreate, error) {
	newUser := domain.NewUserDomain()
	newUser.Name = name
	newUser.Email = email
	newUser.Password = password
	newUser.Roles = []string{"user"}
	err := newUser.UserValidation()
	if err != nil {
		return UserDtoCreate{}, err
	}
	output := UserDtoCreate{
		UserID:   IDgenerator.NewSimpleID(),
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
		Roles:    newUser.Roles,
	}
	err = ur.Repository.
		CreateUser(
			output.UserID,
			output.Name,
			output.Email,
			output.Password,
			output.Roles,
		)
	if err != nil {
		return UserDtoCreate{}, err
	}
	return output, nil
}
