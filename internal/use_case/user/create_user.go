package usecase

import (
	"github.com/rodrigodip/fighting-fantasy/internal/domain"
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
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
		Age:      newUser.Age,
	}
	err = ur.Repository.CreateUser(output.Name, output.Email, output.Password, output.Age)
	if err != nil {
		return UserDtoCreate{}, err
	}
	return output, nil
}
