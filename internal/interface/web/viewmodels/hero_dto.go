package viewmodels

type HeroViewModel struct {
	HeroName    string
	InitialHP   int
	CurrentHP   int
	CurrentDex  int
	CurrentLuck int
}

type Story struct {
	Lore    int
	Choices map[string]string
}
