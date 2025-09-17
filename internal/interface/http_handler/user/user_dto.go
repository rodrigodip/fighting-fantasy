package userhandler

type UserCreateRequest struct {
	UserID   string   `json:"userId,omitempty"`
	Name     string   `json:"name,omitempty"`
	Email    string   `json:"email,omitempty"`
	Password string   `json:"password,omitempty"`
	Roles    []string `json:"roles,omitempty"`
}

type UserCreateResponse struct {
	UserID string   `json:"userId,omitempty"`
	Name   string   `json:"name,omitempty"`
	Email  string   `json:"email,omitempty"`
	Roles  []string `json:"roles,omitempty"`
}
type UserFindByEmailRequest struct {
	UserID string `json:"userId,omitempty"`
	Email  string `json:"email,omitempty"`
}
