package userhandler

type UserCreateRequest struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
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
