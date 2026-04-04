package viewmodels

// PageData represents common data for all pages
type PageData struct {
	Title       string
	User        *User
	Hero        *HeroViewModel
	CurrentPage int
	Story       *Story
	Stats       *GameStats
	FeedBack    *Message
}
type Message struct {
	Success, Error string
}

// User represents an authenticated user
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// TODO: If it's a valuable info MUST be implemented at App Layer first
	// CreatedAt string `json:"created_at"`
}

// GameStats represents final adventure statistics
type GameStats struct {
	BattlesWon    int `json:"battles_won"`
	GoldCollected int `json:"gold_collected"`
	AreasExplored int `json:"areas_explored"`
}
