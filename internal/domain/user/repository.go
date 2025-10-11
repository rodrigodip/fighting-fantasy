package usr

type Repository interface {
	RegisterUser(u User) error
	FindByEmail(email string) (*User, error)
	FindById(id string) (*User, error)
	Update(id, field, newData string) error
	// Login(email, password string) (string, error)
	// VerifyEmail(token string) error
	//TODO: ResendVeriyEmail()

}
