package user

type Repository interface {
	RegisterUser(id, name, email, password string, role []string) error
	//GetUser(id string)(User, error)
	//GetUserByEmail(email, password string) (User, error)
	//UpdateUser(id, name, email string, age int)(User, error)
	//DeleteUser(id string) error
}
