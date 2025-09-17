package auth

type Role string

const (
	RoleUser  Role = "USER"
	RoleAdmin Role = "ADMIN"
)

// Simple rule check
func CanAccessAdmin(role Role) bool {
	return role == RoleAdmin
}
