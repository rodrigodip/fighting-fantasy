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
	Status   string
	Profile  Profile
	Role     string
}
