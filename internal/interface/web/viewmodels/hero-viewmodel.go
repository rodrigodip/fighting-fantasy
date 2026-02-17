package viewmodels

import "github.com/rodrigodip/fighting-fantasy/internal/domain/hero"

type HeroViewModel struct {
	HeroName    string
	InitialHP   int
	CurrentHP   int
	CurrentDex  int
	CurrentLuck int
}

func NewHeroViewModel(h *hero.Hero) HeroViewModel {
	return HeroViewModel{
		HeroName:    h.HeroName,
		InitialHP:   h.Stats.InitialHP,
		CurrentHP:   h.Stats.CurrentHP,
		CurrentDex:  h.Stats.CurrentDex,
		CurrentLuck: h.Stats.CurrentLuck,
	}
}
