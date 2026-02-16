package viewmodels

import "github.com/rodrigodip/fighting-fantasy/internal/domain/hero"

type HeroViewModel struct {
	HeroName string
	HP       int
	Dex      int
	Luck     int
}

func NewHeroViewModel(h *hero.Hero) HeroViewModel {
	return HeroViewModel{
		HeroName: h.HeroName,
		HP:       h.Stats.CurrentHP,
		Dex:      h.Stats.CurrentDex,
		Luck:     h.Stats.CurrentLuck,
	}
}
