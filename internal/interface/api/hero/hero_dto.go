package herohandler

type HeroCreateRequest struct {
	HeroName string `json:"hero_name"`
	Potion   string `json:"selected_potion"`
}
