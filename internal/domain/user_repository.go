package domain

type UserRepository interface {
	CreateUser(id, name, email, pass string, age int) error
	//GetUser(id string)(User, error)
	//UpdateUser(id, name, email string, age int)(User, error)
	//DeleteUser(id string) error
}
