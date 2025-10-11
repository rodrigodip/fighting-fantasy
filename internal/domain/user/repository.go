package user

type Repository interface {
	//RegisterUser(id, name, email string, password []byte, role, status string) error
	RegisterUser(u User) error
	FindByEmail(email string) (*User, error)
	FindById(id string) (*User, error)
	SetVerified(id string) error
	Login(email, password string) (string, error)
	VerifyEmail(token string) error
	//TODO: ResendVeriyEmail()

}
