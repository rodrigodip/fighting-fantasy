package advtrapp

type AdventureRepository interface {
	RegisterLore()
	FindLoreByID()
}

type AdventureUseCase struct {
	repo AdventureRepository
}

func NewAdventureUseCase(r AdventureRepository) *AdventureUseCase {
	return &AdventureUseCase{repo: r}
}

func (uc *AdventureUseCase) Lore() {}
