// Package usecase implements business logic for user_service
package usecase

import (
	"errors"
	"user_services/domain"
	"user_services/repository"
)

// تعريف الأخطاء الموحدة
var (
	ErrUserNotFound   = errors.New("user not found")
	ErrInvalidUserID  = errors.New("invalid user id")
)

// UserUsecase encapsulates user-related business logic.
type UserUsecase struct {
	repo repository.UserRepository
}

// NewUserUsecase creates a new UserUsecase instance.
func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

// GetByID retrieves a user by ID.
func (u *UserUsecase) GetByID(id int64) (*domain.User, error) {
	if id <= 0 {
		return nil, ErrInvalidUserID
	}

	user, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}

// GetAll retrieves all users.
func (u *UserUsecase) GetAll() ([]*domain.User, error) {
	return u.repo.GetAll()
}

// CreateUser creates a new user.
func (u *UserUsecase) CreateUser(user *domain.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("name and email are required")
	}
	return u.repo.Create(user)
}

// UpdateUser updates an existing user.
func (u *UserUsecase) UpdateUser(user *domain.User) error {
	if user.ID <= 0 {
		return ErrInvalidUserID
	}

	existing, err := u.repo.GetByID(user.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrUserNotFound
	}

	return u.repo.Update(user)
}

// DeleteUser deletes a user by ID.
func (u *UserUsecase) DeleteUser(id int64) error {
	if id <= 0 {
		return ErrInvalidUserID
	}

	existing, err := u.repo.GetByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrUserNotFound
	}

	return u.repo.Delete(id)
}
