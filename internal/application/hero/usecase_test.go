package heroapp

import (
	"testing"

	"github.com/rodrigodip/fighting-fantasy/internal/domain/hero"
	"github.com/rodrigodip/fighting-fantasy/internal/mocks"
	"go.uber.org/mock/gomock"
)

func Test_CreateHero(t *testing.T) {

	ctrl := gomock.NewController(t)
	repositoryMock := mocks.NewMockRepository(ctrl)
	heroService := hero.NewHeroService(repositoryMock)
	heroUsecase := NewHeroUseCase(heroService)

	testCases := []struct {
		testName       string
		hero           hero.Hero
		potion         string
		expectedErrMsg string
		expectError    bool
	}{
		{
			testName: "When UserId is Nil",
			hero: hero.Hero{
				UserID:   "",
				HeroName: "Hero",
			},
			potion:         "dexterity",
			expectedErrMsg: "ValidadeHero: userID is Required",
			expectError:    true,
		},
		{
			testName: "When HeroName is Nil",
			hero: hero.Hero{
				UserID:   "123",
				HeroName: "",
			},
			potion:         "dexterity",
			expectedErrMsg: "ValidadeHero: heroName is Required",
			expectError:    true,
		},
		{
			testName: "When HeroName has less then 3 digits",
			hero: hero.Hero{
				UserID:   "123",
				HeroName: "He",
			},
			potion:         "dexterity",
			expectedErrMsg: "ValidadeHero: heroName must have more than 3 digits",
			expectError:    true,
		},
		{
			testName: "When HeroName has more then 20 digits",
			hero: hero.Hero{
				UserID:   "123",
				HeroName: "This name is forbbiden",
			},
			potion:         "dexterity",
			expectedErrMsg: "ValidadeHero: heroName must have less than 20 digits",
			expectError:    true,
		},
		{
			testName: "When Potion is Nil",
			hero: hero.Hero{
				UserID:   "123",
				HeroName: "Hero",
			},
			potion:         "",
			expectedErrMsg: "ValidadeHero: potion is Required",
			expectError:    true,
		},
		{
			testName: "When Potion has a invalid name",
			hero: hero.Hero{
				UserID:   "123",
				HeroName: "Hero",
			},
			potion:         "Dex",
			expectedErrMsg: "SelectPotion: invalid potion name",
			expectError:    true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, err := heroUsecase.CreateHero(
				tc.hero.UserID,
				tc.hero.HeroName,
				tc.potion,
			)
			if err != nil && !tc.expectError {
				t.Error(err.Error())
			}
			//TODO: CREATE custom Errors, do not compare strings
			if tc.expectedErrMsg == err.Error() {
				t.Errorf("Expected: %s >> Received: %s", tc.expectedErrMsg, err)
			}
		})
	}
	t.Run("sucess", func(t *testing.T) {
		testHero := hero.Hero{
			UserID:   "123",
			HeroName: "Test",
		}
		repositoryMock.EXPECT().RegisterHero(testHero).Return(nil)
		err := heroUsecase.HeroService.Save(testHero)
		if err != nil {
			t.Error(err)
		}
	})
}
