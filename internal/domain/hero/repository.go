package hero

type Repository interface {
	RegisterHero(hero Hero) error
}
