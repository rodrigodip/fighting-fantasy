package herohandler

type HeroCreateRequest struct {
	HeroName string `json:"heroName"`
	Potion   string `json:"potion"`
}
