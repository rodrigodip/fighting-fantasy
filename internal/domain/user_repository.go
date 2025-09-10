package domain

type UserRepository interface {
	CreateUser(name, email, pass string, age int) error
	//GetUser(id string)(User, error)
	//GetUserByEmailAndPass(email, password string) (User, error)
	//UpdateUser(id, name, email string, age int)(User, error)
	//DeleteUser(id string) error
}
