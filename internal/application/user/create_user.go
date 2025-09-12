package userapp

import "github.com/rodrigodip/fighting-fantasy/internal/domain/user"

// "github.com/rodrigodip/fighting-fantasy/internal/domain/user"
// IDgenerator "github.com/rodrigodip/fighting-fantasy/internal/pkg/id_generator"
func (uc *UserUseCase) CreateUser(name, email, password string) (*user.User, error) {
	return uc.UserService.CreateUser(name, email, password)
}

// func (ur *UserUseCase) CreateUser(name, email, password string) (UserDtoCreate, error) {
// 	newUser := user.NewUserDomain()
// 	newUser.Name = name
// 	newUser.Email = email
// 	newUser.Password = password
// 	newUser.Roles = []string{"user"}
// 	err := newUser.UserValidation()
// 	if err != nil {
// 		return UserDtoCreate{}, err
// 	}
// 	output := UserDtoCreate{
// 		UserID:   IDgenerator.NewSimpleID(),
// 		Name:     newUser.Name,
// 		Email:    newUser.Email,
// 		Password: newUser.Password,
// 		Roles:    newUser.Roles,
// 	}
// 	err = ur.Repository.
// 		CreateUser(
// 			output.UserID,
// 			output.Name,
// 			output.Email,
// 			output.Password,
// 			output.Roles,
// 		)
// 	if err != nil {
// 		return UserDtoCreate{}, err
// 	}
// 	return output, nil
// }
