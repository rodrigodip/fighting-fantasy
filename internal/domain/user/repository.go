package user

type Repository interface {
	RegisterUser(id, name, email string, password []byte, role, status string) error
	FindByEmail(email string) (*User, error)
	FindById(id string) (*User, error)
	SetVerified(id string) error
}
