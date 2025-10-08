package hero

type Repository interface {
	RegisterHero(hero Hero) error
	FindByOwner(userID string) (*Hero, error)
}
