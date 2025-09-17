package authhandler

type AuthLoginRequest struct {
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}
