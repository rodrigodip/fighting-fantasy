package herohandler

type HeroCreateRequest struct {
	UserID   string `json:"userId"`
	HeroName string `json:"hero_name"`
}
