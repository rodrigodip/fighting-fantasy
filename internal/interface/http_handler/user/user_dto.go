package userhandler

type UserCreateRequest struct {
	UserID   string   `json:"userId"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}

type UserCreateResponse struct {
	UserID string   `json:"userId"`
	Name   string   `json:"name"`
	Email  string   `json:"email"`
	Roles  []string `json:"roles"`
}
