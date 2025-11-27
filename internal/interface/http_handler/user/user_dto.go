package userhandler

type UserCreateRequest struct {
	//UserID   string `json:"userId,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	// Role     string `json:"roles,omitempty"`
}
type AuthLoginRequest struct {
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}
type UserCreateResponse struct {
	UserID string `json:"userId,omitempty"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	Role   string `json:"roles,omitempty"`
}
type UserFindByEmailRequest struct {
	UserID string `json:"userId,omitempty"`
	Email  string `json:"email,omitempty"`
}

type HeroCreateRequest struct {
	UserID   string `json:"userId"`
	HeroName string `json:"hero_name"`
	Potion   string `json:"selected_potion"`
}
