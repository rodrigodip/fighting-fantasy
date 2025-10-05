package herohandler

import (
	// "encoding/json"
	// "net/http"
	// "net/http/httptest"
	// "strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/fighting-fantasy/internal/application/hero"
	"github.com/rodrigodip/fighting-fantasy/internal/domain/hero"
	"github.com/rodrigodip/fighting-fantasy/internal/mocks"
	"go.uber.org/mock/gomock"
)

func Test_RegisterHero(t *testing.T) {
	// stubHero := HeroCreateRequest{
	// 	UserID:   "1234",
	// 	HeroName: "Hero-test",
	// 	Potion:   "Dexterity",
	// }
	// heroJson, _ := json.Marshal(stubHero)

	ctrl := gomock.NewController(t)
	repositoryMock := mocks.NewMockRepository(ctrl)
	heroService := hero.NewHeroService(repositoryMock)
	heroUsecase := heroapp.NewHeroUseCase(heroService)
	heroHandler := NewHeroHandler(heroUsecase)

	r := gin.New()
	r.POST("/heroes", heroHandler.RegisterHero)

	// testCase := []struct {
	// 	name           string
	// 	id             string
	// 	setupMocks     func()
	// 	expectedStatus int
	// 	expectError    bool
	// }{
	// 	{
	// 		name: "",
	// 	},
	// }
	//
	// w := httptest.NewRecorder()
	// userJson, _ := json.Marshal(heroJson)
	// req, _ := http.NewRequest("POST", "/heroes", strings.NewReader(string(userJson)))
	// r.ServeHTTP(w, req)
}
