package hero

type Hero struct {
	UserID          string
	HeroName        string
	CurrentLore     int
	Stats           Stats
	Inventory       Inventory
	Potions         Potions
	Story           []Story
	MonsterDefeated []string
}
type Stats struct {
	InitialDex  int
	CurrentDex  int
	InitialHP   int
	CurrentHP   int
	InitialLuck int
	CurrentLuck int
}
type Inventory struct {
	Equipment  []string
	Backpack   []string
	Provisions int
	Gold       int
	Jewels     int
}
type Potions struct {
	Dexterity bool
	Strength  bool
	Fortune   bool
}
type Story struct {
	Lore    int
	Choices map[string]string
}
