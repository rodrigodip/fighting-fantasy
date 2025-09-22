package user

type Repository interface {
	RegisterUser(id, name, email string, password []byte, role, status string) error
	FindByEmail(email string) (*User, error)
	FindById(id string) (*User, error)
	SetVerified(id string) error
	//GetUser(id string)(User, error)
	//GetUserByEmail(email, password string) (User, error)
	//UpdateUser(id, name, email string, age int)(User, error)
	//DeleteUser(id string) error
}
