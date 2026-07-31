package hero

import (
	"testing"
)

func TestService_ValidateInput(t *testing.T) {
	s := NewHeroService()

	tests := []struct {
		name     string
		heroName string
		potion   string
		wantErr  bool
	}{
		{"empty name", "", "strength", true},
		{"short name", "AB", "strength", true},
		{"long name", "ThisIsAVeryLongHeroNameThatExceedsTwentyCharacters", "strength", true},
		{"empty potion", "HeroName", "", true},
		{"valid input", "HeroName", "strength", false},
		{"valid input min length", "ABC", "dexterity", false},
		{"valid input max length", "Exactly20Characters", "fortune", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.ValidateInput(tt.heroName, tt.potion)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_HasHero(t *testing.T) {
	s := NewHeroService()

	tests := []struct {
		name    string
		userID  string
		hero    Hero
		wantErr bool
	}{
		{"user has hero", "user123", Hero{UserID: "user123"}, true},
		{"user doesn't have hero", "user123", Hero{UserID: "user456"}, false},
		{"empty user IDs match", "", Hero{UserID: ""}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.HasHero(tt.userID, tt.hero)
			if (err != nil) != tt.wantErr {
				t.Errorf("HasHero() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_SelectPotion(t *testing.T) {
	s := NewHeroService()

	tests := []struct {
		name           string
		potion         string
		wantDexterity  bool
		wantStrength   bool
		wantFortune    bool
		wantErr        bool
	}{
		{"dexterity", "dexterity", true, false, false, false},
		{"strength", "strength", false, true, false, false},
		{"fortune", "fortune", false, false, true, false},
		{"invalid potion", "invalid", false, false, false, true},
		{"empty potion", "", false, false, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hero := Hero{}
			result, err := s.SelectPotion(hero, tt.potion)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectPotion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if result.Potions.Dexterity != tt.wantDexterity {
					t.Errorf("SelectPotion() Dexterity = %v, want %v", result.Potions.Dexterity, tt.wantDexterity)
				}
				if result.Potions.Strength != tt.wantStrength {
					t.Errorf("SelectPotion() Strength = %v, want %v", result.Potions.Strength, tt.wantStrength)
				}
				if result.Potions.Fortune != tt.wantFortune {
					t.Errorf("SelectPotion() Fortune = %v, want %v", result.Potions.Fortune, tt.wantFortune)
				}
			}
		})
	}
}
