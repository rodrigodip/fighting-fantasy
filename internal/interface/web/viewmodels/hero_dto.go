package viewmodels

type HeroViewModel struct {
	Name       string
	Strength   int
	Dexterity  int
	Fortune    int
	Provisions int
	Gold       int
	Jewels     int
	Equipped   []string
	Backpack   []string
}

type Story struct {
	Lore    int
	Choices map[string]string
}

type HeroCreateReq struct {
	HeroName string `form:"name"`
	Potion   string `form:"potion"`
}
