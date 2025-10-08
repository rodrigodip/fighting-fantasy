package herohandler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/fighting-fantasy/internal/application/hero"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/hero"
	"github.com/rodrigodip/fighting-fantasy/internal/mocks"
	"go.uber.org/mock/gomock"
)

func Test_RegisterHero(t *testing.T) {

	ctrl := gomock.NewController(t)
	repositoryMock := mocks.NewMockRepository(ctrl)
	heroService := hero.NewHeroService(repositoryMock)
	heroUsecase := heroapp.NewHeroUseCase(heroService)
	heroHandler := NewHeroHandler(heroUsecase)

	r := gin.New()
	r.POST("/heroes", heroHandler.RegisterHero)

	testCases := []struct {
		testName       string
		hero           HeroCreateRequest
		expectedReturn string
		expectedStatus int
		expectError    bool
	}{
		{
			testName:       "When User is Nil",
			hero:           HeroCreateRequest{"", "hero Test", "Dexterity"},
			expectedReturn: "ValidadeHero: userID is Required",
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			testName:       "When HeroName is Nil",
			hero:           HeroCreateRequest{"123", "", "Dexterity"},
			expectedReturn: "ValidadeHero: heroName is Required",
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			testName:       "When HeroName has less than 3 digits",
			hero:           HeroCreateRequest{"", "hero Test", "Dexterity"},
			expectedReturn: "ValidadeHero: heroName must have more than 3 digits",
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			testName:       "When HeroName has more than 20 digits",
			hero:           HeroCreateRequest{"", "hero Test", "Dexterity"},
			expectedReturn: "ValidadeHero: heroName must have less than 20 digits",
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			testName:       "When Potion is Nil",
			hero:           HeroCreateRequest{"", "hero Test", "Dexterity"},
			expectedReturn: "ValidadeHero: a potion name is Required",
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			testName:       "SUCESS",
			hero:           HeroCreateRequest{"", "hero Test", "Dexterity"},
			expectedReturn: "ValidadeHero: a potion name is Required",
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
	}
	for _, tc := range testCases {
		w := httptest.NewRecorder()
		userJson, _ := json.Marshal(tc.hero)
		req, _ := http.NewRequest("POST", "/heroes", strings.NewReader(string(userJson)))
		r.ServeHTTP(w, req)
	}
}
