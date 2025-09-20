package auth

type Status string

const (
	StatusVerified   Status = "VERIFIED"
	StatusUnverified Status = "UNVERIFIED"
)

// StatusCheck
func IsUserVerified(status Status) bool {
	return status == StatusVerified
}
