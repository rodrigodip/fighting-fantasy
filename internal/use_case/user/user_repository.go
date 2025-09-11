package usecase

import "github.com/rodrigodip/fighting-fantasy/internal/domain"

type UserRepository struct {
	Repository domain.UserRepository
}

func NewUserRepository(repository domain.UserRepository) *UserRepository {
	return &UserRepository{Repository: repository}
}
