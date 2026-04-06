package viewmodels

type AdventurePageData struct {
	CurrentPage int
	Story       StoryViewModel
	Hero        HeroViewModel
}
type StoryViewModel struct {
	Text    string
	Battle  *BattleViewModel // nil = regular choices, non-nil = battle mode
	Choices []ChoiceViewModel
}
type BattleViewModel struct {
	Enemy     string
	Strength  int
	Dexterity int
}
type ChoiceViewModel struct {
	Text string
	Page int // page number for hx-get="/adventure/page/{{.Page}}"
}
