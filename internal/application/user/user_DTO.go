package userapp

type UserDtoCreate struct {
	UserID   string   `json:"userId"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}
type UserDtoOutput struct {
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Email string   `json:"email"`
	Roles []string `json:"roles"`
}
type UserDtoLogin struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type UserDtoUpdate struct {
	UserID string   `json:"userId"`
	Name   string   `json:"name"`
	Email  string   `json:"email"`
	Roles  []string `json:"roles"`
	Age    int      `json:"age"`
	Gender string   `json:"gender"`
}
