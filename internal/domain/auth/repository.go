package auth

type Repository interface {
	Login(email, password string) (string, error)
	VerifyEmail(token string) error
	//TODO: ResendVeriyEmail()
}
