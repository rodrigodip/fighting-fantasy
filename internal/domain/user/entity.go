package user

type Profile struct {
	Age    int
	Gender string
}

type User struct {
	UserID   string
	Name     string
	Email    string
	Password []byte
	Status   Status
	Profile  Profile
	Role     string
}
type (
	Status string
	Role   string
)

const (
	StatusVerified   Status = "VERIFIED"
	StatusUnverified Status = "UNVERIFIED"
	RoleUser         Role   = "USER"
	RoleAdmin        Role   = "ADMIN"
)
