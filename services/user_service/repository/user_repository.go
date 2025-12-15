package repository

import "user_service/domain"

type UserRepository interface {
	GetByEmail(email string) (*domain.User, error)
	GetByID(id int64) (*domain.User, error)
	Create(user *domain.User) error
	GetAll() ([]*domain.User, error)
	GetUserWithProfile(userID int64) (*domain.User, error)
	EmailExists(email string) (bool, error)
}
