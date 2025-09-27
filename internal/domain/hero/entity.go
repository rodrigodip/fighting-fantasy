package hero

import "time"

type Hero struct {
	UserID          string
	HeroName        string
	CurrentLore     int
	Stats           Stats
	Inventory       Inventory
	Potions         Potions
	History         []History
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
type History struct {
	Lore      int
	Timestamp time.Time
	Choices   map[string]string
}
