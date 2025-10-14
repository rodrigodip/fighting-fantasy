package usr

type Repository interface {
	RegisterUser(u User) error
	FindByEmail(email string) (*User, error)
	FindById(id string) (*User, error)
	Update(id, field, newData string) error

	//TODO: ResendVeriyEmail()
	//FindHero(userID string) (*Hero, error)
}
