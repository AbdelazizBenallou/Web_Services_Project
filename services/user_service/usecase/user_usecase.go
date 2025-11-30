// Package usecase implements application business logic.
package usecase

import (
	"user_services/domain"
	"user_services/repository"
)

// UserUsecase encapsulates user-related business logic.
type UserUsecase struct {
	repo repository.UserRepository
}

// NewUserUsecase creates a new UserUsecase.
func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

// GetByID retrieves a user by ID.
func (u *UserUsecase) GetByID(id int64) (*domain.User, error) {
	return u.repo.GetByID(id)
}

// GetAll retrieves all users.
func (u *UserUsecase) GetAll() ([]*domain.User, error) {
	return u.repo.GetAll()
}
