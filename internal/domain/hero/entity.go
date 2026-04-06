package hero

type Hero struct {
	UserID      string
	HeroName    string
	CurrentLore int
	Stats       Stats
	Inventory   Inventory
	Potions     Potions
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
	Equipment []string
	Backpack  *Backpack
}
type Backpack struct {
	Provisions int
	Gold       int
	Jewels     int
	Itens      []string
}
type Potions struct {
	Dexterity bool
	Strength  bool
	Fortune   bool
}
