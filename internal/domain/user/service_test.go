package usr

import (
	"testing"
)

func TestService_ValidadeUserInput(t *testing.T) {
	s := NewUserService()

	tests := []struct {
		name     string
		userName string
		email    string
		password string
		wantErr  bool
	}{
		{"empty name", "", "test@example.com", "Password123!", true},
		{"short name", "AB", "test@example.com", "Password123!", true},
		{"empty email", "TestUser", "", "Password123!", true},
		{"invalid email", "TestUser", "notanemail", "Password123!", true},
		{"empty password", "TestUser", "test@example.com", "", true},
		{"valid input", "TestUser", "test@example.com", "Password123!", false},
		{"valid input min name length", "ABC", "test@example.com", "Password123!", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.ValidadeUserInput(tt.userName, tt.email, tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidadeUserInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_ValidatePassword(t *testing.T) {
	s := NewUserService()

	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{"too short", "Pass1!", true},
		{"no lowercase", "PASSWORD123!", true},
		{"no uppercase", "password123!", true},
		{"no digit", "Password!", true},
		{"no special char", "Password123", true},
		{"valid password", "Password123!", false},
		{"valid password with @", "Password123@", false},
		{"valid password with $", "Password123$", false},
		{"valid password with %", "Password123%", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.ValidatePassword(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_IsUserVerified(t *testing.T) {
	s := NewUserService()

	tests := []struct {
		name   string
		status Status
		want   bool
	}{
		{"verified status", StatusVerified, false},
		{"unverified status", StatusUnverified, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := s.IsUserVerified(tt.status)
			if got != tt.want {
				t.Errorf("IsUserVerified() = %v, want %v", got, tt.want)
			}
		})
	}
}
