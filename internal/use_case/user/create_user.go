package usecase

import (
	"github.com/rodrigodip/fighting-fantasy/internal/domain"
	IDgenerator "github.com/rodrigodip/fighting-fantasy/pkg/id_generator"
)

func (ur *UserRepository) CreateUser(name, email, pass string, age int) (UserDtoCreate, error) {
	newUser := domain.NewUserDomain()
	newUser.Name = name
	newUser.Email = email
	newUser.Password = pass
	newUser.Age = age
	err := newUser.UserValidation()
	if err != nil {
		return UserDtoCreate{}, err
	}
	output := UserDtoCreate{
		ID:       IDgenerator.NewSimpleID(),
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
		Age:      newUser.Age,
	}
	//TODO:Insert persistence func
	return output, nil
}
