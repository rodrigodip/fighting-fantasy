package adveture

type Segment struct {
	LoreID        int
	Text          string
	Combat        *Combat
	LuckTest      *LuckTest
	AttributeTest *AttributeTest
	ItemCheck     *ItemCheck
	Effects       []Effect
	Choices       []Choice
	IsTerminal    bool
}

type Combat struct {
	Escape      *Escape
	Penalty     *Penalty
	Waves       []Wave
	VictoryGoto int
	DefeatGoto  *int
}

type Escape struct {
	Available   bool
	AfterRounds int
	Goto        int
	StaminaCost int
}

type Penalty struct {
	HeroAttackModifier int
	Description        string
}

type Wave struct {
	Description string
	Mode        string
	Enemies     []Enemy
}

type Enemy struct {
	Name    string
	Skill   int
	Stamina int
}

type LuckTest struct {
	Description string
	Success     Outcome
	Failure     Outcome
}

type Outcome struct {
	Goto      int
	Effects   []Effect
	Condition *Condition
}

type Condition struct {
	MustSurvive bool
}

type AttributeTest struct {
	// Extensível futuramente
}

type ItemCheck struct {
	Item        string
	SuccessGoto int
	FailureGoto int
}

type Effect struct {
	Type      string
	Operation string
	Value     int
}

type Choice struct {
	Description *string
	Goto        int
}
