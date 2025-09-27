package hero

type Repository interface {
	RegisterHero(userID, heroName string) error
}
